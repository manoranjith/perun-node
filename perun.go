// Copyright (c) 2020 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/hyperledger-labs/perun-node
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package perun

import (
	"context"
	"math/big"
	"time"

	pchannel "perun.network/go-perun/channel"
	ppersistence "perun.network/go-perun/channel/persistence"
	pclient "perun.network/go-perun/client"
	pLog "perun.network/go-perun/log"
	pwallet "perun.network/go-perun/wallet"
	pwire "perun.network/go-perun/wire"
	pnet "perun.network/go-perun/wire/net"
)

// PeerID represents any participant in the off-chain network that the user wants to transact with.
type PeerID struct {
	// Name assigned by user for referring to this PeerID in API requests to the node.
	// It is unique within a session on the node.
	Alias string `yaml:"alias"`

	// Permanent identity used for authenticating the PeerID in the off-chain network.
	OffChainAddr pwire.Address `yaml:"-"`
	// This field holds the string value of address for easy marshaling / unmarshaling.
	OffChainAddrString string `yaml:"offchain_address"`

	// Address for off-chain communication.
	CommAddr string `yaml:"comm_address"`
	// Type of off-chain communication protocol.
	CommType string `yaml:"comm_type"`
}

// OwnAlias is the alias for the entry of the user's own PeerID details.
// It will be used when translating addresses in incoming messages / proposals to aliases.
const OwnAlias = "self"

// IDReader represents the functions to read peer IDs from a cache connected to a peer ID provider.
type IDReader interface {
	ReadByAlias(alias string) (p PeerID, contains bool)
	ReadByOffChainAddr(offChainAddr pwire.Address) (p PeerID, contains bool)
}

// IDProvider represents the functions to read, write peer IDs from and to the local cache connected to a
// peer ID provider. It also includes a function to sync the changes in the cache with the ID provider backend.
type IDProvider interface {
	IDReader
	Write(alias string, p PeerID) error
	Delete(alias string) error
	UpdateStorage() error
}

//go:generate mockery --name CommBackend --output ./internal/mocks

// CommBackend defines the set of methods required for initializing components required for off-chain communication.
// This can be protocols such as tcp, websockets, MQTT.
type CommBackend interface {
	// Returns a listener that can listen for incoming messages at the specified address.
	NewListener(address string) (pnet.Listener, error)

	// Returns a dialer that can dial for new outgoing connections.
	// If timeout is zero, program will use no timeout, but standard OS timeouts may still apply.
	NewDialer() Dialer
}

//go:generate mockery --name Dialer --output ./internal/mocks

// Dialer extends net.Dialer with Registerer interface.
type Dialer interface {
	pnet.Dialer
	Registerer
}

//go:generate mockery --name Registerer --output ./internal/mocks

// Registerer is used to register the commAddr corresponding to an offChainAddr to the wire.Bus in runtime.
type Registerer interface {
	Register(offChainAddr pwire.Address, commAddr string)
}

// Credential represents the parameters required to access the keys and make signatures for a given address.
type Credential struct {
	Addr     pwallet.Address
	Wallet   pwallet.Wallet
	Keystore string
	Password string
}

// User represents a participant in the off-chain network that uses a session on this node for sending transactions.
type User struct {
	PeerID

	OnChain  Credential // Account for funding the channel and the on-chain transactions.
	OffChain Credential // Account (corresponding to off-chain address) used for signing authentication messages.

	// List of participant addresses for this user in each open channel.
	// OffChain credential is used for managing all these accounts.
	PartAddrs []pwallet.Address
}

// Session provides a context for the user to interact with a node. It manages user data (such as keys, peer IDs),
// and channel client.
//
// Once established, a user can establish and transact on state channels. All the channels within a session will use
// the same type and version of communication and state channel protocol. If a user desires to use multiple types or
// versions of any protocol, it should request a separate session for each combination of type and version of those.
type Session struct {
	ID   string // ID uniquely identifies a session instance.
	User User

	ChClient ChClient
}

//go:generate mockery --name Channel --output ./internal/mocks

// Channel represents  state channel established among the participants of the off-chain network.
type Channel interface {
	Close() error
	ID() pchannel.ID
	Idx() pchannel.Index
	IsClosed() bool
	Params() *pchannel.Params
	Peers() []pwire.Address
	Phase() pchannel.Phase
	State() *pchannel.State
	OnUpdate(cb func(from, to *pchannel.State))
	UpdateBy(ctx context.Context, update func(*pchannel.State) error) error
	Register(ctx context.Context) error
	Settle(ctx context.Context, isSecondary bool) error
	Watch(pclient.AdjudicatorEventHandler) error
}

//go:generate mockery --name ChClient --output ./internal/mocks

// ChClient allows the user to establish off-chain channels and transact on these channels.
//
// It allows the user to enable persistence, where all data pertaining to the lifecycle of a channel is
// persisted continuously. When it is enabled, the channel client can be stopped at any point of time and resumed later.
//
// However, the channel client is not responsible if any channel the user was participating in was closed
// with a wrong state when the channel client was not running.
// Hence it is highly recommended not to stop the channel client if there are open channels.
type ChClient interface {
	Registerer
	ProposeChannel(context.Context, pclient.ChannelProposal) (Channel, error)
	Handle(pclient.ProposalHandler, pclient.UpdateHandler)
	Channel(pchannel.ID) (Channel, error)
	Close() error

	EnablePersistence(ppersistence.PersistRestorer)
	OnNewChannel(handler func(Channel))
	Restore(context.Context) error
	RestoreChs(func(Channel)) error

	Log() pLog.Logger
}

//go:generate mockery --name WireBus --output ./internal/mocks

// WireBus is an extension of the wire.Bus interface in go-perun to include a "Close" method.
// pwire.Bus (in go-perun) is a central message bus over which all clients of a channel network
// communicate. It is used as the transport layer abstraction for the ChClient.
type WireBus interface {
	pwire.Bus
	Close() error
}

// ChainBackend wraps the methods required for instantiating and using components for
// making on-chain transactions and reading on-chain values on a specific blockchain platform.
// The timeout for on-chain transaction should be implemented by the corresponding backend. It is
// up to the implementation to make the value user configurable.
//
// It defines methods for deploying contracts; validating deployed contracts and instantiating a funder, adjudicator.
type ChainBackend interface {
	DeployAdjudicator(onChainAddr pwallet.Address) (adjAddr pwallet.Address, _ error)
	DeployAsset(adjAddr, onChainAddr pwallet.Address) (assetAddr pwallet.Address, _ error)
	ValidateContracts(adjAddr, assetAddr pwallet.Address) error
	NewFunder(assetAddr, onChainAddr pwallet.Address) pchannel.Funder
	NewAdjudicator(adjAddr, receiverAddr pwallet.Address) pchannel.Adjudicator
}

// WalletBackend wraps the methods for instantiating wallets and accounts that are specific to a blockchain platform.
type WalletBackend interface {
	ParseAddr(string) (pwallet.Address, error)
	NewWallet(keystore string, password string) (pwallet.Wallet, error)
	UnlockAccount(pwallet.Wallet, pwallet.Address) (pwallet.Account, error)
}

// Currency represents a parser that can convert between string representation of a currency and
// its equivalent value in base unit represented as a big integer.
type Currency interface {
	Parse(string) (*big.Int, error)
	Print(*big.Int) string
}

// NodeConfig represents the configurable parameters of a perun node.
type NodeConfig struct {
	// User configurable values.
	LogLevel         string        // LogLevel represents the log level for the node and all derived loggers.
	LogFile          string        // LogFile represents the file to write logs. Empty string represents stdout.
	ChainURL         string        // Address of the default blockchain node used by the perun node.
	Adjudicator      string        // Address of the default Adjudicator contract used by the perun node.
	Asset            string        // Address of the default Asset Holder contract used by the perun node.
	ChainConnTimeout time.Duration // Timeout for connecting to blockchain node.
	OnChainTxTimeout time.Duration // Timeout to wait for confirmation of on-chain tx.
	ResponseTimeout  time.Duration // Timeout to wait for a response from the peer / user.

	// Hard coded values. See cmd/perunnode/run.go.
	CommTypes            []string // Communication protocols supported by the node for off-chain communication.
	IDProviderTypes      []string // ID Provider types supported by the node.
	CurrencyInterpreters []string // Currencies Interpreters supported by the node.

}

// APIErrorV2 represents the newer version of error returned by node, session
// and channel APIs.
//
// Along with the error message, this error type assigns to each error an
// error category (that describes how the error should be handled by the
// client), error code (that identifies specific types of error), additional
// info (that contains additional data related to the error as key value
// pairs).
//
// TODO: (mano) Once all usage of APIError is replaced with APIErrorV2, the
// older type needs to be removed and this one to be renamed as APIError.
// Also, the usage of the string "V2" in the variables and types related to the
// new error type should be removed.
type APIErrorV2 interface {
	Category() ErrorCategory
	Code() ErrorCode
	Message() string
	AddInfo() interface{}
	Error() string
}

// ErrorCategory represents the category of the error, which describes how the
// error should be handled by the client.
type ErrorCategory int

const (
	// ParticipantError is caused by one of the channel participants not acting
	// as per the perun protocol.
	//
	// To resolve this, the client should negotiate with the peer outside of
	// this system to act in accordance with the perun protocol.
	ParticipantError ErrorCategory = iota

	// ClientError is caused by the errors in the request from the client. It
	// could be errors in arguments or errors in configuration provided by the
	// client to access the external systems or errors in the state of external
	// systems not managed by the node.
	//
	// To resolve this, the client should provide valid arguments, provide
	// correct configuration to access the external systems or fix the external
	// systems; and then retry.
	ClientError

	// ProtocolFatalError is caused when the protocol aborts due to unexpected
	// failure in external system during execution. It could also result in loss
	// of funds.
	//
	// To resolve this, user should manually inspect the error message and
	// handle it.
	ProtocolFatalError
	// InternalError is caused due to unintended behavior in the node software.
	//
	// To resolve this, user should manually inspect the error message and
	// handle it.
	InternalError
)

// String implements the stringer interface for ErrorCategory.
func (c ErrorCategory) String() string {
	return [...]string{
		"Client",
		"Participant",
		"Protocol Fatal",
		"Internal",
	}[c]
}

// ErrorCode is a numeric code assigned to identify the specific type of error.
// The keys in the additional field is fixed for each error code.
type ErrorCode int

// Error code definitions.
const (
	ErrV2PeerResponseTimedOut      ErrorCode = 101
	ErrV2RejectedByPeer            ErrorCode = 102
	ErrV2PeerNotFunded             ErrorCode = 103
	ErrV2UserResponseTimedOut      ErrorCode = 104
	ErrV2ResourceNotFound          ErrorCode = 201
	ErrV2ResourceExists            ErrorCode = 202
	ErrV2InvalidArgument           ErrorCode = 203
	ErrV2FailedPreCondition        ErrorCode = 204
	ErrV2InvalidConfig             ErrorCode = 205
	ErrV2ChainNodeNotReachable     ErrorCode = 206
	ErrV2InvalidContracts          ErrorCode = 207
	ErrV2TxTimedOut                ErrorCode = 301
	ErrV2InsufficientBalForTx      ErrorCode = 302
	ErrV2ChainNodeDisconnected     ErrorCode = 303
	ErrV2InsufficientBalForDeposit ErrorCode = 304
	ErrV2UnknownInternal           ErrorCode = 401
	ErrV2OffChainComm              ErrorCode = 402
)

type (
	// ErrV2InfoResourceNotFound represents the fields in the additional info for
	// ErrResourceNotFound.
	ErrV2InfoResourceNotFound struct {
		Type string
		ID   string
	}

	// ErrV2InfoResourceExists represents the fields in the additional info for
	// ErrResourceExists.
	ErrV2InfoResourceExists struct {
		Type string
		ID   string
	}

	// ErrV2InfoInvalidArgument represents the fields in the additional info for
	// ErrInvalidArgument.
	ErrV2InfoInvalidArgument struct {
		Name        string
		Value       string
		Requirement string
	}
)

// NodeAPI represents the APIs that can be accessed in the context of a perun node.
// Multiple sessions can be opened in a single node. Each instance will have a dedicated
// keystore and ID provider.
type NodeAPI interface {
	Time() int64
	GetConfig() NodeConfig
	Help() []string
	OpenSession(configFile string) (string, []ChInfo, error)

	// This function is used internally to get a SessionAPI instance.
	// Should not be exposed via user API.
	GetSession(string) (SessionAPI, error)

	// GetSessionV2 is a wrapper over GetSession that returns the error in the
	// newly defined APIErrorV2 format instead of the standard error.
	//
	// This is introduced temporarily to facilitate refactoring all the APIs to
	// return a newly defined API Error type instead of the standard error.
	// Since this API is used across all the api calls, two versions are
	// simulataneously required until all the APIs are updated to use the newly
	// defined API Error (APIErrorV2).
	//
	// TODO (mano): Once all the APIs are updated to use the new newly defined API
	// Error, remove the other version of GetSession API that returns standard
	// error type and rename this to GetSession.
	GetSessionV2(string) (SessionAPI, APIErrorV2)
}

//go:generate mockery --name SessionAPI --output ./internal/mocks

// SessionAPI represents the APIs that can be accessed in the context of a perun node.
// First a session has to be instantiated using the NodeAPI. The session can then be used
// open channels and accept channel proposals.
type SessionAPI interface {
	ID() string
	AddPeerID(PeerID) APIErrorV2
	GetPeerID(alias string) (PeerID, APIErrorV2)
	OpenCh(context.Context, BalInfo, App, uint64) (ChInfo, error)
	GetChsInfo() []ChInfo
	SubChProposals(ChProposalNotifier) APIErrorV2
	UnsubChProposals() APIErrorV2
	RespondChProposal(context.Context, string, bool) (ChInfo, error)
	Close(force bool) ([]ChInfo, error)

	// This function is used internally to get a ChAPI instance.
	// Should not be exposed via user API.
	GetCh(string) (ChAPI, error)

	// GetChV2 is a wrapper over GetCh that returns the error in the
	// newly defined APIErrorV2 format instead of the standard error.
	//
	// This is introduced temporarily to facilitate refactoring all the APIs to
	// return a newly defined API Error type instead of the standard error.
	// Since this API is used across all the api calls, two versions are
	// simulataneously required until all the APIs are updated to use the newly
	// defined API Error (APIErrorV2).
	//
	// TODO: Once all the APIs are updated to use the new newly defined API
	// Error, remove the other version of GetCh API that returns standard
	// error type and rename this to GetCh.
	GetChV2(string) (ChAPI, APIErrorV2)
}

type (
	// ChProposalNotifier is the notifier function that is used for sending channel proposal notifications.
	ChProposalNotifier func(ChProposalNotif)

	// ChProposalNotif represents the parameters sent in a channel proposal notifications.
	ChProposalNotif struct {
		ProposalID       string
		OpeningBalInfo   BalInfo
		App              App
		ChallengeDurSecs uint64
		Expiry           int64
	}
)

//go:generate mockery --name ChAPI --output ./internal/mocks

// ChAPI represents the APIs that can be accessed in the context of a perun channel.
// First a channel has to be initialized using the SessionAPI. The channel can then be used
// send and receive updates.
type ChAPI interface {
	// Methods for reading the channel information is doesn't change.
	// These APIs don't use mutex lock.
	ID() string
	Currency() string
	Parts() []string
	ChallengeDurSecs() uint64

	// Methods to transact on, close the channel and read its state.
	// These APIs use a mutex lock.
	SendChUpdate(context.Context, StateUpdater) (ChInfo, error)
	SubChUpdates(ChUpdateNotifier) APIErrorV2
	UnsubChUpdates() APIErrorV2
	RespondChUpdate(context.Context, string, bool) (ChInfo, error)
	GetChInfo() ChInfo
	Close(context.Context) (ChInfo, error)
}

// Enumeration of values for ChUpdateType:
// Open: If accepted, channel will be updated and it will remain in open for off-chain tx.
// Final: If accepted, channel will be updated and closed (settled on-chain and amount withdrawn).
// Closed: Channel has been closed (settled on-chain and amount withdrawn).
const (
	ChUpdateTypeOpen ChUpdateType = iota
	ChUpdateTypeFinal
	ChUpdateTypeClosed
)

type (
	// ChUpdateType is the type of channel update. It can have three values: "open", "final" and "closed".
	ChUpdateType uint8

	// ChUpdateNotifier is the notifier function that is used for sending channel update notifications.
	ChUpdateNotifier func(ChUpdateNotif)

	// ChUpdateNotif represents the parameters sent in a channel update notification.
	// The update can be of two types
	// 1. Regular update proposed by the peer to progress the off-chain state of the channel.
	// 2. Closing update when a channel is closed, balance is settled on the blockchain and
	// the amount corresponding to this user is withdrawn.
	//
	// The two types of updates can be differentiated using the status field,
	// which is "open" or "final" for a regular update and "closed" for a closing update.
	//
	ChUpdateNotif struct {
		// UpdateID denotes the unique ID for this update. It is derived from the channel ID and version number.
		UpdateID       string
		CurrChInfo     ChInfo
		ProposedChInfo ChInfo

		Type ChUpdateType

		// It is with reference to the system clock on the computer running the perun-node.
		// Time (in unix timestamp) before which response to this notification should be sent.
		//
		// It is 0, when no response is expected.
		Expiry int64

		// Error represents any error encountered while processing incoming updates or
		// while a channel is closed by the watcher..
		// When this is non empty, expiry will also be zero and no response is expected
		Error string
	}

	// App represents the app definition and the corresponding app data for a channel.
	App struct {
		Def  pchannel.App
		Data pchannel.Data
	}

	// ChInfo represents the info regarding a channel that will be sent to the user.
	ChInfo struct {
		ChID string
		// Represents the amount held by each participant in the channel.
		BalInfo BalInfo
		// App used in the channel.
		App App
		// Current Version Number for the channel. This will be zero when a channel is opened and will be incremented
		// during each update. When registering the state on-chain, if different participants register states with
		// different versions, channel will be settled according to the state with highest version number.
		Version string
	}

	// BalInfo represents the Balance information of the channel participants.
	// A valid BalInfo should meet the following conditions (will be validated before using the struct):
	//	1. Lengths of Parts list and Balance list are equal.
	//	2. All entries in Parts list are unique.
	//	3. Parts list has an entry "self", that represents the user of the session.
	//	4. No amount in Balance must be negative.
	BalInfo struct {
		Currency string   // Currency interpreter used to interpret the amounts in the balance.
		Parts    []string // List of aliases of channel participants.
		Bal      []string // Amounts held by each participant in this channel for the given currency.
	}

	// StateUpdater function is the function that will be used for applying state updates.
	StateUpdater func(*pchannel.State) error
) // nolint:gofumpt // unknown error, maybe a false positive

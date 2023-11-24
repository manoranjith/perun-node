package peruntcp

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hyperledger-labs/perun-node/api/grpc/pb"
	"github.com/plgd-dev/go-coap/v3/message"
	"github.com/plgd-dev/go-coap/v3/tcp"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	ethchanneltest "perun.network/go-perun/backend/ethereum/channel/test"
	pchannel "perun.network/go-perun/channel"
	channeltest "perun.network/go-perun/channel/test"
	pkgtest "polycry.pt/poly-go/test"
)

func TestServerResponse(t *testing.T) {
	go ServeFundingWatchingAPI(nil, "5688")

	time.Sleep(1 * time.Second)
	co, err := tcp.Dial("localhost:5688")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	ctx := context.Background()

	path := "/a"
	contentFormat := message.AppOctets
	// payload := bytes.NewReader([]byte("Hello, World!"))
	fundingReq := NewRandomFundReq(t)
	fmt.Println(
		"k",
		fundingReq.SessionID,
		fundingReq.Params.Id,
		fundingReq.Params.ChallengeDuration,
		fundingReq.Params.Parts,
		fundingReq.Params.App,
		fundingReq.Params.Nonce,
		fundingReq.Params.LedgerChannel,
		fundingReq.Params.VirtualChannel,
		fundingReq.State.Id,
		fundingReq.State.Version,
		fundingReq.State.App,
		fundingReq.State.Allocation,
		fundingReq.State.Data,
		fundingReq.State.IsFinal,
	)

	require.NoError(t, err)
	payload, err := proto.Marshal(fundingReq)
	require.NoError(t, err)

	resp, err := co.Post(ctx, path, contentFormat, bytes.NewReader(payload))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	log.Printf("Response payload: %v", resp.String())
	body, err := resp.ReadBody()
	if err != nil {
		log.Printf("Response payload error : %v", err)
	}
	log.Printf("Response payload body: %s", body)
	log.Printf("Response payload code: %s", resp.Code().String())
	log.Printf("Response payload token: %s", resp.Token())
	log.Printf("Response payload type: %s", resp.Type().String())
}

const TxFinalityDepth = 3 // For tests that use a constant finality depth.
const blockInterval = 50 * time.Millisecond

func NewRandomFundReq(t *testing.T) *pb.FundReq {
	rng := pkgtest.Prng(t)

	s := ethchanneltest.NewSetup(t, rng, 1, blockInterval, TxFinalityDepth)
	params, state := channeltest.NewRandomParamsAndState(
		rng,
		channeltest.WithChallengeDuration(uint64(100*time.Second)),
		channeltest.WithParts(s.Parts...),
		channeltest.WithAssets(s.Asset),
		channeltest.WithIsFinal(false),
		channeltest.WithLedgerChannel(true),
		channeltest.WithVirtualChannel(false),
	)
	fundingReq := pchannel.FundingReq{
		Params:    params,
		State:     state,
		Idx:       0,
		Agreement: channeltest.NewRandomBalances(rng),
	}

	req, err := pb.FromFundingReq(fundingReq)
	require.NoError(t, err)
	return req
}

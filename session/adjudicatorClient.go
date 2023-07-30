// Copyright (c) 2023 - for information on the respective copyright owner
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

package session

import (
	"context"
	"sync"

	"github.com/hyperledger-labs/perun-node"
	"github.com/hyperledger-labs/perun-node/api/grpc/pb"
	"github.com/pkg/errors"
	"perun.network/go-perun/channel"
	pchannel "perun.network/go-perun/channel"
	"perun.network/go-perun/watcher"
)

type grpcAdjudicator struct {
	apiKey string
	client pb.Funding_APIClient
}

func (a *grpcAdjudicator) Register(
	_ context.Context,
	adjReq pchannel.AdjudicatorReq,
	signedStates []pchannel.SignedState,
) (err error) {
	protoReq := pb.RegisterReq{}
	protoReq.AdjReq, err = pb.FromAdjReq(adjReq)
	if err != nil {
		err = errors.WithMessage(err, "parsing grpc adjudicator request")
		return perun.NewAPIErrUnknownInternal(err)
	}
	protoReq.SignedStates = make([]*pb.SignedState, len(signedStates))
	for i := range signedStates {
		protoReq.SignedStates[i], err = pb.FromSignedState(&signedStates[i])
		if err != nil {
			err = errors.WithMessagef(err, "parsing %d'th signed state", i)
			return perun.NewAPIErrUnknownInternal(err)
		}
	}
	protoReq.SessionID = a.apiKey

	resp, err := a.client.Register(context.Background(), &protoReq)
	if err != nil {
		err = errors.WithMessage(err, "sending the funding request")
		return perun.NewAPIErrUnknownInternal(err)
	}
	if resp.Error != nil && resp.Error.Message != "" {
		// TODO: Proper error handling
		err = errors.WithMessage(err, "registering the channel the channel")
		return perun.NewAPIErrUnknownInternal(errors.New(resp.Error.Message))
	}
	return nil
}

func (a *grpcAdjudicator) Withdraw(
	ctx context.Context,
	adjReq pchannel.AdjudicatorReq,
	stateMap pchannel.StateMap,
) (err error) {
	protoReq := pb.WithdrawReq{}
	protoReq.AdjReq, err = pb.FromAdjReq(adjReq)
	if err != nil {
		err = errors.WithMessage(err, "parsing grpc adjudicator request")
		return perun.NewAPIErrUnknownInternal(err)
	}
	protoReq.StateMap, err = pb.FromStateMap(stateMap)
	if err != nil {
		err = errors.WithMessage(err, "parsing grpc adjudicator request")
		return perun.NewAPIErrUnknownInternal(err)
	}
	protoReq.SessionID = a.apiKey

	resp, err := a.client.Withdraw(context.Background(), &protoReq)
	if err != nil {
		err = errors.WithMessage(err, "sending the funding request")
		return perun.NewAPIErrUnknownInternal(err)
	}
	if resp.Error != nil && resp.Error.Message != "" {
		// TODO: Proper error handling
		err = errors.WithMessage(err, "withdrawing the channel the channel")
		return perun.NewAPIErrUnknownInternal(errors.New(resp.Error.Message))
	}
	return nil
}

func (a *grpcAdjudicator) Progress(
	ctx context.Context,
	progReq pchannel.ProgressReq,
) (err error) {
	protoReq := pb.ProgressReq{}
	protoReq.AdjReq, err = pb.FromAdjReq(progReq.AdjudicatorReq)
	if err != nil {
		err = errors.WithMessage(err, "parsing grpc adjudicator request")
		return perun.NewAPIErrUnknownInternal(err)
	}
	protoReq.NewState, err = pb.FromState(progReq.NewState)
	if err != nil {
		err = errors.WithMessage(err, "parsing grpc adjudicator request")
		return perun.NewAPIErrUnknownInternal(err)
	}
	protoReq.Sig = progReq.Sig
	protoReq.SessionID = a.apiKey

	resp, err := a.client.Progress(context.Background(), &protoReq)
	if err != nil {
		err = errors.WithMessage(err, "sending the funding request")
		return perun.NewAPIErrUnknownInternal(err)
	}
	if resp.Error != nil && resp.Error.Message != "" {
		// TODO: Proper error handling
		err = errors.WithMessage(err, "progressing the channel the channel")
		return perun.NewAPIErrUnknownInternal(errors.New(resp.Error.Message))
	}
	return nil
}

func (a *grpcAdjudicator) Subscribe(
	ctx context.Context,
	chID pchannel.ID,
) (pchannel.AdjudicatorSubscription, error) {
	adjSubReq := &pb.SubscribeReq{
		SessionID: a.apiKey,
		ChID:      chID[:],
	}
	stream, err := a.client.Subscribe(ctx, adjSubReq)
	if err != nil {
		return nil, err
	}
	adjSubRelay := newAdjudicatorEventsSub(chID, a)
	func() {
		subscribeResp, err := stream.Recv()
		if err != nil {
			return
		}
		adjEvent, err := pb.SubscribeResponseToAdjEvent(subscribeResp)
		if err != nil {
			return
		}
		adjSubRelay.publish(adjEvent)

	}()
	return adjSubRelay, nil
}

var _ pchannel.AdjudicatorSubscription = &adjudicatorSub{}

type (
	adjudicatorSub struct {
		isOpen bool
		once   sync.Mutex

		pipe    chan channel.AdjudicatorEvent
		chID    pchannel.ID
		grpcAdj *grpcAdjudicator
	}
)

func newAdjudicatorEventsSub(chID pchannel.ID, grpcAdj *grpcAdjudicator) *adjudicatorSub {
	return &adjudicatorSub{
		isOpen:  true,
		pipe:    make(chan channel.AdjudicatorEvent, adjPubSubBufferSize),
		chID:    chID,
		grpcAdj: grpcAdj,
	}
}

// publish publishes the given adjudicator event to the subscriber.
//
// Panics if the pub-sub instance is already closed. It is implemented this
// way, because
//  1. The watcher will publish on this pub-sub only when it receives an
//     adjudicator event from the blockchain.
//  2. When de-registering a channel from the watcher, watcher will close the
//     subscription for adjudicator events from blockchain, before closing this
//     pub-sub.
//  3. This way, it can be guaranteed that, this method will never be called
//     after the pub-sub instance is closed.
func (a *adjudicatorSub) publish(e channel.AdjudicatorEvent) {
	a.once.Lock()
	a.pipe <- e
	a.once.Unlock()
}

// close closes the publisher instance and the associated subscription. Any
// further call to publish, after a pub-sub is closed will panic.
func (a *adjudicatorSub) Close() error {
	a.once.Lock()
	defer a.once.Unlock()

	unsubReq := &pb.UnsubscribeReq{
		SessionID: a.grpcAdj.apiKey,
		ChID:      a.chID[:],
	}
	a.grpcAdj.client.Unsubscribe(context.Background(), unsubReq)
	// TODO: Handle error.

	return nil
}

// EventStream returns a channel for consuming the published adjudicator
// events. It always returns the same channel and does not support
// multiplexing.
//
// The channel will be closed when the pub-sub instance is closed and Err
// should tell the possible error.
func (a *adjudicatorSub) Next() pchannel.AdjudicatorEvent {
	return <-a.pipe
}

// Err always returns nil. Because, there will be no errors when closing a
// local subscription.
func (a *adjudicatorSub) Err() error {
	return nil // Getting an error is not implemented.
}

var _ watcher.AdjudicatorSub = &adjudicatorPubSub{}

const adjPubSubBufferSize = 10

type (
	adjudicatorPubSub struct {
		once sync.Once
		pipe chan channel.AdjudicatorEvent
	}

	// adjudicatorPub is used by the watcher to publish the adjudicator events
	// received from the blockchain to the client.
	adjudicatorPub interface {
		publish(channel.AdjudicatorEvent)
		close()
	}
)

func newAdjudicatorEventsPubSub() *adjudicatorPubSub {
	return &adjudicatorPubSub{
		pipe: make(chan channel.AdjudicatorEvent, adjPubSubBufferSize),
	}
}

// publish publishes the given adjudicator event to the subscriber.
//
// Panics if the pub-sub instance is already closed. It is implemented this
// way, because
//  1. The watcher will publish on this pub-sub only when it receives an
//     adjudicator event from the blockchain.
//  2. When de-registering a channel from the watcher, watcher will close the
//     subscription for adjudicator events from blockchain, before closing this
//     pub-sub.
//  3. This way, it can be guaranteed that, this method will never be called
//     after the pub-sub instance is closed.
func (a *adjudicatorPubSub) publish(e channel.AdjudicatorEvent) {
	a.pipe <- e
}

// close closes the publisher instance and the associated subscription. Any
// further call to publish, after a pub-sub is closed will panic.
func (a *adjudicatorPubSub) close() {
	a.once.Do(func() { close(a.pipe) })
}

// EventStream returns a channel for consuming the published adjudicator
// events. It always returns the same channel and does not support
// multiplexing.
//
// The channel will be closed when the pub-sub instance is closed and Err
// should tell the possible error.
func (a *adjudicatorPubSub) EventStream() <-chan channel.AdjudicatorEvent {
	return a.pipe
}

// Err always returns nil. Because, there will be no errors when closing a
// local subscription.
func (a *adjudicatorPubSub) Err() error {
	return nil
}

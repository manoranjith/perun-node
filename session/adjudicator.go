// Copyright (c) 2022 - for information on the respective copyright owner
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

	"github.com/hyperledger-labs/perun-node"
	pchannel "perun.network/go-perun/channel"
	pwallet "perun.network/go-perun/wallet"
)

func (s *Session) Register(
	ctx context.Context,
	adjReq perun.AdjudicatorReq,
	signedStates []pchannel.SignedState,
) perun.APIError {
	s.WithField("method", "Register").Infof("\nReceived request with params %+v, %+v", adjReq, signedStates)

	pAdjReq, err := toPChannelAdjudicatorReq(adjReq, s.user.OffChain.Wallet)
	if err != nil {
		apiErr := perun.NewAPIErrUnknownInternal(err)
		s.WithFields(perun.APIErrAsMap("Register", apiErr)).Error(apiErr.Message())
		return apiErr
	}

	err = s.adjudicator.Register(ctx, pAdjReq, signedStates)
	if err != nil {
		apiErr := perun.NewAPIErrUnknownInternal(err)
		s.WithFields(perun.APIErrAsMap("Register", apiErr)).Error(apiErr.Message())
		return apiErr
	}
	s.WithField("method", "Register").Infof("Registered successfully: %+v ", adjReq.Params.ID())
	return nil
}

func (s *Session) Withdraw(
	ctx context.Context,
	adjReq perun.AdjudicatorReq,
	stateMap pchannel.StateMap,
) perun.APIError {
	s.WithField("method", "Withdraw").Infof("\nReceived request with params %+v, %+v", adjReq, stateMap)

	pAdjReq, err := toPChannelAdjudicatorReq(adjReq, s.user.OffChain.Wallet)
	if err != nil {
		apiErr := perun.NewAPIErrUnknownInternal(err)
		s.WithFields(perun.APIErrAsMap("Withdraw", apiErr)).Error(apiErr.Message())
		return apiErr
	}

	err = s.adjudicator.Withdraw(ctx, pAdjReq, stateMap)
	if err != nil {
		apiErr := perun.NewAPIErrUnknownInternal(err)
		s.WithFields(perun.APIErrAsMap("Withdraw", apiErr)).Error(apiErr.Message())
		return apiErr
	}
	s.user.OnChain.Wallet.DecrementUsage(s.user.OnChain.Addr)
	s.WithField("method", "Withdraw").Infof("Withdrawn successfully: %+v ", adjReq.Params.ID())
	return nil
}

func (s *Session) Progress(ctx context.Context, progReq perun.ProgressReq) perun.APIError {
	s.WithField("method", "Progress").Infof("\nReceived request with params %+v", progReq)

	pProgReq, err := toPChannelProgressReq(progReq, s.user.OffChain.Wallet)
	if err != nil {
		apiErr := perun.NewAPIErrUnknownInternal(err)
		s.WithFields(perun.APIErrAsMap("Progress", apiErr)).Error(apiErr.Message())
		return apiErr
	}

	err = s.adjudicator.Progress(ctx, pProgReq)
	if err != nil {
		apiErr := perun.NewAPIErrUnknownInternal(err)
		s.WithFields(perun.APIErrAsMap("Progress", apiErr)).Error(apiErr.Message())
		return apiErr
	}
	s.WithField("method", "Progress").Infof("Progressed successfully: %+v ", progReq.Params.ID())
	return nil
}

func (s *Session) Subscribe(
	ctx context.Context,
	chID pchannel.ID,
) (pchannel.AdjudicatorSubscription, perun.APIError) {
	s.WithField("method", "Subscribe").Infof("\nReceived request with params %+v", chID)

	adjSub, err := s.adjudicator.Subscribe(ctx, chID)
	if err != nil {
		apiErr := perun.NewAPIErrUnknownInternal(err)
		s.WithFields(perun.APIErrAsMap("Progress", apiErr)).Error(apiErr.Message())
		return nil, apiErr
	}
	s.WithField("method", "Subscribe").Infof("Subscribed successfully: %+v ", chID)
	return adjSub, nil
}

func toPChannelAdjudicatorReq(in perun.AdjudicatorReq, w pwallet.Wallet) (out pchannel.AdjudicatorReq, err error) {
	out.Acc, err = w.Unlock(in.Acc)
	if err != nil {
		return out, err
	}
	out.Params = in.Params
	out.Tx = in.Tx
	out.Idx = in.Idx
	out.Secondary = in.Secondary
	return out, nil
}

func toPChannelProgressReq(in perun.ProgressReq, w pwallet.Wallet) (out pchannel.ProgressReq, err error) {
	out.AdjudicatorReq, err = toPChannelAdjudicatorReq(in.AdjudicatorReq, w)
	if err != nil {
		return out, err
	}
	out.NewState = in.NewState
	out.Sig = in.Sig
	return out, nil
}

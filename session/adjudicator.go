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
)

func (s *Session) Register(
	ctx context.Context,
	adjReq pchannel.AdjudicatorReq,
	signedStates []pchannel.SignedState,
) perun.APIError {
	s.WithField("method", "Register").Infof("\nReceived request with params %+v, %+v", adjReq, signedStates)

	err := s.adjudicator.Register(ctx, adjReq, signedStates)
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
	adjReq pchannel.AdjudicatorReq,
	stateMap pchannel.StateMap,
) perun.APIError {
	s.WithField("method", "Withdraw").Infof("\nReceived request with params %+v, %+v", adjReq, stateMap)

	err := s.adjudicator.Withdraw(ctx, adjReq, stateMap)
	if err != nil {
		apiErr := perun.NewAPIErrUnknownInternal(err)
		s.WithFields(perun.APIErrAsMap("Withdraw", apiErr)).Error(apiErr.Message())
		return apiErr
	}
	s.user.OnChain.Wallet.DecrementUsage(s.user.OnChain.Addr)
	s.WithField("method", "Withdraw").Infof("Withdrawn successfully: %+v ", adjReq.Params.ID())
	return nil
}

func (s *Session) Progress(ctx context.Context, progReq pchannel.ProgressReq) perun.APIError {
	s.WithField("method", "Progress").Infof("\nReceived request with params %+v", progReq)

	err := s.adjudicator.Progress(ctx, progReq)
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

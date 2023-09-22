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

package grpc

import (
	"context"

	"github.com/pkg/errors"
	pchannel "perun.network/go-perun/channel"

	"github.com/hyperledger-labs/perun-node"
	"github.com/hyperledger-labs/perun-node/api/grpc/pb"
	"github.com/hyperledger-labs/perun-node/api/handlers"
)

// fundingServer represents a grpc server that can serve funding API.
type fundingServer struct {
	pb.UnimplementedFunding_APIServer
	*handlers.FundingHandler
}

// Fund wraps session.Fund.
func (a *fundingServer) Fund(ctx context.Context, req *pb.FundReq) (*pb.FundResp, error) {
	return a.FundingHandler.Fund(ctx, req)
}

// RegisterAssetERC20 is a stub that always returns false. Because, the remote
// funder does not support use of assets other than the default ERC20 asset.
//
// TODO: Make actual implementation.
func (a *fundingServer) RegisterAssetERC20(ctx context.Context, req *pb.RegisterAssetERC20Req) (
	*pb.RegisterAssetERC20Resp, error,
) {
	return a.FundingHandler.RegisterAssetERC20(ctx, req)
}

// IsAssetRegistered wraps session.IsAssetRegistered.
func (a *fundingServer) IsAssetRegistered(ctx context.Context, req *pb.IsAssetRegisteredReq) (
	*pb.IsAssetRegisteredResp,
	error,
) {
	return a.FundingHandler.IsAssetRegistered(ctx, req)
}

// Register wraps session.Register.
func (a *fundingServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return a.FundingHandler.Register(ctx, req)
}

// Withdraw wraps session.Withdraw.
func (a *fundingServer) Withdraw(ctx context.Context, req *pb.WithdrawReq) (*pb.WithdrawResp, error) {
	return a.FundingHandler.Withdraw(ctx, req)
}

// Progress wraps session.Progress.
func (a *fundingServer) Progress(ctx context.Context, req *pb.ProgressReq) (*pb.ProgressResp, error) {
	return a.FundingHandler.Progress(ctx, req)
}

// Subscribe wraps session.Subscribe.

func (a *fundingServer) Subscribe(req *pb.SubscribeReq, stream pb.Funding_API_SubscribeServer) error {
	sess, err := a.N.GetSession(req.SessionID)
	if err != nil {
		return errors.WithMessage(err, "retrieving session")
	}

	var chID pchannel.ID
	copy(chID[:], req.ChID)

	adjSub, err := sess.Subscribe(context.Background(), chID)
	if err != nil {
		return errors.WithMessage(err, "setting up subscription")
	}

	a.Lock()
	if a.Subscribes[req.SessionID] == nil {
		a.Subscribes[req.SessionID] = make(map[pchannel.ID]pchannel.AdjudicatorSubscription)
	}
	a.Subscribes[req.SessionID][chID] = adjSub
	a.Unlock()

	// This stream is anyways closed when StopWatching is called for.
	// Hence, that will act as the exit condition for the loop.
	go func() {
		// will return nil, when the sub is closed.
		// so, we need a mechanism to call close on the server side.
		// so, add a call Unsubscribe, which simply calls close.
		for {
			adjEvent := adjSub.Next()
			if adjEvent == nil {
				err := errors.WithMessage(adjSub.Err(), "sub closed with error")
				notif := &pb.SubscribeResp_Error{
					Error: pb.FromError(perun.NewAPIErrUnknownInternal(err)),
				}
				// TODO: Proper error handling. For now, ignore this error.
				_ = stream.Send(&pb.SubscribeResp{Response: notif}) //nolint: errcheck
				return
			}
			notif, err := pb.SubscribeResponseFromAdjEvent(adjEvent)
			if err != nil {
				return
			}
			err = stream.Send(notif)
			if err != nil {
				return
			}
		}
	}()

	return nil
}

func (a *fundingServer) Unsubscribe(_ context.Context, req *pb.UnsubscribeReq) (*pb.UnsubscribeResp, error) {
	errResponse := func(err perun.APIError) *pb.UnsubscribeResp {
		return &pb.UnsubscribeResp{
			Error: pb.FromError(err),
		}
	}

	var chID pchannel.ID
	copy(chID[:], req.ChID)

	a.Lock()
	if _, ok := a.Subscribes[req.SessionID]; !ok {
		return errResponse(perun.NewAPIErrUnknownInternal(errors.New("unknown session id"))), nil
	}
	adjSub, ok := a.Subscribes[req.SessionID][chID]
	if !ok {
		return errResponse(perun.NewAPIErrUnknownInternal(errors.New("unknown channel id"))), nil
	}
	delete(a.Subscribes[req.SessionID], chID)
	a.Unlock()

	if err := adjSub.Close(); err != nil {
		return errResponse(perun.NewAPIErrUnknownInternal(errors.WithMessage(err, "closing sub"))), nil
	}

	return &pb.UnsubscribeResp{
		Error: nil,
	}, nil
}

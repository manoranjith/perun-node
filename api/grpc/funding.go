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

	pchannel "perun.network/go-perun/channel"

	"github.com/hyperledger-labs/perun-node"
	"github.com/hyperledger-labs/perun-node/api/grpc/pb"
)

// fundingServer represents a grpc server that can serve funding API.
type fundingServer struct {
	pb.UnimplementedFunding_APIServer
	n perun.NodeAPI
}

// Fund wraps session.Fund.
func (a *fundingServer) Fund(ctx context.Context, grpcReq *pb.FundReq) (*pb.FundResp, error) {
	errResponse := func(err perun.APIError) *pb.FundResp {
		return &pb.FundResp{
			Error: toGrpcError(err),
		}
	}

	sess, apiErr := a.n.GetSession(grpcReq.SessionID)
	if apiErr != nil {
		return errResponse(apiErr), nil
	}
	req, err := fromGrpcFundingReq(grpcReq)
	if err != nil {
		return errResponse(perun.NewAPIErrUnknownInternal(err)), nil
	}

	err = sess.Fund(ctx, req)
	if err != nil {
		return errResponse(perun.NewAPIErrUnknownInternal(err)), nil
	}

	return &pb.FundResp{
		Error: nil,
	}, nil
}

func fromGrpcFundingReq(protoReq *pb.FundReq) (req pchannel.FundingReq, err error) {
	if req.Params, err = fromGrpcParams(protoReq.Params); err != nil {
		return req, err
	}
	if req.State, err = fromGrpcState(protoReq.State); err != nil {
		return req, err
	}

	req.Idx = pchannel.Index(protoReq.Idx)
	req.Agreement = fromGrpcBalances(protoReq.Agreement.Balances)
	return req, nil
}

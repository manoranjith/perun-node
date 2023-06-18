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
	"fmt"
	"math"
	"math/big"

	"github.com/pkg/errors"
	pchannel "perun.network/go-perun/channel"
	pwallet "perun.network/go-perun/wallet"

	"github.com/hyperledger-labs/perun-node/api/grpc/pb"
)

func fromGrpcParams(protoParams *pb.Params) (*pchannel.Params, error) {
	app, err := toApp(protoParams.App)
	if err != nil {
		return nil, err
	}
	parts, err := toWalletAddrs(protoParams.Parts)
	if err != nil {
		return nil, errors.WithMessage(err, "parts")
	}
	params := pchannel.NewParamsUnsafe(
		protoParams.ChallengeDuration,
		parts,
		app,
		(new(big.Int)).SetBytes(protoParams.Nonce),
		protoParams.LedgerChannel,
		protoParams.VirtualChannel)

	return params, nil
}

func toApp(protoApp []byte) (app pchannel.App, err error) {
	if len(protoApp) == 0 {
		app = pchannel.NoApp()
		return app, nil
	}
	appDef := pwallet.NewAddress()
	err = appDef.UnmarshalBinary(protoApp)
	if err != nil {
		return app, err
	}
	app, err = pchannel.Resolve(appDef)
	return app, err
}

func toWalletAddrs(protoAddrs [][]byte) ([]pwallet.Address, error) {
	addrs := make([]pwallet.Address, len(protoAddrs))
	for i := range protoAddrs {
		addrs[i] = pwallet.NewAddress()
		err := addrs[i].UnmarshalBinary(protoAddrs[i])
		if err != nil {
			return nil, errors.WithMessagef(err, "%d'th address", i)
		}
	}
	return addrs, nil
}

func fromGrpcState(protoState *pb.State) (state *pchannel.State, err error) {
	state = &pchannel.State{}
	copy(state.ID[:], protoState.Id)
	state.Version = protoState.Version
	state.IsFinal = protoState.IsFinal
	allocation, err := toAllocation(protoState.Allocation)
	if err != nil {
		return nil, errors.WithMessage(err, "allocation")
	}
	state.Allocation = *allocation
	state.App, state.Data, err = toAppAndData(protoState.App, protoState.Data)
	return state, err
}

func toAllocation(protoAlloc *pb.Allocation) (alloc *pchannel.Allocation, err error) {
	alloc = &pchannel.Allocation{}
	alloc.Assets = make([]pchannel.Asset, len(protoAlloc.Assets))
	for i := range protoAlloc.Assets {
		alloc.Assets[i] = pchannel.NewAsset()
		err = alloc.Assets[i].UnmarshalBinary(protoAlloc.Assets[i])
		if err != nil {
			return nil, errors.WithMessagef(err, "%d'th asset", i)
		}
	}
	alloc.Locked = make([]pchannel.SubAlloc, len(protoAlloc.Locked))
	for i := range protoAlloc.Locked {
		alloc.Locked[i], err = fromGrpcSubAlloc(protoAlloc.Locked[i])
		if err != nil {
			return nil, errors.WithMessagef(err, "%d'th sub alloc", i)
		}
	}
	alloc.Balances = fromGrpcBalances(protoAlloc.Balances.Balances)
	return alloc, nil
}

func toAppAndData(protoApp, protoData []byte) (app pchannel.App, data pchannel.Data, err error) {
	if len(protoApp) == 0 {
		app = pchannel.NoApp()
		data = pchannel.NoData()
		return app, data, nil
	}
	appDef := pwallet.NewAddress()
	err = appDef.UnmarshalBinary(protoApp)
	if err != nil {
		return nil, nil, err
	}
	app, err = pchannel.Resolve(appDef)
	if err != nil {
		return
	}
	data = app.NewData()
	return app, data, data.UnmarshalBinary(protoData)
}

func fromGrpcSubAlloc(protoSubAlloc *pb.SubAlloc) (subAlloc pchannel.SubAlloc, err error) {
	subAlloc = pchannel.SubAlloc{}
	subAlloc.Bals = fromGrpcBalance(protoSubAlloc.Bals)
	if len(protoSubAlloc.Id) != len(subAlloc.ID) {
		return subAlloc, errors.New("sub alloc id has incorrect length")
	}
	copy(subAlloc.ID[:], protoSubAlloc.Id)
	subAlloc.IndexMap, err = fromGrpcIndexMap(protoSubAlloc.IndexMap.IndexMap)
	return subAlloc, err
}

func fromGrpcBalances(protoBalances []*pb.Balance) [][]*big.Int {
	balances := make([][]*big.Int, len(protoBalances))
	for i, protoBalance := range protoBalances {
		balances[i] = make([]*big.Int, len(protoBalance.Balance))
		for j := range protoBalance.Balance {
			balances[i][j] = (&big.Int{}).SetBytes(protoBalance.Balance[j])
			balances[i][j].SetBytes(protoBalance.Balance[j])
		}
	}
	return balances
}

func fromGrpcBalance(protoBalance *pb.Balance) (balance []pchannel.Bal) {
	balance = make([]pchannel.Bal, len(protoBalance.Balance))
	for j := range protoBalance.Balance {
		balance[j] = new(big.Int).SetBytes(protoBalance.Balance[j])
	}
	return balance
}

func fromGrpcIndexMap(protoIndexMap []uint32) (indexMap []pchannel.Index, err error) {
	indexMap = make([]pchannel.Index, len(protoIndexMap))
	for i := range protoIndexMap {
		if protoIndexMap[i] > math.MaxUint16 {
			//nolint:goerr113  // We do not want to define this as constant error.
			return nil, fmt.Errorf("%d'th index is invalid", i)
		}
		indexMap[i] = pchannel.Index(uint16(protoIndexMap[i]))
	}
	return indexMap, nil
}

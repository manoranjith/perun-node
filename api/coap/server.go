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

package peruncoap

import (
	coap "github.com/plgd-dev/go-coap/v3"
	"github.com/plgd-dev/go-coap/v3/mux"
	pchannel "perun.network/go-perun/channel"

	"github.com/hyperledger-labs/perun-node"
	"github.com/hyperledger-labs/perun-node/api/handlers"
)

// ServeFundingWatchingAPI starts a payment channel API server that listens for incoming grpc
// requests at the specified address and serves those requests using the node API instance.
func ServeFundingWatchingAPI(n perun.NodeAPI, port string) error {
	fundingServer := &fundingServer{
		FundingHandler: &handlers.FundingHandler{
			N:          n,
			Subscribes: make(map[string]map[pchannel.ID]pchannel.AdjudicatorSubscription),
		},
	}

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Handle("/fund", mux.HandlerFunc(fundingServer.Fund))
	r.Handle("/register", mux.HandlerFunc(fundingServer.Register))
	r.Handle("/progress", mux.HandlerFunc(fundingServer.Progress))
	r.Handle("/withdraw", mux.HandlerFunc(fundingServer.Withdraw))

	return coap.ListenAndServe("tcp", ":"+port, r)
}

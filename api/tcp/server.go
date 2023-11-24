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

package peruntcp

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	pchannel "perun.network/go-perun/channel"
	"perun.network/go-perun/log"
	"polycry.pt/poly-go/sync"

	"github.com/hyperledger-labs/perun-node"
	"github.com/hyperledger-labs/perun-node/api/grpc/pb"
	"github.com/hyperledger-labs/perun-node/api/handlers"
)

type Server struct {
	sync.Closer

	server net.Listener

	fundingServer *fundingServer
}

// ServeFundingWatchingAPI starts a payment channel API server that listens for incoming grpc
// requests at the specified address and serves those requests using the node API instance.
func ServeFundingWatchingAPI(n perun.NodeAPI, port string) error {
	fundingServer := &fundingServer{
		FundingHandler: &handlers.FundingHandler{
			N:          n,
			Subscribes: make(map[string]map[pchannel.ID]pchannel.AdjudicatorSubscription),
		},
	}

	tcpServer, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("listener: %w", err)
	}

	s := &Server{
		server:        tcpServer,
		fundingServer: fundingServer,
	}
	s.OnCloseAlways(func() { tcpServer.Close() })

	for {
		conn, err := s.server.Accept()
		if err != nil {
			return err
		}

		go s.Handle(conn)
	}
}

func (s *Server) Handle(conn io.ReadWriteCloser) {
	defer conn.Close()
	s.OnCloseAlways(func() { conn.Close() })

	// var m sync.Mutex

	for {
		msg, err := recvMsg(conn)
		if err != nil {
			log.Errorf("decoding message failed: %v", err)
			return
		}

		go func() {
			switch msg := msg.GetMsg().(type) {
			case *pb.Message_FundingRequest:
				log.Warn("Server: Got Funding request")
				req, err := ParseFundingRequestMsg(msg.FundingRequest)
				if err != nil {
					log.Errorf("Invalid update message: %v", err)
					return
				}
				if err := s.funder.Fund(s.Ctx(), pchannel.FundingReq{
					Params:    &req.Params,
					State:     &req.InitialState,
					Idx:       req.Participant,
					Agreement: req.FundingAgreement,
				}); err != nil {
					log.Errorf("Funding failed: %v", err)
				}
				sendMsg(&m, conn, &pb.Message{Msg: &pb.Message_FundingResponse{
					FundingResponse: &pb.FundingResponseMsg{
						ChannelId: req.InitialState.ID[:],
						Success:   err == nil}}})
			}
		}()
	}
}

func recvMsg(conn io.Reader) (*pb.Fndi, error) {
	var size uint16
	if err := binary.Read(conn, binary.BigEndian, &size); err != nil {
		return nil, fmt.Errorf("reading size of data from wire: %w", err)
	}
	data := make([]byte, size)
	if _, err := io.ReadFull(conn, data); err != nil {
		return nil, fmt.Errorf("reading data from wire: %w", err)
	}
	var msg pb.Message
	if err := protobuf.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("unmarshalling message: %w", err)
	}
	return &msg, nil
}

func sendMsg(m *sync.Mutex, conn io.Writer, msg *pb.Message) error {
	m.Lock()
	defer m.Unlock()
	data, err := protobuf.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshalling message: %w", err)
	}
	if err := binary.Write(conn, binary.BigEndian, uint16(len(data))); err != nil {
		return fmt.Errorf("writing length to wire: %w", err)
	}
	if _, err = conn.Write(data); err != nil {
		return fmt.Errorf("writing data to wire: %w", err)
	}
	return nil
}

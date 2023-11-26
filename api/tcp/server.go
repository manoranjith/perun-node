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
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"google.golang.org/protobuf/proto"
	pchannel "perun.network/go-perun/channel"
	"perun.network/go-perun/log"
	"polycry.pt/poly-go/sync"

	"github.com/hyperledger-labs/perun-node"
	"github.com/hyperledger-labs/perun-node/api/grpc/pb"
	"github.com/hyperledger-labs/perun-node/api/handlers"
	"github.com/hyperledger-labs/perun-node/app/payment"
)

type Server struct {
	sync.Closer

	server net.Listener

	fundingHandler *handlers.FundingHandler

	sessionID string // For timebeing use hard-coded session-id
}

// ServeFundingWatchingAPI starts a payment channel API server that listens for incoming grpc
// requests at the specified address and serves those requests using the node API instance.
func ServeFundingWatchingAPI(n perun.NodeAPI, port string) error {
	var err error
	sessionID, _, err := payment.OpenSession(n, "api/session.yaml")
	if err != nil {
		return err
	}

	fundingServer := &handlers.FundingHandler{
		N:          n,
		Subscribes: make(map[string]map[pchannel.ID]pchannel.AdjudicatorSubscription),
	}

	tcpServer, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("listener: %w", err)
	}

	s := &Server{
		server:         tcpServer,
		fundingHandler: fundingServer,
		sessionID:      sessionID,
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

	var m sync.Mutex

	for {
		msg, err := recvMsg(conn)
		if err != nil {
			log.Errorf("decoding message failed: %v", err)
			return
		}

		go func() {
			switch msg := msg.GetMsg().(type) {
			case *pb.APIMessage_FundReq:
				log.Warnf("Server: Got Funding request: %+v", msg)
				// TODO: error is always nil. Remove that return argument.
				msg.FundReq.SessionID = s.sessionID
				fundResp, err := s.fundingHandler.Fund(context.Background(), msg.FundReq)
				if err != nil {
					log.Errorf("fund response error +%v", err)
				}
				sendMsg(&m, conn, &pb.APIMessage{Msg: &pb.APIMessage_FundResp{
					FundResp: fundResp}})
			}
		}()
	}
}

func recvMsg(conn io.Reader) (*pb.APIMessage, error) {
	var size uint16
	if err := binary.Read(conn, binary.BigEndian, &size); err != nil {
		return nil, fmt.Errorf("reading size of data from wire: %w", err)
	}
	data := make([]byte, size)
	if _, err := io.ReadFull(conn, data); err != nil {
		return nil, fmt.Errorf("reading data from wire: %w", err)
	}
	var msg pb.APIMessage
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("unmarshalling message: %w", err)
	}
	return &msg, nil
}

func sendMsg(m *sync.Mutex, conn io.Writer, msg *pb.APIMessage) error {
	m.Lock()
	defer m.Unlock()
	data, err := proto.Marshal(msg)
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

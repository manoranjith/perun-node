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
	"bytes"
	"log"

	"github.com/plgd-dev/go-coap/v3/message"
	"github.com/plgd-dev/go-coap/v3/message/codes"
	"github.com/plgd-dev/go-coap/v3/mux"
	"google.golang.org/protobuf/proto"
)

func validateAndParseRequest(w mux.ResponseWriter, r *mux.Message, req proto.Message) bool {
	if r.Code() != codes.POST {
		setErrorResponse(w, codes.BadOption, "Only POST method is supported")
		return false
	}

	if r.Type() != message.Confirmable {
		setErrorResponse(w, codes.BadOption, "Only Confirmable message type is supported")
		return false
	}

	cf, _ := r.ContentFormat()
	if cf != message.AppOctets {
		setErrorResponse(w, codes.UnsupportedMediaType, "Only Octet stream media type is supported")
		return false
	}

	err := parseRequest(r, req)
	if err != nil {
		setErrorResponse(w, codes.BadOption, "Cannot parse message body")
		return false
	}

	return true
}

func parseRequest(r *mux.Message, req proto.Message) error {
	payload, err := r.ReadBody()
	if err != nil {
		return err
	}

	err = proto.Unmarshal(payload, req)
	return err
}

func setErrorResponse(w mux.ResponseWriter, code codes.Code, info string) {
	err := w.SetResponse(code, message.TextPlain, bytes.NewReader([]byte(info)))
	if err != nil {
		log.Printf("cannot set response: %v", err)
	}
}

func setResponse(w mux.ResponseWriter, code codes.Code, resp proto.Message) {
	payload, err := proto.Marshal(resp)
	if err != nil {
		setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	err = w.SetResponse(code, message.AppOctets, bytes.NewReader(payload))
	if err != nil {
		log.Printf("cannot set response: %v", err)
	}
}

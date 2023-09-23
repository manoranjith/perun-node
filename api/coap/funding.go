package peruncoap

import (
	"bytes"
	"context"
	"log"

	"github.com/hyperledger-labs/perun-node/api/grpc/pb"
	"github.com/hyperledger-labs/perun-node/api/handlers"
	"github.com/plgd-dev/go-coap/v3/message"
	"github.com/plgd-dev/go-coap/v3/message/codes"
	"github.com/plgd-dev/go-coap/v3/mux"
	"google.golang.org/protobuf/proto"
)

// fundingServer represents a grpc server that can serve funding API.
type fundingServer struct {
	*handlers.FundingHandler
}

func loggingMiddleware(next mux.Handler) mux.Handler {
	return mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		log.Printf("ClientAddress %v, %v\n\n", w.Conn().RemoteAddr(), r.String())
		next.ServeCOAP(w, r)
	})
}

func (f *fundingServer) Fund(w mux.ResponseWriter, r *mux.Message) {
	req := pb.FundReq{}
	if !f.validateAndParseRequest(w, r, &req) {
		return
	}

	resp, err := f.FundingHandler.Fund(context.TODO(), &req)
	if err != nil {
		f.sendErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	f.sendResponse(w, codes.Content, resp)
}

func (f *fundingServer) Register(w mux.ResponseWriter, r *mux.Message) {
	req := pb.RegisterReq{}
	if !f.validateAndParseRequest(w, r, &req) {
		return
	}

	resp, err := f.FundingHandler.Register(context.TODO(), &req)
	if err != nil {
		f.sendErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	f.sendResponse(w, codes.Content, resp)
}

func (f *fundingServer) Withdraw(w mux.ResponseWriter, r *mux.Message) {
	req := pb.WithdrawReq{}
	if !f.validateAndParseRequest(w, r, &req) {
		return
	}

	resp, err := f.FundingHandler.Withdraw(context.TODO(), &req)
	if err != nil {
		f.sendErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	f.sendResponse(w, codes.Content, resp)
}

func (f *fundingServer) Progress(w mux.ResponseWriter, r *mux.Message) {
	req := pb.ProgressReq{}
	if !f.validateAndParseRequest(w, r, &req) {
		return
	}

	resp, err := f.FundingHandler.Progress(context.TODO(), &req)
	if err != nil {
		f.sendErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	f.sendResponse(w, codes.Content, resp)
}

func (f *fundingServer) validateAndParseRequest(w mux.ResponseWriter, r *mux.Message, req proto.Message) bool {
	if r.Code() != codes.POST {
		f.sendErrorResponse(w, codes.BadOption, "Only POST method is supported")
		return false
	}

	if r.Type() != message.Confirmable {
		f.sendErrorResponse(w, codes.BadOption, "Only Confirmable message type is supported")
		return false
	}

	cf, _ := r.ContentFormat()
	if cf != message.AppOctets {
		f.sendErrorResponse(w, codes.UnsupportedMediaType, "Only Octet stream media type is supported")
		return false
	}

	err := f.parseRequest(r, req)
	if err != nil {
		f.sendErrorResponse(w, codes.BadOption, "Cannot parse message body")
		return false
	}

	return true
}

func (f *fundingServer) parseRequest(r *mux.Message, req proto.Message) error {
	payload, err := r.ReadBody()
	if err != nil {
		return err
	}

	err = proto.Unmarshal(payload, req)
	return err
}

func (f *fundingServer) sendErrorResponse(w mux.ResponseWriter, code codes.Code, info string) {
	err := w.SetResponse(code, message.TextPlain, bytes.NewReader([]byte(info)))
	if err != nil {
		log.Printf("cannot set response: %v", err)
	}
}

func (f *fundingServer) sendResponse(w mux.ResponseWriter, code codes.Code, resp proto.Message) {
	payload, err := proto.Marshal(resp)
	if err != nil {
		f.sendErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	err = w.SetResponse(code, message.AppOctets, bytes.NewReader(payload))
	if err != nil {
		log.Printf("cannot set response: %v", err)
	}
}

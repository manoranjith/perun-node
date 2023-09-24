package peruncoap

import (
	"context"
	"log"

	"github.com/hyperledger-labs/perun-node/api/grpc/pb"
	"github.com/hyperledger-labs/perun-node/api/handlers"
	"github.com/plgd-dev/go-coap/v3/message/codes"
	"github.com/plgd-dev/go-coap/v3/mux"
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
	if !validateAndParseRequest(w, r, &req) {
		return
	}

	resp, err := f.FundingHandler.Fund(context.TODO(), &req)
	if err != nil {
		setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	setResponse(w, codes.Content, resp)
}

func (f *fundingServer) RegisterAssetERC20(w mux.ResponseWriter, r *mux.Message) {
	req := pb.RegisterAssetERC20Req{}
	if !validateAndParseRequest(w, r, &req) {
		return
	}

	resp, err := f.FundingHandler.RegisterAssetERC20(context.TODO(), &req)
	if err != nil {
		setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	setResponse(w, codes.Content, resp)
}
func (f *fundingServer) IsAssetRegistered(w mux.ResponseWriter, r *mux.Message) {
	req := pb.IsAssetRegisteredReq{}
	if !validateAndParseRequest(w, r, &req) {
		return
	}

	resp, err := f.FundingHandler.IsAssetRegistered(context.TODO(), &req)
	if err != nil {
		setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	setResponse(w, codes.Content, resp)
}
func (f *fundingServer) Register(w mux.ResponseWriter, r *mux.Message) {
	req := pb.RegisterReq{}
	if !validateAndParseRequest(w, r, &req) {
		return
	}

	resp, err := f.FundingHandler.Register(context.TODO(), &req)
	if err != nil {
		setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	setResponse(w, codes.Content, resp)
}

func (f *fundingServer) Withdraw(w mux.ResponseWriter, r *mux.Message) {
	req := pb.WithdrawReq{}
	if !validateAndParseRequest(w, r, &req) {
		return
	}

	resp, err := f.FundingHandler.Withdraw(context.TODO(), &req)
	if err != nil {
		setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	setResponse(w, codes.Content, resp)
}

func (f *fundingServer) Progress(w mux.ResponseWriter, r *mux.Message) {
	req := pb.ProgressReq{}
	if !validateAndParseRequest(w, r, &req) {
		return
	}

	resp, err := f.FundingHandler.Progress(context.TODO(), &req)
	if err != nil {
		setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	setResponse(w, codes.Content, resp)
}

package peruntcp

import (
	"context"

	"github.com/hyperledger-labs/perun-node/api/grpc/pb"
	"github.com/hyperledger-labs/perun-node/api/handlers"
)

// fundingServer represents a grpc server that can serve funding API.
type fundingServer struct {
	*handlers.FundingHandler
}

func (f *fundingServer) Fund(req *pb.FundReq) {
	resp, err := f.FundingHandler.Fund(context.TODO(), req)
	if err != nil {
		// setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
		return
	}

	_, _ = resp, err

	// setResponse(w, codes.Content, resp)
}

// func (f *fundingServer) RegisterAssetERC20(w mux.ResponseWriter, r *mux.Message) {
// 	req := pb.RegisterAssetERC20Req{}
// 	if !validateAndParseRequest(w, r, &req) {
// 		return
// 	}

// 	resp, err := f.FundingHandler.RegisterAssetERC20(context.TODO(), &req)
// 	if err != nil {
// 		setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
// 		return
// 	}

// 	setResponse(w, codes.Content, resp)
// }
// func (f *fundingServer) IsAssetRegistered(w mux.ResponseWriter, r *mux.Message) {
// 	req := pb.IsAssetRegisteredReq{}
// 	if !validateAndParseRequest(w, r, &req) {
// 		return
// 	}

// 	resp, err := f.FundingHandler.IsAssetRegistered(context.TODO(), &req)
// 	if err != nil {
// 		setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
// 		return
// 	}

// 	setResponse(w, codes.Content, resp)
// }
// func (f *fundingServer) Register(w mux.ResponseWriter, r *mux.Message) {
// 	req := pb.RegisterReq{}
// 	if !validateAndParseRequest(w, r, &req) {
// 		return
// 	}

// 	resp, err := f.FundingHandler.Register(context.TODO(), &req)
// 	if err != nil {
// 		setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
// 		return
// 	}

// 	setResponse(w, codes.Content, resp)
// }

// func (f *fundingServer) Withdraw(w mux.ResponseWriter, r *mux.Message) {
// 	req := pb.WithdrawReq{}
// 	if !validateAndParseRequest(w, r, &req) {
// 		return
// 	}

// 	resp, err := f.FundingHandler.Withdraw(context.TODO(), &req)
// 	if err != nil {
// 		setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
// 		return
// 	}

// 	setResponse(w, codes.Content, resp)
// }

// func (f *fundingServer) Progress(w mux.ResponseWriter, r *mux.Message) {
// 	req := pb.ProgressReq{}
// 	if !validateAndParseRequest(w, r, &req) {
// 		return
// 	}

// 	resp, err := f.FundingHandler.Progress(context.TODO(), &req)
// 	if err != nil {
// 		setErrorResponse(w, codes.InternalServerError, "Cannot parse the response")
// 		return
// 	}

// 	setResponse(w, codes.Content, resp)
// }

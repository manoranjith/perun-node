package peruncoap

import (
	"bytes"
	"fmt"
	"log"

	"github.com/hyperledger-labs/perun-node/api/grpc/pb"
	coap "github.com/plgd-dev/go-coap/v3"
	"github.com/plgd-dev/go-coap/v3/message"
	"github.com/plgd-dev/go-coap/v3/message/codes"
	"github.com/plgd-dev/go-coap/v3/mux"
	"google.golang.org/protobuf/proto"
)

func loggingMiddleware(next mux.Handler) mux.Handler {
	return mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		log.Printf("ClientAddress %v, %v\n\n", w.Conn().RemoteAddr(), r.String())
		next.ServeCOAP(w, r)
	})
}

func handleA(w mux.ResponseWriter, r *mux.Message) {
	requestCode := r.Code().String()
	requestType := r.Type().String()
	requestToken := r.Token().String()
	requestMessageID := r.MessageID()

	contentFormat, err := r.ContentFormat()
	if err != nil {
		// handler error
	}
	payload, err := r.ReadBody()
	if err != nil {
		// handler error
	}
	protoFundingReq := pb.FundReq{}
	err = proto.Unmarshal(payload, &protoFundingReq)
	if err != nil {
		// handler error
	}
	fmt.Println(
		"g",
		protoFundingReq.SessionID,
		protoFundingReq.Params.Id,
		protoFundingReq.Params.ChallengeDuration,
		protoFundingReq.Params.Parts,
		protoFundingReq.Params.App,
		protoFundingReq.Params.Nonce,
		protoFundingReq.Params.LedgerChannel,
		protoFundingReq.Params.VirtualChannel,
		protoFundingReq.State.Id,
		protoFundingReq.State.Version,
		protoFundingReq.State.App,
		protoFundingReq.State.Allocation,
		protoFundingReq.State.Data,
		protoFundingReq.State.IsFinal,
	)
	fundingReq, err := pb.ToFundingReq(&protoFundingReq)
	if err != nil {
		// handler error
	}
	fmt.Println(
		requestCode,
		requestType,
		requestToken,
		requestMessageID,
		contentFormat.String(),
		fundingReq)
	err = w.SetResponse(codes.Content, message.TextPlain, bytes.NewReader([]byte("hello world")))
	if err != nil {
		log.Printf("cannot set response: %v", err)
	}
}

func handleB(w mux.ResponseWriter, r *mux.Message) {
	fmt.Println("BBB")
	customResp := w.Conn().AcquireMessage(r.Context())
	defer w.Conn().ReleaseMessage(customResp)
	customResp.SetCode(codes.Content)
	customResp.SetToken(r.Token())
	customResp.SetContentFormat(message.TextPlain)
	customResp.SetBody(bytes.NewReader([]byte("B hello world")))
	err := w.Conn().WriteMessage(customResp)
	if err != nil {
		log.Printf("cannot set response: %v", err)
	}
}

func Serve() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Handle("/a", mux.HandlerFunc(handleA))
	r.Handle("/b", mux.HandlerFunc(handleB))

	log.Fatal(coap.ListenAndServe("udp", ":5688", r))
}

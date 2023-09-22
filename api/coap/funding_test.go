package peruncoap

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/plgd-dev/go-coap/v3/udp"
)

func TestServerResponse(t *testing.T) {
	go Serve()

	co, err := udp.Dial("localhost:5688")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	path := "/a"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := co.Get(ctx, path)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	log.Printf("Response payload: %v", resp.String())
	body, err := resp.ReadBody()
	if err != nil {
		log.Printf("Response payload error : %v", err)
	}
	log.Printf("Response payload body: %s", body)
	log.Printf("Response payload code: %s", resp.Code().String())
	log.Printf("Response payload token: %s", resp.Token())
	log.Printf("Response payload type: %s", resp.Type().String())
}

package peruncoap

import (
	"context"
	"log"
	"os"
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
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := co.Get(ctx, path)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	log.Printf("Response payload: %v", resp.String())
}

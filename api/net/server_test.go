package perunnet

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/plgd-dev/go-coap/v3/message"
	"github.com/plgd-dev/go-coap/v3/tcp"
	"github.com/plgd-dev/go-coap/v3/tcp/coder"
	"github.com/stretchr/testify/require"
)

func TestRunCoapServer(t *testing.T) {
	ServeFundingWatchingAPI(nil, "5688")
}

func TestDecode(t *testing.T) {

	go ServeFundingWatchingAPI(nil, "5688")

	co, err := tcp.Dial("localhost:5688")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	ctx := context.Background()

	path := "/fund"
	contentFormat := message.AppOctets
	payload := []byte("Hello, World!")

	msg, err := co.NewPostRequest(ctx, path, contentFormat, bytes.NewReader(payload))
	require.NoError(t, err)
	msgBytes, err := msg.MarshalWithEncoder(coder.DefaultCoder)
	require.NoError(t, err)

	fmt.Printf("\nMESSAGE:%+vEND\n-", msg)
	fmt.Printf("\nMESSAGE:%+vEND\n-", msgBytes)

	// encoded := []byte{68, 2, 0, 21, 0, 0, 57, 116, 180, 102, 117, 110, 100,
	// 	17, 42, 255, 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100,
	// 	33}
	// encoded := []byte{64, 2, 0, 0, 180, 102, 117, 110, 100, 17, 42, 255, 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	// encoded := []byte{68, 1, 93, 31, 0, 0, 57, 116, 57, 108, 111, 99, 97, 108, 104, 111, 115, 116, 131, 116, 118, 49}
	// // encoded := []byte{72, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 180, 102, 117, 110, 100, 17, 42, 255, 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}

	// msg := message.Message{}

	// msg.MessageID = 21
	// msg.Type = message.Confirmable

	// _, err := coder.DefaultCoder.Decode(encoded, &msg)
	// require.NoError(t, err)
	// fmt.Printf("\n%+v\n", msg)
}

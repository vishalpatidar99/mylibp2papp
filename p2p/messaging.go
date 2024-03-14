package p2p

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"

	"github.com/libp2p/go-libp2p/core/network"
)

type MessageRequestBody struct {
	Message string
}

func MessageStreamHandler(stream network.Stream) {
	ctx := context.Background()
	defer ctx.Done()

	r := bufio.NewReader(stream)
	str, err := r.ReadString('\n')
	if err != nil {
		fmt.Printf("failed to read from new stream buffer - %v", err)
		return
	}

	msgReq := MessageRequestBody{}
	err = json.Unmarshal([]byte(str), &msgReq)
	if err != nil {
		fmt.Printf("error on unmarshal the message %v", err)
		return
	}
	fmt.Println("Got the message hurray! - - - - - ", msgReq)
}

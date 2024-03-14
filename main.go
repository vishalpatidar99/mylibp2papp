package main

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/vishalpatidar99/mylibp2papp/p2p"
)

func main() {
	ctx := context.Background()
	remotNodeID := "QmfLKonCWTkvZEAzze3pXoMzx8QhrvdzRUgjjtQnbqZ2dN"
	remotPort := 42984

	// fmt.Println(node_id)

	host, err := p2p.RunNode()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Self Node info: ", host.ID().String())
	spDMSAddr := fmt.Sprintf("/ip4/127.0.0.1/tcp/%d/p2p/%s", remotPort, remotNodeID)

	// Parse multiaddress of the SP DMS node
	spDMSInfo, err := peer.AddrInfoFromString(spDMSAddr)
	if err != nil {
		fmt.Println("Error parsing multiaddress of the SP DMS node:", err)
		return
	}

	// Connect to the SP DMS node
	err = host.Connect(ctx, *spDMSInfo)
	if err != nil {
		fmt.Println("Error connecting to the SP DMS node:", err)
		return
	}

	fmt.Println("Connected to the SP DMS node!")

	select {}

}

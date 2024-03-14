package p2p

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/protocol"
)

const ShellProtocolID = "/nunet/dms/shell/0.0.1"

func RunNode() (host.Host, error) {
	// ctx := context.Background()
	// Create a libp2p host
	host, err := NewHost()
	if err != nil {
		return nil, fmt.Errorf("error while creating libp2p host %v", err)
	}

	// setting stream handler
	host.SetStreamHandler(protocol.ID(ShellProtocolID), MessageStreamHandler)

	return host, nil
}

func NewHost() (host.Host, error) {
	host, err := libp2p.New()
	if err != nil {
		return nil, fmt.Errorf("error creating libp2p host: %v", err)
	}

	return host, nil
}

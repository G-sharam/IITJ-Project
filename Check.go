package main

import (
	"context"
	"fmt"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	peer "github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	ctx := context.Background()

	// Create a libp2p Host with a random peer ID
	host, err := libp2p.New()
	if err != nil {
		log.Fatal(err)
	}

	// Print your Peer ID and multiaddress
	fmt.Println("Peer ID:", host.ID().Pretty())

	for _, addr := range host.Addrs() {
		fmt.Println("Listening on:", addr.Encapsulate(ma.StringCast("/p2p/" + host.ID().Pretty())))
	}

	select {} // keep running
}

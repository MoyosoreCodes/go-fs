package main

import (
	"fmt"
	"log"

	"github.com/MoyosoreCodes/go-fs/p2p"
)

func main() {
	fmt.Println("Starting P2P node...")
	
	tr := p2p.InitTransport(":4444")
	fmt.Printf("Initialized transport on port %s\n", ":4000")

	if err := tr.InitListener(); err != nil {
		log.Fatalf("TCP listen error: %s\n", err)
	}
	fmt.Println("Listening for connections...")
	
	select {}
}
package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	address string
	listener net.Listener 

	mu sync.RWMutex
	peers map[net.Addr]Peer
}

type TCPPeer struct {
	conn net.Conn
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}
}

func InitTransport(address string) *TCPTransport {
	return &TCPTransport{
		address: address,
	}
}

func (t *TCPTransport) InitListener() error {
	listener, err := net.Listen("tcp", t.address)
	if err != nil {
		fmt.Printf("TCP listen error: %s\n", err)
		return err
	}
	t.listener = listener
	

	go t.initAcceptLoop()
	
	return nil
}

func (t *TCPTransport) initAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		go t.handleConnection(conn)
	}

}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	defer conn.Close()
	peer := NewTCPPeer(conn, true)

	fmt.Printf("new incoming connection %+v\n", peer)
}

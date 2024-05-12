package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	ListenAddress string
	Listener      net.Listener
	peers         map[net.Addr]Peer
	mu            sync.RWMutex
}

func NewTcpTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		ListenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.Listener, err = net.Listen("tcp", t.ListenAddress)

	if err != nil {
		return err
	}

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()

		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Printf("new incoming connection %+v\n", conn)
}

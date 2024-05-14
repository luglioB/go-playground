package main

import (
	"fs/p2p"
	"log"
)

func main() {
	tr := p2p.NewTcpTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}

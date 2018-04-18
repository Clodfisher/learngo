package main

import (
	"github.com/Clodfisher/learngo/rpcregister"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(&rpcregister.RegisterServer{
		RegisterMap: make(map[string]string),
	})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}

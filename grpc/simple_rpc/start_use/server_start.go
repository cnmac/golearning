package main

import (
	"github.com/cnmac/golearning/grpc/simple_rpc/rpc_mod"
	"log"
	"net"
	"net/rpc"
)

func main() {
	rpc_mod.RegisterHelloService(new(rpc_mod.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}

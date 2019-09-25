package main

import (
	"fmt"
	"github.com/cnmac/golearning/grpc/simple_rpc/rpc_mod"
	"log"
)

func main() {
	client, err := rpc_mod.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.RpcHello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

package rpc_mod

import (
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	RpcHello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

func (p *HelloService) RpcHello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

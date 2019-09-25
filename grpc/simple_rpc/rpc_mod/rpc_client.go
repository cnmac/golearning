package rpc_mod

import (
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func (p *HelloServiceClient) RpcHello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".RpcHello", request, reply)
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

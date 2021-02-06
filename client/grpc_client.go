package client

import (
	"context"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"go-kit-microservice-demo/endpoint"
	"go-kit-microservice-demo/pb"
	"go-kit-microservice-demo/service"
	"google.golang.org/grpc"
)

func NewGRPCClient(conn *grpc.ClientConn) service.Service {
	rpcLoginEndpoint := kitgrpc.NewClient(conn, "pb.User", "RpcLogin", RequestLogin, ResponseLogin, pb.LoginReply{}).Endpoint()
	return &endpoint.Set{RpcLoginEndpoint:rpcLoginEndpoint}
}

func RequestLogin (ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Login)
	return &pb.Login{
		Username:             req.GetUsername(),
		Password:             req.GetPassword(),
	}, nil
}

func ResponseLogin (ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginReply)
	return &pb.LoginReply{
		Token:                resp.GetToken(),
	}, nil
}
package transport

import (
	"context"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"go-kit-microservice-demo/endpoint"
	"go-kit-microservice-demo/pb"
)

type grpcServer struct {
	login kitgrpc.Handler
}

func NewRPCServer (ep endpoint.Set)  pb.UserServer{
	//options := []kitgrpc.ServerOption{
	//	kitgrpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
	//
	//	}),
	//}
	return &grpcServer{login:kitgrpc.NewServer(
			ep.RpcLoginEndpoint,
			RequestGrpcLogin,
			ResponseGrpcLogin,
		)}
}

func (s *grpcServer) RpcLogin(ctx context.Context, req *pb.Login) (*pb.LoginReply, error) {
	_, resp , err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.LoginReply), nil
}

func RequestGrpcLogin (ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Login)
	return &pb.Login{
		Username:             req.GetUsername(),
		Password:             req.GetPassword(),
	}, nil
}

func ResponseGrpcLogin (ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginReply)
	return &pb.LoginReply{
		//Token:                resp.GetToken(),
		Token:                resp.Token,
	}, nil
}

package client

import (
	"context"
	"go-kit-microservice-demo/pb"
	"google.golang.org/grpc"
	"testing"
)

func TestNewGRPCClient(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	svc := NewGRPCClient(conn)
	reply, err := svc.RpcLogin(context.Background(), &pb.Login{
		Username:             "vj",
		Password:             "123",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(reply.Token)
}
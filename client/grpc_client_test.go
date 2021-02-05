package client

import (
	"context"
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
	token, err := svc.Login(context.Background(), "vj", "111")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(token)
}
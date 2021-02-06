package main

import (
	"fmt"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"go-kit-microservice-demo/endpoint"
	"go-kit-microservice-demo/pb"
	"go-kit-microservice-demo/service"
	"go-kit-microservice-demo/transport"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main () {
	//logger := log.NewLogfmtLogger(os.Stderr)

	server := service.NewService()
	//server = LoggingMiddleware{logger, server}
	endpoints := endpoint.NewEndpointSet(server)
	//httpHandle := transport.NewHTTPHandler(endpoints)

	grpcServer := transport.NewRPCServer(endpoints)

	grpcListen, err := net.Listen("tcp", ":9001")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))

	pb.RegisterUserServer(baseServer, grpcServer)
	if err = baseServer.Serve(grpcListen); err != nil {
		fmt.Println(err)
		os.Exit(0);
	}
	//_ = http.ListenAndServe(":8999", httpHandle)
}

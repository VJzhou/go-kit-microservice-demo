package main

import (
	"github.com/go-kit/kit/log"
	"go-kit-microservice-demo/endpoint"
	"go-kit-microservice-demo/service"
	"go-kit-microservice-demo/transport"
	"net/http"
	"os"
)

func main () {
	logger := log.NewLogfmtLogger(os.Stderr)

	server := service.NewService()
	server = LoggingMiddleware{logger, server}
	endpoints := endpoint.NewEndpointSet(server)
	httpHandle := transport.NewHTTPHandler(endpoints)

	_ = http.ListenAndServe(":8999", httpHandle)
}

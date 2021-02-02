package main

import (
	"go-kit-microservice-demo/endpoint"
	"go-kit-microservice-demo/service"
	"go-kit-microservice-demo/transport"
	"go-kit-microservice-demo/util"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
)

func main () {
	logger := log.NewLogfmtLogger(os.Stderr)

	server := service.NewService()
	server = util.LoggingMiddleware{logger, server}
	endpoints := endpoint.NewEndpointSet(server)
	httpHandle := transport.NewHTTPHandler(endpoints)

	_ = http.ListenAndServe(":8999", httpHandle)
}

package main

import (
	"awesomeProject/endpoint"
	"awesomeProject/service"
	"awesomeProject/transport"
	"awesomeProject/util"
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

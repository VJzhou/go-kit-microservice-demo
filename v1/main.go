package main

import (
	"awesomeProject/endpoint"
	"awesomeProject/service"
	"awesomeProject/transport"
	"net/http"
)

func main () {
	server := service.NewService()
	endpoints := endpoint.NewEndpointSet(server)
	httpHandle := transport.NewHTTPHandler(endpoints)

	_ = http.ListenAndServe(":8999", httpHandle)
}

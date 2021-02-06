package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"go-kit-microservice-demo/endpoint"
	"go-kit-microservice-demo/service"
	"go-kit-microservice-demo/transport"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main () {
	logger := log.NewLogfmtLogger(os.Stderr)
	errChan := make(chan  error)

	fieldKeys := []string{"method", "err"}

	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace:"test_group",
		Subsystem:"AddService",
		Name:"request_count",
		Help: "Number of requests received",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace:   "test_group",
		Subsystem:   "AddService",
		Name:        "request_latency",
		Help:        "Total duration of request in microsecond",
	}, fieldKeys)


	server := service.NewService()
	server = LoggingMiddleware{logger, server}
	server = service.Metrics(requestCount, requestLatency)(server)

	endpoints := endpoint.NewEndpointSet(server)
	httpHandle := transport.NewHTTPHandler(endpoints)

	go func() {
		fmt.Println("Listen server at 8999 port")
		errChan <- http.ListenAndServe(":8999", httpHandle)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c , syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("over %s", <-c)
	}()

	fmt.Println(<- errChan)
}

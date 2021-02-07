package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"go-kit-microservice-demo/register/endpoint"
	"go-kit-microservice-demo/register/service"
	"go-kit-microservice-demo/register/transport"
	"go-kit-microservice-demo/register/util"
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
		Subsystem:"add_service",
		Name:"request_count",
		Help: "Number of requests received",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace:   "test_group",
		Subsystem:   "add_service",
		Name:        "request_latency",
		Help:        "Total duration of request in microsecond",
	}, fieldKeys)


	server := service.NewService()
	server = service.LoggingMiddleware{logger, server}

	server = service.MetricMiddleware{Service: server, RequestCount:requestCount, RequestLatency:requestLatency}

	endpoints := endpoint.NewEndpointSet(server)
	httpHandle := transport.NewHTTPHandler(endpoints)

	// 注册服务对象
	rg := util.Register("192.168.1.203", "8500", "192.168.1.203", "8999", logger)


	go func() {
		fmt.Println("Listen server at 8999 port")
		rg.Register()
		errChan <- http.ListenAndServe(":8999", httpHandle)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c , syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("over %s", <-c)
	}()

	fmt.Println(<- errChan)
	rg.Deregister()
}

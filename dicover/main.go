package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main () {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestamp)
		logger  = log.With(logger, "caller", log.DefaultCaller)
	}

	var client consul.Client
	{
		consulConfig := api.DefaultConfig()
		consulConfig.Address = "http://127.0.0.1:8500"
		consulClient, err := api.NewClient(consulConfig)

		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}

		client = consul.NewClient(consulClient)
	}

	ctx := context.Background()

	ep := NewEndpoint(ctx, client, logger)

	r := MakeHTTPHandler(ep)

	errChan := make(chan error)

	go func() {
		logger.Log("transport", "HTTP", "addr", "9001")
		errChan <- http.ListenAndServe(":9001", r)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	logger.Log(<-errChan)
}
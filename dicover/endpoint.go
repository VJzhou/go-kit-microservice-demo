package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/lb"
	"time"
)

type Set struct {
	AddEndpoint endpoint.Endpoint
	LoginEndpoint endpoint.Endpoint
}

func NewEndpoint (ctx context.Context, client consul.Client, logger log.Logger) Set{
	return Set{
		AddEndpoint:   MakeAddDiscoveryEndpoint(ctx, client, logger),
		LoginEndpoint: MakeLoginDiscoveryEndpoint(ctx, client, logger),
	}
}

func MakeAddDiscoveryEndpoint(ctx context.Context, client consul.Client, logger log.Logger) endpoint.Endpoint {
	serviceName := "add_service"
	tags := []string{"add_service", "vj"}
	passingOnly := true

	duration := 500 * time.Millisecond

	instancer := consul.NewInstancer(client, logger, serviceName, tags, passingOnly)

	factory := addServiceFactory(ctx, "POST", "/add")
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(1, duration, balancer)

	return retry

}

func MakeLoginDiscoveryEndpoint(ctx context.Context, client consul.Client, logger log.Logger) endpoint.Endpoint {
	serviceName := "add_service"
	tags := []string{"add_service", "vj"}
	passingOnly := true

	duration := 500 * time.Millisecond

	instancer := consul.NewInstancer(client, logger, serviceName, tags, passingOnly)

	factory := addServiceFactory(ctx, "POST", "/login")
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(1, duration, balancer)

	return retry

}
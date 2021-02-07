package dicover

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/lb"
	"time"
)

func MakeDiscoveryEndpoint(ctx context.Context, client consul.Client, logger log.Logger, method, path string) endpoint.Endpoint {
	serviceName := "add_service"
	tags := []string{"add_service", "vj"}
	passingOnly := true

	duration := 500 * time.Millisecond

	instancer := consul.NewInstancer(client, logger, serviceName, tags, passingOnly)

	factory := addServiceFactory(ctx, method, path)
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(1, duration, balancer)

	return retry

}
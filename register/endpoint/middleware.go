package endpoint

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go-kit-microservice-demo/register/util"
	uberlimit "go.uber.org/ratelimit"
	"golang.org/x/time/rate"

)

func AuthMiddleware () endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			token := fmt.Sprint(ctx.Value(util.JWT_CONTEXT_KEY))
			if token == "" {
				err = errors.New("请登陆")
				return nil, err
			}

			info , err := util.ParseToken(token)
			if err != nil {
				return nil, err
			}
			if v, ok := info["username"]; ok {
				ctx = context.WithValue(ctx, "username", v)
			}
			return next(ctx, request)
		}
	}
}

func UberRateLimitMiddleware (limit uberlimit.Limiter) endpoint.Middleware{
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			limit.Take()
			return next(ctx, request)
		}
	}
}

func GolandRateLimitMiddleware (limit *rate.Limiter)endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, errors.New("bad request")
			}
			return next(ctx, request)
		}
	}
}



package endpoint

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go-kit-microservice-demo/util"
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

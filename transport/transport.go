package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"go-kit-microservice-demo/endpoint"
	"go-kit-microservice-demo/util"
	"golang.org/x/time/rate"
	"net/http"
	"os"
	"time"
)

var logger log.Logger

func init () {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestamp)
}

func NewHTTPHandler (ep endpoint.Set) http.Handler {
	m := http.NewServeMux()

	options := []kithttp.ServerOption{
		kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			ctx = context.WithValue(ctx, util.JWT_CONTEXT_KEY, request.Header.Get("Authorization"))
			return ctx
		}),
	}

	//uberLimit := ratelimit.New(10)
	limit := rate.NewLimiter(rate.Every(10 * time.Second),1)

	addAddEndpoint := ep.AddEndpoint
	addAddEndpoint = endpoint.AuthMiddleware()(addAddEndpoint)
	//addAddEndpoint = endpoint.RateLimitMiddleware(uberLimit)(addAddEndpoint)
	addAddEndpoint = endpoint.GolandRateLimitMiddleware(limit)(addAddEndpoint)

	m.Handle("/add", kithttp.NewServer(
		addAddEndpoint,
		decodeHTTPAddRequest,
		encodeResponse,
		options...
		))

	m.Handle("/login", kithttp.NewServer(
		ep.LoginEndpoint,
		decodeHTTPLoginRequest,
		encodeResponse,
		options...
	))

	return m
}

func decodeHTTPAddRequest (_ context.Context, r *http.Request) (interface{}, error){
	var req endpoint.AddRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Log(err.Error())
		return req, err
	}
	return req, nil
}

func decodeHTTPLoginRequest (ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse (ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error" : err.Error(),
	})
}

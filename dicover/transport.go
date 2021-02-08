package main

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler(ep Set) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/add").Handler(kithttp.NewServer(ep.AddEndpoint,
		decodeAddRequest,
		encodeJSONResponse,
		))

	r.Methods("POST").Path("/login").Handler(kithttp.NewServer(ep.LoginEndpoint,
		decodeLoginRequest,
		encodeJSONResponse,
	))

	return r
}

func decodeAddRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
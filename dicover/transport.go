package dicover

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
	kithttp "github.com/go-kit/kit/transport/http"

)

func MakeHTTPHandler(endpoint endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/add").Handler(kithttp.NewServer(endpoint,
		decodeAddRequest,
		encodeJSONResponse,
		))

	r.Methods("POST").Path("/login").Handler(kithttp.NewServer(endpoint,
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
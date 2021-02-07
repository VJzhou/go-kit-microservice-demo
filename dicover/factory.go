package dicover

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	kithttp "github.com/go-kit/kit/transport/http"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func addServiceFactory (_ context.Context, method, path string) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, err error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}

		target, err:= url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}

		target.Path = path
		var (
			enc kithttp.EncodeRequestFunc
			dec kithttp.DecodeResponseFunc
		)

		switch path {
		case "/login":
			enc, dec = encodeJSONRequest, decodeLoginResponse
		case "/add":
			enc, dec = encodeJSONRequest, decodeAddResponse
		default:
			return nil, nil, fmt.Errorf("unknow path %q", path)
		}
		return kithttp.NewClient(method, target, enc, dec).Endpoint(), nil, nil
	}
}


func decodeAddResponse (ctx context.Context, resp *http.Response) (interface{}, error) {
	var response AddResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeLoginResponse(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func encodeJSONRequest(_ context.Context, req *http.Request, request interface{}) error {
	// Both uppercase and count requests are encoded in the same way:
	// simple JSON serialization to the request body.
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}
func encodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}


// 请求参数解析
type AddRequest struct {
	Num1 int `json:"num_1"`
	Num2 int `json:"num_2"`
}

// 返回参数解析
type AddResponse struct {
	Sum int `json:"sum"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
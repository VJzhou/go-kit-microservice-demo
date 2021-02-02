package endpoint

import (
	"go-kit-microservice-demo/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

// 请求参数解析
type AddRequest struct {
	Num1 int `json:"num_1"`
	Num2 int `json:"num_2"`
}

// 返回参数解析
type AddResponse struct {
	Sum int `json:"sum"`
}

// endpoint 集合
type Set struct {
	AddEndpoint endpoint.Endpoint
}

// 集合工厂函数
func NewEndpointSet(svc service.Service) Set {
	return Set{AddEndpoint:MakeAddEndpoint(svc)}
}

func MakeAddEndpoint (svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddRequest)
		sum := svc.Add(req.Num1, req.Num2)
		return AddResponse{Sum:sum}, nil
	}
}

func (s *Set) Add (ctx context.Context, num1, num2 int) int {
	resp, _ := s.AddEndpoint(ctx, AddRequest{num1, num2})
	getResp := resp.(AddResponse)
	return getResp.Sum
}



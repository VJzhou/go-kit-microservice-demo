package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-kit-microservice-demo/register/service"
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

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

type HealthCheckRequest struct {

}

type HealthCheckResponse struct {
	Status bool `json:"status"`
}


// endpoint 集合
type Set struct {
	AddEndpoint endpoint.Endpoint
	LoginEndpoint endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
}

// 集合工厂函数
func NewEndpointSet(svc service.Service) Set {
	return Set{
		AddEndpoint:MakeAddEndpoint(svc),
		LoginEndpoint:MakeLoginEndpoint(svc),
		HealthCheckEndpoint:MakeHealthEndpoint(svc),
	}
}

func MakeAddEndpoint (svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddRequest)
		sum := svc.Add(ctx, req.Num1, req.Num2)
		return AddResponse{Sum:sum}, nil
	}
}

func MakeLoginEndpoint (svc service.Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(LoginRequest)
		token ,err := svc.Login(ctx, req.Username, req.Password)
		if err != nil {
			return nil, err
		}
		return LoginResponse{token }, nil
	}
}

func MakeHealthEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			res := svc.HealthCheck(ctx)
			return HealthCheckResponse{Status:res}, nil
	}
}

func (s *Set) Add (ctx context.Context, num1, num2 int) int {
	resp, _ := s.AddEndpoint(ctx, AddRequest{num1, num2})
	getResp := resp.(AddResponse)
	return getResp.Sum
}

func (s *Set) Login (ctx context.Context, username, password string) (string ,  error) {

	resp, err := s.LoginEndpoint(ctx, LoginRequest{username, password})
	if err != nil {
		return "", err
	}
	getResp := resp.(LoginResponse)
	return getResp.Token, nil
}

func (s *Set) HealthCheck(ctx context.Context) bool {
	resp, _:= s.HealthCheckEndpoint(ctx, HealthCheckRequest{})
	getResp := resp.(HealthCheckResponse)
	return getResp.Status
}


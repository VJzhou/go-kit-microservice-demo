package service

import (
	"context"
	"errors"
	"go-kit-microservice-demo/register/util"
)

// 定义接口
type Service interface {
	Add (ctx context.Context,num1 ,num2 int) int
	Login(ctx context.Context,username, password string) (string , error)
	HealthCheck(ctx context.Context) bool
}

type ServiceMiddleware func(service Service) Service

type addService struct {}


// 工厂函数
func NewService() Service {
	return &addService{}
}

// 实现接口
func (a addService) Add(ctx context.Context,num1, num2 int) int {
	return num1 + num2
}

func (a addService) Login(ctx context.Context,username, password string) (string, error) {
	if username == "vj" && password == "111" {
		return util.GenerateToken(username, 1)
	}
	return "", errors.New("Account error")
}

func (a addService) HealthCheck(_ context.Context) bool {
	return true
}


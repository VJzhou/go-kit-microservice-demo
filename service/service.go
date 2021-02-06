package service

import (
	"errors"
	"go-kit-microservice-demo/util"
)

// 定义接口
type Service interface {
	Add (num1 ,num2 int) int
	Login(username, password string) (string , error)
}

type ServiceMiddleware func(service Service) Service

type addService struct {}

// 工厂函数
func NewService() Service {
	return &addService{}
}

// 实现接口
func (a addService) Add(num1, num2 int) int {
	return num1 + num2
}

func (a addService) Login(username, password string) (string, error) {
	if username == "vj" && password == "111" {
		return util.GenerateToken(username, 1)
	}
	return "", errors.New("Account error")
}

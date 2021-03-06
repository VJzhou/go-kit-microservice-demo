package main

import (
	"go-kit-microservice-demo/service"
	"fmt"
	"github.com/go-kit/kit/log"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next service.Service
}

func (l LoggingMiddleware) Login(username, password string) (ret string, err error) {
	defer func(begin time.Time) {
		_ = l.Logger.Log(
			"method", "login",
			"input", fmt.Sprintf("username:%s, password:%s", username, password),
			"output" , ret,
			"took", time.Since(begin),
		)
	}(time.Now())
	ret, _= l.Next.Login(username, password)
	return
}

func (l LoggingMiddleware) Add (num1, num2 int) (ret int) {
	defer func(begin time.Time) {
		_ = l.Logger.Log(
			"method", "add",
				"input", fmt.Sprintf("num1:%d, num2:%d", num1, num2),
				"output" , ret,
				"took", time.Since(begin),
		)
	}(time.Now())
	ret = l.Next.Add(num1, num2)
	return
}
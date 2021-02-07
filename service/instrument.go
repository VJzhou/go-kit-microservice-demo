package service

import (
	"fmt"
	"github.com/go-kit/kit/metrics"
	"time"
)

type MetricMiddleware struct {
	Service
	RequestCount metrics.Counter
	RequestLatency metrics.Histogram
}

func Metrics(requestCount metrics.Counter, requestLatency metrics.Histogram) ServiceMiddleware {
	return func(service Service) Service {
		return MetricMiddleware{
			Service:        service,
			RequestCount:   requestCount,
			RequestLatency: requestLatency,
		}
	}
}

func (mw MetricMiddleware) Add (num1, num2 int) (ret int) {

	defer func(begin time.Time) {
		lvs := []string{"method", "Add", "err", ""}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	ret = mw.Service.Add(num1, num2)
	return
}

func (mw MetricMiddleware) Login(username, password string) (token string, err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", "login", "err", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	token , err = mw.Service.Login(username, password)
	return
}
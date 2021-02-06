package service

import (
	"fmt"
	"github.com/go-kit/kit/metrics"
	"time"
)

type metricMiddleware struct {
	Service
	requestCount metrics.Counter
	requestLatency metrics.Histogram
}

func Metrics(requestCount metrics.Counter, requestLatency metrics.Histogram) ServiceMiddleware {
	return func(service Service) Service {
		return metricMiddleware{
			Service:        service,
			requestCount:   requestCount,
			requestLatency: requestLatency,
		}
	}
}

func (mw metricMiddleware) Add (num1, num2 int) (ret int) {

	defer func(begin time.Time) {
		lvs := []string{"method", "Add", "err", ""}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())


	ret = mw.Service.Add(num1, num2)
	return
}

func (mw metricMiddleware) Login(username, password string) (token string, err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", "login", "err", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())


	token , err = mw.Service.Login(username, password)
	return
}
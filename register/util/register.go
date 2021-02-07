package util

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/pborman/uuid"
	"os"
	"strconv"
)


func Register(consulHost, consulPort, svcHost, svcPort string, logger log.Logger) (register sd.Registrar) {

	// 创建consul 客户端链接
	var client consul.Client
	{
		consulCientfg := api.DefaultConfig()
		consulCientfg.Address = consulHost + ":" + consulPort
		consulClient, err := api.NewClient(consulCientfg)
		if err != nil {
			logger.Log("create consul client error:", err)
			os.Exit(1)
		}
		client = consul.NewClient(consulClient)
	}

	// 设置consul服务健康检查参数
	check := api.AgentServiceCheck{
		HTTP: "http://" + svcHost + ":" + svcPort + "/health",
		Interval: "10s",
		Timeout:"1s",
		Notes: "Consul check service health status",
	}

	port, _ := strconv.Atoi(svcPort)

	// 设置微服五向Consul 注册的信息
	regi := api.AgentServiceRegistration{
		ID: "add_service" + uuid.New(),
		Name: "add_service",
		Address: svcHost,
		Port: port,
		Tags: []string{"add_service", "vj"},
		Check: &check,
	}

	// 注册
	register = consul.NewRegistrar(client, &regi, logger)
	return
}
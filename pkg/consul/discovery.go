package consul

import (
	"errors"
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
)

// Discovery 服务发现
func Discovery(addr string, serviceName string) (string, error) {
	cfg := &consulapi.Config{Address: addr}
	client, err := consulapi.NewClient(cfg)
	if err != nil {
		return "", err
	}

	services, _, err := client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return "", err
	}

	if len(services) == 0 {
		return "", errors.New("没有找到服务")
	}

	return fmt.Sprintf("%s:%d", services[0].Service.Address, services[0].Service.Port), nil
}

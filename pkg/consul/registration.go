package consul

import (
	"context"
	"fmt"
	"net"

	"github.com/google/uuid"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Service struct {
	client      *consulapi.Client
	serviceID   string
	serviceName string
	serviceHost string
	servicePort int
}

// Health struct health
type Health struct{}

func NewConsul(address string, serviceName string, servicePort int) (*Service, error) {
	config := consulapi.DefaultConfig()
	config.Address = address
	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Service{
		client:      client,
		serviceName: serviceName,
		serviceHost: localIP(),
		servicePort: servicePort,
	}, nil
}

func (c *Service) Register() error {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	registration := &consulapi.AgentServiceRegistration{
		ID:      uid.String(),
		Name:    c.serviceName,
		Address: c.serviceHost,
		Port:    c.servicePort,
		Tags:    []string{"primary"},
		// http
		//Check: &consulapi.AgentServiceCheck{
		//	HTTP:                           fmt.Sprintf("http://%s:%d/health", serverHost, serverPort),
		//	Timeout:                        "5s",
		//	Interval:                       "5s",
		//	DeregisterCriticalServiceAfter: "120s", // 故障检查失败120s后 consul自动将注册服务删除
		//},
		// grpc
		Check: &consulapi.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%v:%v/%v", c.serviceHost, c.servicePort, c.serviceName),
			Interval:                       "5s",
			Timeout:                        "2s",
			DeregisterCriticalServiceAfter: "20s", // 故障检查失败120s后 consul自动将注册服务删除
		},
	}

	err = c.client.Agent().ServiceRegister(registration)
	if err != nil {
		return err
	}

	c.serviceID = uid.String()
	return nil
}

func (c *Service) Deregister() error {
	return c.client.Agent().ServiceDeregister(c.serviceID)
}

// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (h *Health) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *Health) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

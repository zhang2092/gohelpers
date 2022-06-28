package etcd

import (
	"context"
	"sync"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientV3 "go.etcd.io/etcd/client/v3"
)

type RemoteService struct {
	name  string
	node  map[string]string
	mutex sync.Mutex
}

type Resolver struct {
	v3       *clientV3.Client
	endpoint []string
}

// NewResolver 构造 resolver 对象
func NewResolver(endpoint []string) (*Resolver, error) {
	client, err := clientV3.New(clientV3.Config{
		Endpoints:   endpoint,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return &Resolver{v3: client, endpoint: endpoint}, nil
}

// Close 关闭client
func (r *Resolver) Close() error {
	return r.v3.Close()
}

// Discovery 发现服务
func (r *Resolver) Discovery(serviceName string) (*RemoteService, error) {
	service := &RemoteService{
		name: serviceName,
		node: make(map[string]string, 1),
	}

	kv := clientV3.NewKV(r.v3)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := kv.Get(ctx, serviceName, clientV3.WithPrefix())
	if err != nil {
		return nil, err
	}

	service.mutex.Lock()
	for _, kv := range resp.Kvs {
		service.node[string(kv.Key)] = string(kv.Value)
	}
	service.mutex.Unlock()

	go r.watchServiceUpdate(service)

	return service, nil
}

// watchServiceUpdate 监控服务目录下的事件
func (r *Resolver) watchServiceUpdate(service *RemoteService) {
	watcher := clientV3.NewWatcher(r.v3)
	// Watch 服务目录下的更新
	watchChan := watcher.Watch(context.TODO(), service.name, clientV3.WithPrefix())
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			service.mutex.Lock()
			switch event.Type {
			case mvccpb.PUT: // PUT事件，目录下有了新key
				service.node[string(event.Kv.Key)] = string(event.Kv.Value)
			case mvccpb.DELETE: // DELETE事件，目录中有key被删掉(Lease过期，key 也会被删掉)
				delete(service.node, string(event.Kv.Key))
			}
			service.mutex.Unlock()
		}
	}
}

// GetName 获取服务名称
func (r *RemoteService) GetName() string {
	return r.name
}

// GetNode 获取服务列表
func (r *RemoteService) GetNode() map[string]string {
	return r.node
}

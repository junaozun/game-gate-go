package register_center

import (
	"context"
	"sync"

	"github.com/junaozun/gogopkg/etcdx"
	clientv3 "go.etcd.io/etcd/client/v3"
)

/*注册中心服务器直接使用Etcd*/

type RegisterServer struct {
	client   etcdx.IClient
	stopChan chan struct{}
	ttl      int
	mutex    sync.RWMutex
	leaseId  clientv3.Lease
	key      string
	value    string
}

func NewRegisterCenter(servers string, ttl int, key string) (*RegisterServer, error) {
	client, err := etcdx.NewClientWithConfig(etcdx.Config{
		Servers:        servers,
		DialTimeout:    5,
		RequestTimeout: 5,
	})
	if err != nil {
		return nil, err
	}
	return &RegisterServer{
		client:   client,
		stopChan: make(chan struct{}),
		ttl:      ttl,
		key:      key,
	}, nil
}

func (r *RegisterServer) Start(ctx context.Context) error {

}

func (r *RegisterServer) Stop(ctx context.Context) error {

}

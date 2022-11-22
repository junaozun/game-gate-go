package gateway

import (
	"fmt"
	"sync"

	"github.com/junaozun/game-gate-go/pkg/ip"
	"github.com/junaozun/game-gate-go/pkg/snowflake"
	"github.com/junaozun/game-gate-go/server/client_session"
	"github.com/junaozun/game-gate-go/server/register_center"
	"github.com/junaozun/game-gate-go/server/server_session"
	"github.com/junaozun/gogopkg/app"
	"github.com/junaozun/gogopkg/logrusx"
)

const (
	host = "0.0.0.0:"
)

type GateServer struct {
	ip        string
	port      string
	registry  *register_center.RegisterServer // 注册中心服务
	logicSrv  *server_session.LogicServer     // gateway 与 logic 的长连接服务
	clientSrv *client_session.ClientServer    // gateway 与 client 的长连接服务
	wg        sync.WaitGroup
	stopChan  chan struct{}
	IdWorker  *snowflake.IDWorker
}

func NewGateWay(cfg *GateServerConfig) *GateServer {
	g := &GateServer{
		stopChan: make(chan struct{}),
		IdWorker: snowflake.NewIDWorker(cfg.Gate.MachineId),
		port:     cfg.Client.Port,
	}
	g.logicSrv = server_session.NewLogicServer(g, host+cfg.Logic.Port)
	g.clientSrv = client_session.NewClientServer(g, host+g.port)

	// 注册中心
	// 1.获取本机器的ip地址
	gateIp, err := ip.GetNetworkCardIp(cfg.Gate.NetworkCard)
	if err != nil {
		panic(fmt.Errorf("get networkCard[%v] ip error: %v", cfg.Gate.NetworkCard, err))
	}
	g.ip = gateIp

	registerKey := fmt.Sprintf("%valive/%v:%v", cfg.Etcd.Key, gateIp, g.port)
	// 2.new 注册中心服务器
	rs, err := register_center.NewRegisterCenter(cfg.Etcd.ServerAddress, cfg.Etcd.TTL, registerKey)
	if err != nil {
		panic(err)
	}
	g.registry = rs
	return g
}

func (g *GateServer) Run() {
	addr := g.ip + ":" + g.port
	gate := app.New(
		app.OnBeginHook(func() {
			logrusx.Log.WithFields(logrusx.Fields{
				"addr": addr,
			}).Info("gateServer start .....")
		}),
		app.OnExitHook(func() {
			logrusx.Log.WithFields(logrusx.Fields{
				"addr": addr,
			}).Info("gateServer close .....")
		}),
		app.Name("gateway"),
		app.Runners(g.logicSrv, g.clientSrv, g.registry),
	)
	if err := gate.Run(); err != nil {
		panic(err)
	}
}

package server_session

import (
	"context"
	"sync"

	"github.com/junaozun/game-gate-go/network"
	"github.com/junaozun/game-gate-go/server/gateway"
	"github.com/junaozun/game-gate-go/transport"
)

type LogicServer struct {
	GateServer      *gateway.GateServer
	netWork         string
	address         string
	onlineConn      map[uint64]*network.Conn
	transportServer transport.IServerTransport
	mutex           sync.RWMutex
}

func NewLogicServer(gateServer *gateway.GateServer, address string) *LogicServer {
	srv := &LogicServer{
		address:    address,
		GateServer: gateServer,
		netWork:    "tcp",
		onlineConn: make(map[uint64]*network.Conn),
	}
	return srv
}

func (l *LogicServer) Start(ctx context.Context) error {
	tops := []transport.ServerTransportOption{
		transport.WithServerAddress(l.address),
		transport.WithIdGenerator(l.GateServer.IdWorker), //
	}
	l.transportServer = transport.GetServerTransport(l.netWork)
	return l.transportServer.ListenAndServe(ctx, tops...)
}

func (l *LogicServer) Stop(ctx context.Context) error {
	// 关闭所有的在线连接
	l.mutex.Lock()
	defer l.mutex.Unlock()
	for uuid, conn := range l.onlineConn {
		conn.Close()
		delete(l.onlineConn, uuid)
	}
	return nil
}

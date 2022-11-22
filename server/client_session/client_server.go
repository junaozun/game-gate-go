package client_session

import (
	"context"

	"github.com/junaozun/game-gate-go/network"
	"github.com/junaozun/game-gate-go/server/gateway"
	"github.com/junaozun/game-gate-go/transport"
)

type ClientServer struct {
	GateServer      *gateway.GateServer
	address         string
	network         string
	onlineConn      map[uint64]*network.Conn
	transportServer transport.IServerTransport
}

func NewClientServer(gateServer *gateway.GateServer, address string) *ClientServer {
	srv := &ClientServer{
		address:    address,
		network:    "tcp",
		GateServer: gateServer,
		onlineConn: make(map[uint64]*network.Conn),
	}
	return srv
}

func (c *ClientServer) Start(ctx context.Context) error {
	tops := []transport.ServerTransportOption{
		transport.WithServerAddress(c.address),
		transport.WithIdGenerator(c.GateServer.IdWorker), //
		transport.WithMaxReadBodyBytes(111),
		transport.WithMaxWriteBodyBytes(111),
	}
	c.transportServer = transport.GetServerTransport(c.network)
	return c.transportServer.ListenAndServe(ctx, tops...)
}

func (c *ClientServer) Stop(ctx context.Context) error {
}

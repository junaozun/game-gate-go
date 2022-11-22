package udp_server

import (
	"context"

	"github.com/junaozun/game-gate-go/network"
	"github.com/junaozun/game-gate-go/transport"
)

type UdpServer struct {
	config network.Config
}

func (u *UdpServer) ListenAndServe(context.Context, ...transport.ServerTransportOption) error {
	return nil
}

func (u *UdpServer) Config() network.Config {
	return u.config
}

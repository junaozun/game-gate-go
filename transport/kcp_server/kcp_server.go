package kcp_server

import (
	"context"

	"github.com/junaozun/game-gate-go/network"
	"github.com/junaozun/game-gate-go/transport"
)

type KcpServer struct {
	config network.Config
}

func (k *KcpServer) ListenAndServe(context.Context, ...transport.ServerTransportOption) error {
	return nil
}

func (k *KcpServer) Config() network.Config {
	return k.config
}

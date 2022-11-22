package ws_server

import (
	"context"

	"github.com/junaozun/game-gate-go/network"
	"github.com/junaozun/game-gate-go/transport"
)

type WsServer struct {
	config network.Config
}

func (w *WsServer) ListenAndServe(context.Context, ...transport.ServerTransportOption) error {
	return nil
}

func (w *WsServer) Config() network.Config {
	return w.config
}

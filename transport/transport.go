package transport

import (
	"context"
)

const (
	Tansport_tcp  = "tcp"
	Tansport_udp  = "udp"
	Tansport_ws   = "ws"
	Tansport_kcp  = "kcp"
	Tansport_http = "http"
)

// IServerTransport 传输层主要提供一种监听和处理请求的能力
type IServerTransport interface {
	ListenAndServe(context.Context, ...ServerTransportOption) error
}

var serverTransportMap = make(map[string]IServerTransport)

func RegisterServerTransport(name string, transport IServerTransport) {
	if serverTransportMap == nil {
		serverTransportMap = make(map[string]IServerTransport)
	}
	serverTransportMap[name] = transport
}

func GetServerTransport(name string) IServerTransport {
	if v, ok := serverTransportMap[name]; ok {
		return v
	}
	return nil
}

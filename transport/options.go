package transport

import (
	"context"
	"time"

	"github.com/junaozun/game-gate-go/pkg/generatorid"
)

type ServerTransportOptions struct {
	Address           string                  // address，e.g: ip://127.0.0.1：8080
	Timeout           time.Duration           // transport layer request timeout ，default: 2 min
	Handler           Handler                 // handler
	IdGenerator       generatorid.IDGenerator // id自增生成器
	MaxReadBodyBytes  int                     // 最大读
	MaxWriteBodyBytes int                     // 最大写
	SerializationType string                  // proto、json、msgpack
	KeepAlivePeriod   time.Duration           // keepalive period
}

type Handler interface {
	Handle(context.Context, []byte) ([]byte, error)
}

type ServerTransportOption func(*ServerTransportOptions)

func WithServerAddress(address string) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.Address = address
	}
}

func WithIdGenerator(gen generatorid.IDGenerator) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.IdGenerator = gen
	}
}

func WithMaxReadBodyBytes(maxRead int) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.MaxReadBodyBytes = maxRead
	}
}

func WithMaxWriteBodyBytes(maxWrite int) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.MaxWriteBodyBytes = maxWrite
	}
}

func WithHandler(handler Handler) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.Handler = handler
	}
}

func WithServerTimeout(timeout time.Duration) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.Timeout = timeout
	}
}

func WithSerializationType(serializationType string) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.SerializationType = serializationType
	}
}

func WithKeepAlivePeriod(keepAlivePeriod time.Duration) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.KeepAlivePeriod = keepAlivePeriod
	}
}

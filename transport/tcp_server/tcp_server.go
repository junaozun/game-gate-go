package tcp_server

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/junaozun/game-gate-go/transport"
)

type TcpServer struct {
	opts *transport.ServerTransportOptions
	mu   sync.RWMutex
}

var DefaultTcpServerTransport = NewTcpServerTransport()

var NewTcpServerTransport = func() *TcpServer {
	return &TcpServer{
		opts: &transport.ServerTransportOptions{},
	}
}

func init() {
	transport.RegisterServerTransport("tcp", DefaultTcpServerTransport)
}

func (t *TcpServer) ListenAndServe(ctx context.Context, opts ...transport.ServerTransportOption) error {

	for _, o := range opts {
		o(t.opts)
	}

	ln, err := net.Listen("tcp", t.opts.Address)
	if err != nil {
		return err
	}

	return t.serve(ctx, ln)
}

// serve listen
func (t *TcpServer) serve(ctx context.Context, lis net.Listener) error {
	var tempDelay time.Duration

	tl, ok := lis.(*net.TCPListener)
	if !ok {
		return errors.New("[TcpServer] serve not tcp listener")
	}
	for {
		select {
		case <-ctx.Done(): // 监听到取消server
			err := lis.Close()
			if err != nil {
				return err
			}
			return ctx.Err()
		default:
		}
		tcpConn, err := tl.AcceptTCP()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				time.Sleep(tempDelay)
				continue
			}
			return err
		}
		c := NewTcpConn(tcpConn, t)
		c.serveConn()
	}
}

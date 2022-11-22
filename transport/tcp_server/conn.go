package tcp_server

import (
	"bufio"
	"net"
)

type TcpConn struct {
	uuid     int64 // 用户uuid
	conn     *net.TCPConn
	server   *TcpServer
	bufRead  *bufio.Reader // 带有缓冲的读conn数据的buf
	bufWrite *bufio.Writer // 带有缓冲的写数据入conn的buf
}

func NewTcpConn(conn *net.TCPConn, server *TcpServer) *TcpConn {
	return &TcpConn{
		conn:     conn,
		server:   server,
		bufRead:  bufio.NewReaderSize(conn, server.opts.MaxReadBodyBytes+16),
		bufWrite: bufio.NewWriterSize(conn, server.opts.MaxWriteBodyBytes+16),
	}
}

func (c *TcpConn) serveConn() {

}

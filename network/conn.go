package network

import (
	"net"
)

type Conn struct {
	uuid int64 // 用户uuid
	conn *net.Conn
}

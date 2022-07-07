package tcp

import (
	"context"
	"net"
)

// handler 是应用层服务器的抽象
type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close() error
}

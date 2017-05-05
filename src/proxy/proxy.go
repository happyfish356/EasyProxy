package proxy

import (
	"config"
	"net"
)

type Proxy interface {
	Init(config *config.Config)
	Check()
	Clean(url string)
	Recover(url string)
	Dispatch(con net.Conn)
	Close()
}

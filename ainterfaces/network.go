package ainterfaces

import (
	"net"
	"nso/core"
)

type ISender interface {
	SendMessage(message *core.Message)
}

type ISession interface {
	ISender
	GetID() int
	ReadMessage() *core.Message
	Update()
	Disconnect()
	SendKey()
	GetUser() IUser
}

type INetworkLoop interface {
	ListenSync() error
	Accept(conn net.Conn) error
	Close() error
}

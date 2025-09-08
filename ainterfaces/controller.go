package ainterfaces

import "nso/core"

type IController interface {
	OnMessage(session ISession, message *core.Message) error
}

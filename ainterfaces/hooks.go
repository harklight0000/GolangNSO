package ainterfaces

import (
	. "nso/request"
	. "nso/response"
)

type IHooks interface {
	ListenASync() (err error)
	AddHookAsync(request *HookRequest) chan *HookResponse
}

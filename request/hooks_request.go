package request

import (
	. "nso/constants"
	"nso/logging"
	"nso/response"
)

type HookRequest struct {
	Data interface{}
	Done chan *response.HookResponse
	*logging.Caller
	HookType HookType
}

func NewHookRequest(typeHook HookType, data interface{}) *HookRequest {
	return &HookRequest{
		Data:     data,
		Done:     make(chan *response.HookResponse),
		Caller:   logging.NewCaller(),
		HookType: typeHook,
	}
}

func (r *HookRequest) String() string {
	return "Hook type: " + r.HookType.String() + " at " + r.Caller.String()
}

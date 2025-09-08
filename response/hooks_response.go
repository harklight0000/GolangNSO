package response

import . "nso/constants"

type HookResponse struct {
	Data  interface{}
	Error error
	Type  HookType
}

func NewHookResponse(data interface{}, err error, hookType HookType) *HookResponse {
	return &HookResponse{Data: data, Error: err, Type: hookType}
}

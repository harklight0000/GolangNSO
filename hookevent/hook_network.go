package hookevent

import (
	"fmt"
	"github.com/rotisserie/eris"
	"net"
	. "nso/constants"
	"nso/logging"
	. "nso/request"
	. "nso/response"
)

func (h *Hooks) hookNetwork(request *HookRequest) *HookResponse {
	switch request.HookType {
	case HOOK_NEW_CONNECTION_ENTER:
		conn, ok := request.Data.(net.Conn)
		if !ok {
			logging.Logger.Info(fmt.Sprintf("New connection from %s", conn.RemoteAddr().String()))
			logging.Logger.Info("Hooking network add new connection")
			return NewHookResponse(nil, nil, request.HookType)
		} else {
			return NewHookResponse(nil, eris.New("Invalid request data"), request.HookType)
		}
	}
	return NewHookResponse(nil, eris.New("Unknown hook type"), request.HookType)
}

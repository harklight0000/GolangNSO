package hookevent

import (
	"github.com/rotisserie/eris"
	. "nso/constants"
	"nso/logging"
	. "nso/request"
	. "nso/response"
)

func (h *Hooks) hookApp(request *HookRequest) *HookResponse {
	logging.Logger.Info("Hooking app " + request.String())
	switch request.HookType {
	case HOOK_APP_START:
		fallthrough
	case HOOK_APP_END:
		return NewHookResponse(nil, eris.New("The hook type is preserved"), request.HookType)
	case HOOK_APP_TERMINATE:
		// do some logic to clean up
		var err = h.appCtx.Close()
		if err != nil {
			return NewHookResponse(nil, eris.Wrap(err, "Close app fail"), request.HookType)
		} else {
			logging.Logger.Info("Close app success")
			return NewHookResponse("Stop application success", nil, request.HookType)
		}
	}
	return NewHookResponse(nil, eris.New("Unknown hook type"), request.HookType)
}

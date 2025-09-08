package hookevent

import (
	"context"
	. "fmt"
	"github.com/rotisserie/eris"
	. "nso/ainterfaces"
	. "nso/constants"
	. "nso/logging"
	. "nso/request"
	. "nso/response"
	"time"
)

type Hooks struct {
	appCtx          IAppContext
	requests        chan *HookRequest
	ctx             context.Context
	cancel          context.CancelFunc
	nWorkers        int
	responseTimeOut int64
}

func NewHooks(nWorkers int, responseTimeOut int64, queueSize int, c IAppContext) *Hooks {
	var h = &Hooks{
		requests:        make(chan *HookRequest, queueSize),
		nWorkers:        nWorkers,
		responseTimeOut: responseTimeOut,
		appCtx:          c,
	}
	h.ctx, h.cancel = context.WithCancel(h.appCtx.GetContext())
	return h
}

func (h *Hooks) ListenASync() error {
	if h.nWorkers <= 0 {
		return eris.New("Number of workers must be greater than 0")
	}
	Logger.Info(Sprintf("Starting hook listeners with %d workers", h.nWorkers))
	for i := 0; i < h.nWorkers; i++ {
		go func(workerNumber int) {
			Logger.Info(Sprintf("Starting hook listener %d", workerNumber))
		event:
			for {
				select {
				case <-h.ctx.Done():
					return
				case request := <-h.requests:
					Logger.Info(Sprintf("Worker %d handling hook request type %d", workerNumber, request.HookType))
					select {
					case request.Done <- h.handle(request):
					case <-time.After(time.Duration(h.responseTimeOut) * time.Second):
						Logger.Info("Timeout waiting for hook response")
						continue event
					}
				}
			}
		}(i)
	}
	return nil
}

func (h *Hooks) AddHookAsync(request *HookRequest) chan *HookResponse {
	if h.appCtx == nil {
		go func() {
			select {
			case request.Done <- NewHookResponse(nil, eris.New("AppContext must be set through SetAppContext method to run trigger hook"), request.HookType):
			case <-time.After(time.Duration(h.responseTimeOut) * time.Second):
				Logger.Info("Timeout waiting for hook response")
				return
			}
		}()
		return request.Done
	}
	if request == nil {
		res := make(chan *HookResponse)
		go func() {
			select {
			case res <- NewHookResponse(nil, eris.New("HookRequest cannot be nil"), request.HookType):
			case <-time.After(time.Duration(h.responseTimeOut) * time.Second):
				Logger.Info("Timeout waiting for hook response")
				return
			}
		}()
		return res
	}
	h.requests <- request
	return request.Done
}

func (h *Hooks) handle(request *HookRequest) (r *HookResponse) {
	defer func() {
		if r := recover(); r != nil {
			Logger.Error(Sprintf("Recovering from panic in hook handler: %v", r))
			r = NewHookResponse(nil, eris.New(Sprintf("Recovering from panic in hook handler: %v", r)), request.HookType)
		}
	}()
	var hookType = request.HookType
	Logger.Info(Sprintf("Received hook %s", request.String()))
	switch {
	case HOOK_NET_WORK_START < hookType && hookType < HOOK_NET_WORK_STOP:
		return h.hookNetwork(request)
	case HOOK_APP_START < hookType && hookType < HOOK_APP_END:
		return h.hookApp(request)
	}
	return NewHookResponse(nil, eris.New("Unknown hook type"), request.HookType)
}

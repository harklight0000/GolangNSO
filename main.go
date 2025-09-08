package main

import (
	"errors"
	"github.com/rotisserie/eris"
	"github.com/ztrue/shutdown"
	"go.uber.org/zap"
	"net"
	. "nso/ainterfaces"
	"nso/app"
	"nso/constants"
	"nso/errs"
	"nso/logging"
	"nso/request"
)

func run() {
	var err error

	var appContext IAppContext = app.NewAppContext()
	err = appContext.Init()
	if err != nil {
		logging.Logger.Error("Error initializing app context", zap.Error(err))
		return
	}
	var netLoop = appContext.GetNetworkLoop()
	// Hooks listen
	hooks := appContext.GetHooks()
	shutdown.Add(func() {
		defer appContext.CloseGoroutines()
		resp := hooks.AddHookAsync(request.NewHookRequest(constants.HOOK_APP_TERMINATE, nil))
		select {
		case r := <-resp:
			logging.Logger.Info("Hook response received")
			if r.Error != nil {
				logging.Logger.Info("Hook response error", zap.Error(r.Error))
			}

		}
	})
	go shutdown.Listen()
	err = hooks.ListenASync()
	if err != nil {
		logging.Logger.Error(errs.ToString(err))
	}
	// Network listen
	err = netLoop.ListenSync()
	if err != nil {
		root := eris.Unwrap(err)
		if errors.Is(root, net.ErrClosed) {
			logging.Logger.Info("Listener closed")
		} else {
			logging.Logger.Error(errs.ToString(err))
		}
	}
	logging.Logger.Info("Server exited")
}

func main() {
	run()
}

package networking

import (
	"context"
	"errors"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"
	"net"
	. "nso/ainterfaces"
	"nso/config"
	"nso/constants"
	"nso/logging"
	"nso/request"
)

type NetworkLoop struct {
	appCtx   IAppContext
	ctx      context.Context
	Cancel   context.CancelFunc
	ln       net.Listener
	cfg      *config.AppConfig
	hooks    IHooks
	sessions ISession
}

func (n *NetworkLoop) ListenSync() error {
	if n.appCtx == nil {
		return eris.New("AppContext must be set through SetAppContext method to run listen")
	}
	logging.Logger.Info("App listening on port: " + n.cfg.ServerPort)
	for {
		conn, er := n.ln.Accept()
		if er != nil {
			if errors.Is(er, net.ErrClosed) {
				logging.Logger.Info("Listener closed")
			}
			return eris.Wrap(er, "Error accepting connection from std lib")
		}
		er = n.Accept(conn)
		if er != nil {
			return eris.Wrap(er, "Error accepting connection from my code")
		}
	}
}

func (n *NetworkLoop) Accept(conn net.Conn) error {
	<-n.hooks.AddHookAsync(request.NewHookRequest(constants.HOOK_NEW_CONNECTION_ENTER, conn))
	logging.Logger.Info("Go through hooks blocking")
	session := NewSession(conn, n.appCtx)
	session.Update()
	return nil
}

func (n *NetworkLoop) Close() error {
	err := n.ln.Close()
	if err != nil {
		logging.Logger.Info("Closed listener error", zap.Error(eris.Wrap(err, "Error closing listener")))
		return eris.Wrap(err, "Error closing listener")
	}
	logging.Logger.Info("Closed listener success")
	return nil
}

func NewNetworkLoop(c IAppContext) *NetworkLoop {
	var n = &NetworkLoop{
		appCtx: c,
	}
	n.ctx, n.Cancel = context.WithCancel(n.appCtx.GetContext())
	n.cfg = n.appCtx.GetConfig()
	n.hooks = n.appCtx.GetHooks()
	ln, err := net.Listen("tcp", ":"+n.cfg.ServerPort)
	if err != nil {
		logging.Logger.Panic("Error listening on port: "+n.cfg.ServerPort, zap.Error(err))
	}
	n.ln = ln
	return n
}

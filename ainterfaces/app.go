package ainterfaces

import (
	"context"
	"io"
	"nso/config"
)

type IAppContext interface {
	GetConfig() *config.AppConfig
	GetDatabase() IDatabase
	GetContext() context.Context
	GetHooks() IHooks
	GetNetworkLoop() INetworkLoop
	io.Closer
	CloseGoroutines()
	Init() error
	GameConfig() *config.GameConfig
	GetGameData() IGameData
	GetController() IController
	GetUserManager() IUserManager
	GetMapManager() IMapManager
	GetMenuFactory() IMenuFactory
	GetEffectFactory() IEffectFactory
	ClansManager() IClansManager
	GetItemSellListManager() IItemSellListManager
}

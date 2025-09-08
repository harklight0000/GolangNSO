package app

import (
	"context"
	"go.uber.org/zap"
	. "nso/ainterfaces"
	"nso/config"
	"nso/controllers"
	"nso/database"
	"nso/hookevent"
	"nso/logging"
	"nso/menu"
	"nso/networking"
	"nso/objects"
)

func NewAppContext() IAppContext {
	ctx, cancel := context.WithCancel(context.Background())
	c := &ContextImpl{
		config:  config.GetAppConfig(),
		gameCfg: config.GetGameConfig(),
		ctx:     ctx,
		cancel:  cancel,
	}
	cfg := c.config
	if cfg.UseSQL {
		c.database = database.NewSQLDatabaseAdapter(database.InitSQLDB())
	} else {
		c.database = database.NewMongoDatabaseAdapter(database.InitMongoDB())
	}
	c.hooks = hookevent.NewHooks(cfg.NHooksWorkers, cfg.HooksTimeOut, cfg.HookQueueSize, c)
	c.netWork = networking.NewNetworkLoop(c)
	c.gameData = objects.NewGameData(c)
	c.controller = controllers.NewController(c)
	c.userManager = objects.NewUserManager(c)
	c.mapManager = objects.NewMapManager(c, c.gameData.MapEntities())
	c.menuFactory = menu.NewProcessorFactory(c)
	c.effectFactory = objects.NewEffectFactory(c.gameData.Effects())
	c.clanManager = objects.NewClansManager()
	c.itemSellListManager = objects.NewItemSellListManager(c)
	return c
}

type ContextImpl struct {
	config              *config.AppConfig
	gameCfg             *config.GameConfig
	database            IDatabase
	ctx                 context.Context
	hooks               IHooks
	netWork             INetworkLoop
	cancel              context.CancelFunc
	gameData            IGameData
	controller          IController
	userManager         IUserManager
	mapManager          IMapManager
	menuFactory         IMenuFactory
	effectFactory       IEffectFactory
	clanManager         *objects.ClansManager
	itemSellListManager IItemSellListManager
}

func (this *ContextImpl) GetItemSellListManager() IItemSellListManager {
	return this.itemSellListManager
}

func (this *ContextImpl) ClansManager() IClansManager {
	return this.clanManager
}

func (this *ContextImpl) GetEffectFactory() IEffectFactory {
	return this.effectFactory
}

func (this *ContextImpl) GetMenuFactory() IMenuFactory {
	return this.menuFactory
}

func (this *ContextImpl) GetMapManager() IMapManager {
	return this.mapManager
}

func (this *ContextImpl) GetUserManager() IUserManager {
	return this.userManager
}

func (this *ContextImpl) Init() error {
	logging.Logger.Info("Initializing app context")
	return nil
}

func (this *ContextImpl) CloseGoroutines() {
	this.cancel()
	logging.Logger.Info("Closed all goroutines")
}

func (this *ContextImpl) GetConfig() *config.AppConfig {
	return this.config
}

func (this *ContextImpl) GetDatabase() IDatabase {
	return this.database
}

func (this *ContextImpl) GetContext() context.Context {
	return this.ctx
}

func (this *ContextImpl) GetHooks() IHooks {
	return this.hooks
}

func (this *ContextImpl) GetNetworkLoop() INetworkLoop {
	return this.netWork
}

func (this *ContextImpl) Close() error {
	err2 := this.netWork.Close()
	if err2 != nil {
		logging.Logger.Error("Error closing network: ", zap.Error(err2))
	}

	err1 := this.database.Close()
	if err1 != nil {
		logging.Logger.Error("Error closing database: ", zap.Error(err1))
	}
	return nil
}

func (this *ContextImpl) GetController() IController {
	return this.controller
}

func (this *ContextImpl) GetGameData() IGameData {
	return this.gameData
}

func (this *ContextImpl) GameConfig() *config.GameConfig {
	return this.gameCfg
}

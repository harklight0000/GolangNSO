package objects

import (
	. "nso/ainterfaces"
	"nso/entity"
	. "nso/logging"
)

func NewGameData(ctx IAppContext) IGameData {

	this := &GameData{
		appContext: ctx,
	}
	this.LoadData()
	this.LoadGame()
	createCahe(this)
	LoadCache()
	this.itemFactory = NewItemFactory(this.items)
	this.clansManager = NewClansManager()
	return this
}

type GameData struct {
	appContext   IAppContext
	maps         []*entity.MapEntity
	npcs         []*entity.NPCEntity
	mobs         []*entity.MobEntity
	optionItems  []*entity.OptionItemEntity
	items        []*entity.ItemEntity
	njParts      []*entity.NjPartEntity
	tasks        []*entity.TasksEntity
	levels       []*entity.LevelEntity
	effects      []*entity.EffectEntity
	itemSells    map[int]*entity.ItemSellEntity
	GNpcs        []NpcInfo
	clanManagers []*ClanManager
	gMaps        []*IMap
	*LevelManager
	*GameScr
	skills       []*entity.SkillEntity
	itemFactory  IITemFactory
	clansManager IClansManager
}

func (this *GameData) Effects() []*entity.EffectEntity {
	return this.effects
}

func (this *GameData) Npcs() []NpcInfo {
	return this.GNpcs
}

func (this *GameData) LoadGame() {
	db := this.appContext.GetDatabase()
	Logger.Info("Start loading game data")
	this.loadClanManager(db)
	this.loadItemShinwa(db)
	this.loadGameNpc()
	this.updateCaveIDAll(db)
	Logger.Info("Load game data successfully")
}

func (this *GameData) LoadData() {
	db := this.appContext.GetDatabase()
	Logger.Info("Start loading game data")
	this.loadMap(db)
	this.loadNpc(db)
	this.loadMob(db)
	this.loadOptionItem(db)
	this.loadItem(db)
	this.loadSkills(db)
	this.loadNjPart(db)
	this.loadTask(db)
	this.loadLevel(db)
	this.loadEffect(db)
	this.loadItemSell(db)
	this.LevelManager = NewLevelManager(this.levels)
	this.GameScr = NewGameScr()

	Logger.Info("Load game data successfully")
}

func (this *GameData) MapEntities() []*entity.MapEntity {
	return this.maps
}

func (this *GameData) loadItemShinwa(db IDatabase) {
	// TODO: load item shinwa
}

func (this *GameData) updateCaveIDAll(db IDatabase) {
	// TODO: update cave id all to -1
}

func (this *GameData) Mobs() []*entity.MobEntity {
	return this.mobs
}

func (this *GameData) GetItemFactory() IITemFactory {
	return this.itemFactory
}

func (this *GameData) GetClansManager() IClansManager {
	return this.clansManager
}

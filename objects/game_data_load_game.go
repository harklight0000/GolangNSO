package objects

import (
	"go.uber.org/zap"
	. "nso/ainterfaces"
	"nso/entity"
	. "nso/logging"
)

func (this *GameData) loadClanManager(db IDatabase) {
	var clans []entity.ClanEntity
	err := db.FindAll(&clans)
	if err != nil {
		Logger.Panic("Error when load clan manager", zap.Error(err))
		return
	}
	for _, clan := range clans {
		manager := NewClanManager(&clan, this.appContext)
		er := manager.ParseData()
		if er != nil {
			Logger.Panic("Error when parse clan manager", zap.Error(er))
			return
		}
		this.clanManagers = append(this.clanManagers)
	}
}

func (this *GameData) loadGameNpc() {
	for _, npc := range this.npcs {
		n := NpcInfo{
			NPCEntity: *npc,
		}
		er := n.ParseData()
		if er != nil {
			Logger.Panic("Error when parse npc", zap.Error(er))
			return
		}
		this.GNpcs = append(this.GNpcs, n)
	}
}

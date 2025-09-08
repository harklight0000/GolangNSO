package ainterfaces

import (
	"nso/entity"
)

type IGameData interface {
	MapEntities() []*entity.MapEntity
	LoadData()
	LoadGame()
	Mobs() []*entity.MobEntity
	Npcs() []NpcInfo
	Effects() []*entity.EffectEntity
	GetItemFactory() IITemFactory
}

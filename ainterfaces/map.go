package ainterfaces

import (
	"time"
)

type IMap interface {
	IGameObject
	SetTimeMap(timeMap time.Time)
	Close() error
	GetTemplateMob() []MobInfo
	GetID() int
	GetCave() ICave
	GetXHD() int
	IsVdmq() bool
	GetFreeArea() IArea
	Npcs() []*NpcInfo
	IsGtcMap() bool
}

type ICave interface {
	UpdateXP(xp int64)
	UpdatePoint(point int)
	Reset()
	Finish()
	GetMaps() []IMap
}

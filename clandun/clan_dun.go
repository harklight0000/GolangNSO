package clandun

import (
	. "nso/ainterfaces"
	. "nso/objects"
	"nso/utils"
	"sync"
)

type IClanDunState interface {
	CanEnter() bool
	ShouldLeave() bool
	ChangeState(state IClanDunState)
}

type IClanDun interface {
	IGameObject
	InitMaps()
	GetTick() int64
	SetTick(tick int64)
	GetDuration() int64
	GetClanManager() IClanManager
	SetDuration(duration int64)
	GetEntrance() IMap
	GetCurrentState() IClanDunState
	EnterEntrance(user *User)
	GetMemClanData(id int) *MemberDunData
	Inform(msg string)
}

type ClanDun struct {
	*ClanManager
	baseClanID *utils.AtomicInteger
	maps       sync.Map
}

func NewClanDun(clanManager *ClanManager) *ClanDun {
	this := &ClanDun{ClanManager: clanManager}
	this.baseClanID = utils.NewAtomicInteger(1)

	return this
}

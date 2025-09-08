package ainterfaces

import "nso/core"

type IParty interface {
	GetMaster() INinja
	IsLocked() bool
	GetCave() ICave
	GetID() int
	GetMapID() int
	GetNinjaByID(id int) INinja
	OpenCave(cave ICave)
	RefreshTeam()
	SendText(text string)
	ChatMessage(ninja INinja, text string)
	GetNinjas() []INinja
	GetBattle() IBattle
	SetBattle(battle IBattle)
	UpdateEffect(effect IEffect)
	GetName() string
	ChangeTeamLeader(index int)
	AddToParty(ninja INinja)
	AcceptParty(ninja INinja)
	RemoveMember(ninja INinja)
	ExitParty(ninja INinja)
	EnterArea(area IArea)
	ClearBattle()
	SendToAll(message *core.Message)
	Size() byte
}

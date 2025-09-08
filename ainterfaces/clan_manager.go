package ainterfaces

import "nso/core"

type IClanManager interface {
	GetName() string
	UpdateCoin(coin int)
	GetTocPho() ClanMember
	GetTocTruong() ClanMember
	GetMembers() []ClanMember
	GetElders() []ClanMember
	GetMemberByID(cID int) ClanMember
	GetMemberByName(name string) ClanMember
	GetMemberMax() int
	GetExpNext() int
	GetFreeCoin() int
	GetCoinOpen() int
	GetCoinUp() int
	ISender
	PayFeeClan()
	UpExp(exp int)
	AddItem(item IItem)
	RemoveItem(item IItem, quantity int)
	Chat(user IUser, message string)
	ChangeTypeClan(user IUser, m core.Message)
	OpenItemLevel(user IUser)
	SendClanItem(user IUser, m core.Message)
	SetAlert(user IUser, m core.Message)
	MoveOutClan(user IUser, m core.Message)
	CreateMessage(m string)
	OutClan(user IUser)
	ClanUpLevel(user IUser)
	InputCoinClan(user IUser, m core.Message)
	InformAll(message string)
	WriteLog(name string, num int, number int, date string)
	LogClan(p IUser)
	RequestClanInfo(user IUser)
	RequestClanMember(user IUser)
	RequestClanItem(user IUser)
	GetLevel() int
	SetLevel(level int)
	InviteToDun(name string)
	SetClanTerritory(clanTerritory interface{})
	SetClanBattle(clanBattle interface{})
	GetClanTerritory() interface{}
	GetClanBattle() interface{}
	SetClanBattleData(clanBattleData interface{})
	GetClanBattleData() interface{}
	RestoreClanBattle() bool
	GetCurrentThanThu() interface{}
	GetThanThuIndex() int
	SetThanThuIndex(index int)
	Close() error
}

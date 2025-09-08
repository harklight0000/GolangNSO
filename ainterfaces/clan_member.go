package ainterfaces

type ClanMember struct {
	CharID        int
	CName         string
	ClanName      string
	TypeClan      int8
	CLevel        int
	NClass        byte
	PointClan     int
	PointClanWeek int
	Ninja         INinja
}

func NewClanMember(clanName string, ninja INinja) *ClanMember {
	this := &ClanMember{ClanName: clanName, Ninja: ninja}
	this.TypeClan = -1
	this.ClanName = clanName
	this.CName = ninja.GetName()
	this.CharID = ninja.GetID()
	this.NClass = ninja.NClass()
	this.CLevel = ninja.GetLevel()
	return this
}

package objects

import (
	. "nso/constants"
	"nso/core"
)

func RefreshMob(receiver *Area, mob *Mob) {
	m := core.NewMessage(NPC_LIVE)
	m.WriteIByte(mob.ID)
	m.WriteByte(mob.Sys)
	m.WriteByte(mob.LevelBoss)
	m.WriteInt(mob.HpMax)
	receiver.LoopAll(func(user *User) {
		user.SendMessage(m)
	})
}

func CancelTrade(u *User) {
	m := core.NewMessage(TRADE_CANCEL)
	u.SendMessage(m)
}

func SendMapInfo(u *User, this *Area) {
	CancelTrade(u)
	ma := this.Map
	m := core.NewMessage(MAP_INFO)
	m.WriteIByte(ma.GetID())
	m.WriteIByte(ma.TileID)
	m.WriteIByte(ma.BgID)
	m.WriteIByte(ma.TypeMap)
	m.WriteUTF(ma.Name)
	m.WriteIByte(this.ID)
	m.WriteShort(u.Get().(*Ninja).X)
	m.WriteShort(u.Get().(*Ninja).Y)
	vgo := ma.Vgos
	m.WriteIByte(len(vgo))
	for _, v := range vgo {
		m.WriteIShort(v.MinX)
		m.WriteIShort(v.MinY)
		m.WriteIShort(v.MaxX)
		m.WriteIShort(v.MaxY)
	}
	m.WriteIByte(len(this.mobs))
	for _, mo := range this.mobs {
		m.WriteBool(mo.IsDisable)
		m.WriteBool(mo.IsDontMove)
		m.WriteBool(mo.IsFire)
		m.WriteBool(mo.IsIce)
		m.WriteBool(mo.IsWind)
		m.WriteIByte(mo.TemplateId)
		m.WriteByte(mo.Sys)
		m.WriteInt(mo.Hp)
		m.WriteIByte(mo.Level)
		m.WriteInt(mo.HpMax)
		m.WriteShort(mo.X)
		m.WriteShort(mo.Y)
		m.WriteByte(mo.Status)
		m.WriteByte(mo.LevelBoss)
		m.WriteBool(mo.IsBoss)
	}
	// Khuc go
	m.WriteByte(0)
	m.WriteIByte(len(ma.npcs))
	for _, npc := range ma.npcs {
		m.WriteIByte(npc.NPCEntity.Type)
		m.WriteShort(npc.X)
		m.WriteShort(npc.Y)
		m.WriteIByte(npc.NPCEntity.ID)
	}
	this.lockItemMap.RLock()
	defer this.lockItemMap.RUnlock()
	m.WriteIByte(len(this.itemMaps))
	for _, im := range this.itemMaps {
		m.WriteShort(im.ItemMapID)
		m.WriteShort(im.Item.ID)
		m.WriteShort(im.X)
		m.WriteShort(im.Y)
	}
	m.WriteUTF(this.Map.Name)
	m.WriteIByte(0)
	u.SendMessage(m)
}

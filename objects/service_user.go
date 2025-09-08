package objects

import (
	. "nso/constants"
	"nso/core"
)

func (this *User) SendCharInfo(nj *Ninja, other *User) {
	m := core.NewMessage(PLAYER_ADD)
	m.WriteInt(nj.GetID())
	m.WriteUTF(nj.clan.ClanName)
	if nj.clan.ClanName != "" {
		m.WriteSByte(nj.clan.TypeClan)
	}
	// Visible
	m.WriteBool(false)
	m.WriteByte(nj.GetTypePk())
	m.WriteByte(nj.NClass())
	m.WriteSByte(nj.Gender)
	m.WriteShort(nj.PartHead())
	m.WriteUTF(nj.Name)
	m.WriteInt(nj.hp)
	m.WriteInt(nj.GetMaxHp())
	m.WriteIByte(nj.GetLevel())
	m.WriteShort(nj.Weapon())
	m.WriteShort(nj.PartBody())
	m.WriteShort(nj.PartLeg())
	m.WriteSByte(-1) // Mob me
	m.WriteShort(nj.X)
	m.WriteShort(nj.Y)
	m.WriteShort(nj.GetEff5BuffHP())
	m.WriteShort(nj.GetEff5BuffMP())
	m.WriteByte(0) // Effect
	m.WriteBool(nj.IsHuman)
	m.WriteBool(!nj.IsHuman)
	m.WriteShort(nj.PartHead())
	m.WriteShort(nj.Weapon())
	m.WriteShort(nj.PartBody())
	m.WriteShort(nj.PartLeg())
	it := nj.ItemBody[18] // Dau chan than
	if it != nil {
		switch it.ID {
		case 795: // Thien nguyet chi nu
			m.WriteShort(37)
			m.WriteShort(38)
			m.WriteShort(39)
		case 796: // Nhat tu lam phong
			m.WriteShort(40)
			m.WriteShort(41)
			m.WriteShort(42)
		case 804: // Hjiro
			m.WriteShort(58)
			m.WriteShort(59)
			m.WriteShort(60)
		case 805: // Shiraiji
			m.WriteShort(55)
			m.WriteShort(56)
			m.WriteShort(57)
		case 850: // Mat na ho
			m.WriteShort(69 - int16(nj.Gender*3))
			m.WriteShort(70 - int16(nj.Gender*3))
			m.WriteShort(71 - int16(nj.Gender*3))
		default:
			m.WriteShort(-1)
			m.WriteShort(-1)
			m.WriteShort(-1)
		}
	} else {
		m.WriteShort(-1)
		m.WriteShort(-1)
		m.WriteShort(-1)
	}
	it = nj.ItemBody[17]
	if it != nil {
		if it.ID == 799 { // Gay mat trang
			m.WriteShort(44)
		} else if it.ID == 800 {
			m.WriteShort(46)
		} else {
			m.WriteShort(-1)
		}
	} else {
		m.WriteShort(-1)
	}

	it = nj.ItemBody[12]
	if it != nil {
		if it.ID == 797 {
			m.WriteShort(43)
		} else {
			m.WriteShort(-1)
		}
	} else {
		m.WriteShort(-1)
	}
	m.WriteShort(-1)
	it = nj.ItemMounts[4] // thu cuoi
	if it != nil {
		if it.ID == 798 { // Lan su vu
			m.WriteShort(36)
		} else if it.ID == 801 {
			m.WriteShort(47)
		} else if it.ID == 802 {
			m.WriteShort(48)
		} else if it.ID == 803 {
			m.WriteShort(49)
		} else if it.ID == 839 {
			m.WriteShort(63)
		} else if it.ID == 851 {
			m.WriteShort(72)
		} else {
			m.WriteShort(-1)
		}
	} else {
		m.WriteShort(-1)
	}
	m.WriteShort(-1)
	it = nj.ItemBody[27] // Mat na
	if it != nil {
		if it.ID == 813 { // Mat na shin ah
			m.WriteShort(54)
		} else if it.ID == 814 { // Mat na vo dien
			m.WriteShort(53)
		} else if it.ID == 815 { // Mat na oni
			m.WriteShort(52)
		} else if it.ID == 816 { // Mat na kuma
			m.WriteShort(51)
		} else if it.ID == 817 { // Mat na inu
			m.WriteShort(50)
		} else {
			m.WriteShort(-1)
		}
	} else {
		m.WriteShort(-1)
	}
	// Bien hinh
	it = nj.ItemBody[26]
	if it != nil {
		if it.ID == 825 { // Pet bong ma
			m.WriteShort(61)
		} else if it.ID == 826 { // Pt yeu tinh
			m.WriteShort(62)
		} else if it.ID == 852 {
			m.WriteShort(74)
		} else {
			m.WriteShort(-1)
		}
	} else {
		m.WriteShort(-1)
	}
	for i := 16; i < 32; i++ {
		it = nj.ItemBody[i]
		if it != nil {
			m.WriteShort(it.ID)
			m.WriteByte(it.GetUpgrade())
			m.WriteByte(it.Sys)
		} else {
			m.WriteShort(-1)
		}
	}
	other.SendMessage(m)
	if nj.mobMe != nil {
		m = core.MessageSubCommand(PLAYER_LOAD_THU_NUOI)
		m.WriteInt(nj.ID)
		m.WriteIByte(nj.mobMe.Template.ID)
		m.WriteBool(nj.mobMe.IsBoss)
		other.SendMessage(m)
	}
}

func (this *User) SendCoat(other *User) {
	m := core.MessageSubCommand(PLAYER_LOAD_AO_CHOANG)
	nj := this.Get().(*Ninja)
	if nj.ItemBody[TYPE_YOROI] == nil {
		return
	}
	m.WriteInt(nj.ID)
	m.WriteInt(nj.hp)
	m.WriteInt(nj.GetMaxHp())
	m.WriteShort(nj.ItemBody[TYPE_YOROI].ID)
	other.SendMessage(m)
}

func (this *User) SendGlove(u *User) {
	m := core.MessageSubCommand(PLAYER_LOAD_GLOVE)
	nj := this.Get().(*Ninja)
	if nj.ItemBody[TYPE_GIA_TOC] == nil {
		return
	}
	m.WriteInt(nj.ID)
	m.WriteInt(nj.hp)
	m.WriteInt(nj.GetMaxHp())
	m.WriteShort(nj.ItemBody[TYPE_GIA_TOC].ID)
	u.SendMessage(m)
}

package objects

import (
	. "nso/ainterfaces"
	"nso/constants"
	"nso/core"
)

func RefreshTeam(this *Party) {
	m := core.NewMessage(constants.PLAYER_IN_PARTY)
	m.WriteBool(this.IsLocked())
	for _, n := range this.ninjas {
		m.WriteInt(n.GetID())
		m.WriteByte(n.NClass())
		m.WriteUTF(n.GetName())
	}
	this.SendToAll(m)
}

func ChatParty(this *Party, name string, text string) {
	m := core.NewMessage(constants.CHAT_PARTY)
	m.WriteString(name)
	m.WriteString(text)
	this.SendToAll(m)
}

func SendInvitationParty(this *Party, ninja INinja) {
	m := core.NewMessage(79)
	m.WriteInt(this.master.GetID())
	m.WriteString(this.master.Name)
	ninja.SendMessage(m)
}

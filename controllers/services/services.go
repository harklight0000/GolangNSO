package services

import (
	. "nso/ainterfaces"
	. "nso/constants"
	"nso/core"
	. "nso/objects"
)

func SendMenuArray(sender ISender, menu []string) {
	m := core.NewMessage(OPEN_UI_NEWMENU)
	for _, v := range menu {
		m.WriteUTF(v)
	}
	sender.SendMessage(m)
}

func SendMenuWithChat(sender ISender, npcID int16, menu []string, chat string) {
	m := core.NewMessage(OPEN_UI_CONFIRM)
	m.WriteShort(npcID)
	m.WriteUTF(chat)
	m.WriteIByte(len(menu))
	for _, v := range menu {
		m.WriteUTF(v)
	}
	sender.SendMessage(m)
}

func MoveMessage(nj *Ninja) *core.Message {
	m := core.NewMessage(PLAYER_MOVE)
	m.WriteInt(nj.ID)
	m.WriteShort(nj.X)
	m.WriteShort(nj.Y)
	return m
}
func ResetPoint(char *Ninja) {
	m := core.NewMessage(RESET_POINT)
	m.WriteShort(char.X)
	m.WriteShort(char.Y)
	char.SendMessage(m)
}

func EndLoad(sender ISender, canvas bool) {
	m := core.NewMessage(GIAODO)
	m.WriteBool(!canvas)
	sender.SendMessage(m)
}

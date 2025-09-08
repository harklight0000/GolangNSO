package ainterfaces

import (
	. "nso/constants"
	"nso/core"
)

type MenuInfo struct {
	NpcInfo
	App   IAppContext
	Index int
}

func (this MenuInfo) DefaultMenu(user IUser) []string {
	var menu []string
	for _, talk := range this.Talks {
		menu = append(menu, talk[0])
	}
	return menu
}

func (this MenuInfo) Chat(user IUser, text string) {
	m := core.NewMessage(OPEN_UI_SAY)
	m.WriteShort(int16(this.ID))
	m.WriteUTF(text)
	user.SendMessage(m)
}

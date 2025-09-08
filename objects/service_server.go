package objects

import (
	. "nso/ainterfaces"
	"nso/constants"
	"nso/core"
)

func SendTB(sender ISender, title string, s string) {
	m := core.NewMessage(constants.ALERT_MESSAGE)
	m.WriteUTF(title)
	m.WriteUTF(s)
	sender.SendMessage(m)
}

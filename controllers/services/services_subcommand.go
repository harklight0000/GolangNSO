package services

import (
	. "nso/ainterfaces"
	. "nso/constants"
	"nso/core"
)

func SendSkills(ss ISession, textSkill string, length int, skill []byte) {
	m := core.MessageSubCommand(SEND_SKILL)
	m.WriteUTF(textSkill)
	m.WriteInt(length)
	m.WriteFull(skill)
	m.WriteByte(0)
	ss.SendMessage(m)
}

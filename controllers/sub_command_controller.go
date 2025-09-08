package controllers

import (
	"fmt"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"
	. "nso/ainterfaces"
	. "nso/constants"
	"nso/controllers/services"
	"nso/core"
	"nso/logging"
	. "nso/objects"
	. "nso/utils"
)

func (this *Controller) MessageSubCommand(session ISession, message *core.Message) error {
	cmd := MessageSubCommand(message.ReadByte())
	logging.Logger.Info(fmt.Sprintf("MessageSubCommand cmd %d name= %s", cmd, cmd.String()))
	switch cmd {
	case SEND_SKILL:
		return this.sendSkills(session, message)
	case SAVE_RMS:
		return this.saveRMS(session, message)
	case FIND_PARTY:
		return this.findParties(session)
	default:
		this.LogNotImplementedSubCommand(cmd)
	}
	return nil
}

func (this *Controller) LogNotImplementedSubCommand(cmd MessageSubCommand) {
	logging.Logger.Warn("Not implemented sub command", zap.String("sub command", cmd.String()))
}

func (this *Controller) sendSkills(ss ISession, message *core.Message) error {
	user, ok := ss.GetUser().(*User)
	if !ok {
		return eris.New("Session is not supported *User")
	}
	ninja, ok := user.Get().(*Ninja)
	if !ok {
		return eris.New("User is not supported *Ninja")
	}
	skillType := message.ReadUTF()
	var arrSkill []byte
	var length int
	if skillType == "KSkill" {
		length = len(ninja.KSkill)
		arrSkill = make([]byte, length)
		for i, v := range ninja.KSkill {
			arrSkill[i] = Byte(v)
		}
	} else if skillType == "OSkill" {
		length = len(ninja.OSkill)
		arrSkill = make([]byte, length)
		for i, v := range ninja.OSkill {
			arrSkill[i] = Byte(v)
		}
	} else if skillType == "CSkill" {
		length = 1
		arrSkill = make([]byte, length)
		arrSkill[0] = Byte(-1)
		skill := ninja.GetCSkill()
		if skill != nil {
			if skill.Type != 2 {
				arrSkill[0] = skill.ID
			}
		}
		if arrSkill[0] == Byte(-1) && len(ninja.GetSkills()) > 0 {
			arrSkill[0] = ninja.GetSkills()[0].ID
		}
	}
	if arrSkill == nil {
		return eris.New("Cannot get skill of type " + skillType)
	}
	services.SendSkills(ss, skillType, length, arrSkill)
	return nil
}

func (this *Controller) saveRMS(session ISession, message *core.Message) error {
	t1 := message.ReadUTF()
	t2 := message.ReadUTF()
	length := message.ReadShort()
	nj, ok := session.GetUser().(*User).Get().(*Ninja)
	if !ok {
		return eris.New("User is not supported *Ninja")
	}
	logging.Logger.Info("Save RMS", zap.String("t1", t1), zap.String("t2", t2), zap.Int("length", int(length)))
	switch t1 {
	case "KSkill":
		for i := range nj.KSkill {
			id := message.ReadByte()
			if id != -1 {
				skill := nj.GetSkill(Byte(id))
				if skill != nil && skill.Type != 0 {
					nj.KSkill[i] = int8(skill.ID)
				}
			}
		}
	case "OSkill":
		for i := range nj.OSkill {
			id := message.ReadByte()
			if id != -1 {
				skill := nj.GetSkill(Byte(id))
				if skill != nil && skill.Type != 0 {
					nj.OSkill[i] = int8(skill.ID)
				}
			}
		}
	}
	return nil
}

func (this *Controller) findParties(session ISession) error {
	user, ok := session.GetUser().(*User)
	if !ok {
		return eris.New("Session is not supported *User")
	}
	ninja, ok := user.Get().(*Ninja)
	if !ok {
		return eris.New("User is not supported *Ninja")
	}
	if !IsNil(ninja.GetParty()) {
		ninja.SendYellowMessage("Bạn đã có nhóm rồi")
		return nil
	}
	parties := user.Area.GetParties()
	m := core.MessageSubCommand(FIND_PARTY)
	for _, p := range parties {
		master := p.GetMaster().(*Ninja)
		m.WriteByte(master.NClass())
		m.WriteIByte(master.GetLevel())
		m.WriteUTF(master.GetName())
		m.WriteByte(p.Size())
	}
	session.SendMessage(m)
	return nil
}

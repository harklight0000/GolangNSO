package controllers

import (
	. "nso/constants"
	"nso/core"
	. "nso/networking"
	"nso/objects"
)

func (this *Controller) updateData(ss *Session) error {
	m := core.MessageNotMap(UPDATE_DATA)
	m.WriteFull(objects.Cache[objects.DATA_CACHE])
	ss.SendMessage(m)
	return nil
}

func (this *Controller) updateItem(ss *Session) error {
	m := core.MessageNotMap(UPDATE_ITEM)
	m.WriteFull(objects.Cache[objects.ITEM_CACHE])
	ss.SendMessage(m)
	return nil
}

func (this *Controller) updateSkill(ss *Session) error {
	m := core.MessageNotMap(UPDATE_SKILL)
	m.WriteFull(objects.Cache[objects.SKILL_CACHE])
	ss.SendMessage(m)
	return nil
}

func (this *Controller) updateMap(ss *Session) error {
	m := core.MessageNotMap(UPDATE_MAP)
	if ss.IsNew() {
		m.WriteFull(objects.Cache[objects.MAP_CACHE_NEW])
	} else {
		m.WriteFull(objects.Cache[objects.MAP_CACHE])
	}
	ss.SendMessage(m)
	return nil
}

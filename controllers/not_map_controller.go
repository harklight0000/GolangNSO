package controllers

import (
	"fmt"
	"github.com/rotisserie/eris"
	. "nso/constants"
	"nso/core"
	"nso/entity"
	"nso/logging"
	. "nso/networking"
	. "nso/objects"
	. "nso/sqlplugins"
	. "nso/utils"
	"path"
)

func (this *Controller) LogMessageNotMap(cmd MessageNotMap) error {
	logging.Logger.Info(fmt.Sprintf("Not map not implemented command id = %d, name = %s", cmd, cmd.String()))
	return nil
}

func (this *Controller) MessageNotMap(ss *Session, m *core.Message) error {
	cmd := MessageNotMap(m.ReadByte())
	logging.Logger.Info(fmt.Sprintf("MessageNotMap command id = %d, name = %s", cmd, cmd.String()))
	switch cmd {
	case UPDATE_DATA:
		return this.updateData(ss)
	case UPDATE_ITEM:
		return this.updateItem(ss)
	case UPDATE_SKILL:
		return this.updateSkill(ss)
	case UPDATE_MAP:
		return this.updateMap(ss)
	case SELECT_PLAYER:
		return this.selectPlayer(ss, m)
	case CLIENT_OK:
		return this.renderSelectPlayer(ss)
	case REQUEST_MOB_TEMPLATE:
		return this.requestMobTemplate(ss, m)
	case REQUEST_ICON:
		return this.requestIcon(ss, m)
	default:
		return this.LogMessageNotMap(cmd)
	}
}

func (this *Controller) selectPlayer(ss *Session, m *core.Message) error {
	user, ok := ss.User.(*User)
	if !ok {
		return eris.New("User is not supported *User")
	}
	if user == nil {
		return eris.New("User is nil")
	}

	if user.Ninja != nil {
		return eris.New("User is already selected")
	}
	name := m.ReadUTF()
	var contains = false
	for _, n := range user.NinjasArray {
		if n == name {
			contains = true
			break
		}
	}
	if !contains {
		return eris.New("Ninja is not exist in user account")
	}
	for _, nj := range user.NinjaObjects {
		if nj != nil && nj.Name == name {
			user.Ninja = nj
			break
		}
	}
	user.NinjaObjects = nil
	user.Ninja.IsHuman = true
	err := user.Ninja.ParseData()
	if err != nil {
		return eris.Wrap(err, "Error occur when parse ninja data")
	}
	user.Ninja.SendInfo()
	mapManager := ss.AppCtx.GetMapManager()
	foundMap := mapManager.GetMapByID(int(user.Ninja.MapID))
	if foundMap == nil {
		foundMap = mapManager.GetMapByID(22)
	}
	userManager := ss.AppCtx.GetUserManager()
	err = userManager.AddUser(user)
	if err != nil {
		return eris.Wrap(err, "Error occur when add user")
	}
	err = mapManager.Enter(user, foundMap)
	SendTB(ss, "Server", "Số lượng thành viên "+ToString(userManager.Size()))
	_ = this.ChatMap(user, "Code quá 180 phút có hại cho sức khoẻ")
	if err != nil {
		ss.SendServerDialog("Có lỗi khi chuyển map")
		return eris.Wrap(err, "Error occur when enter map")
	}
	err = user.UpdateAsync()
	if err != nil {
		return eris.Wrap(err, "Error occur when update user")
	}
	return nil
}

func (this *Controller) renderSelectPlayer(ss *Session) error {
	user, ok := ss.User.(*User)
	if !ok {
		ss.SendServerDialog("Error occur when render select player")
		return nil
	}
	var count byte
	for _, nj := range user.NinjasArray {
		if nj != "" {
			count++
		}
	}
	m := core.MessageNotMap(SELECT_PLAYER)
	m.WriteByte(count)
	db := ss.AppCtx.GetDatabase()
	user.Ninja = nil
	for _, name := range user.NinjasArray {
		if name == "" {
			continue
		}
		var ninja entity.NinjaEntity
		err := db.FindOne(&ninja, Eq("name", name))
		if err != nil {
			return eris.Wrap(err, "Error occur when parse ninja data")
		}
		nj := NewNinja(&ninja, user)
		err = nj.Awake()
		if err != nil {
			return eris.Wrap(err, "Error occur when awake ninja object")
		}
		err = nj.Start()
		if err != nil {
			return eris.Wrap(err, "Error occur when run update ninja")
		}
		user.NinjaObjects = append(user.NinjaObjects, nj)
		m.WriteSByte(nj.Gender)
		m.WriteUTF(nj.Name)
		m.WriteUTF(NClassToClassName[nj.NClass()])
		m.WriteIByte(nj.Level)
		maskItem := nj.ItemBody[11]
		var head = nj.Head
		if maskItem != nil {
			head = getHeadByItemID(maskItem.ID)
			if head == -1 {
				head = maskItem.GetData().Part
			}
		}

		weaponItem := nj.ItemBody[1]
		var weapon int16 = -1
		if weaponItem != nil {
			weapon = weaponItem.GetData().Part
		}
		var body int16 = -1
		if nj.ItemBody[2] != nil {
			body = nj.ItemBody[2].GetData().Part
		}
		var leg int16 = -1
		if nj.ItemBody[6] != nil {
			leg = nj.ItemBody[6].GetData().Part
		}
		if head == 185 || head == 188 || head == 258 || head == 264 || head == 267 || head == 270 || head == 273 || head == 276 || head == 279 {
			body = head + 1
			leg = head + 2
		}
		m.WriteShort(head)
		m.WriteShort(weapon)
		m.WriteShort(body)
		m.WriteShort(leg)
	}
	ss.SendMessage(m)
	return nil
}

func getHeadByItemID(id int16) int16 {
	switch id {
	case 541:
		return 185
	case 542:
		return 188
	case 745:
		return 264
	case 774:
		return 267
	case 786:
		return 270
	case 787:
		return 276
	case 853:
		return 273
	case 854:
		return 279
	}
	return -1
}

func (this *Controller) requestMobTemplate(ss *Session, m *core.Message) error {
	var id int
	if ss.IsNew() {
		id = int(m.ReadUnsignedShort())
	} else {
		id = int(m.ReadUByte())
	}
	if id < 0 {
		return nil
	}
	data, err := ReadAll(path.Join(this.cfg.ResPath, "cache", "mob", ToString(int(ss.ZoomLevel)), ToString(int(id))))
	if err != nil {
		return eris.Wrap(err, "Error occur when read mob template")
	}
	m = core.NewMessage(NOT_MAP)
	m.WriteFull(data)
	ss.SendMessage(m)
	return nil
}

func (this *Controller) requestIcon(ss *Session, m *core.Message) error {
	id := m.ReadInt()
	data, er := ReadAll(path.Join(this.cfg.ResPath, "icon", ToString(int(ss.ZoomLevel)), ToString(id)+".png"))
	if er != nil {
		return eris.Wrap(er, "Error occur when read icon")
	}
	m = core.MessageNotMap(REQUEST_ICON)
	m.WriteInt(id)
	m.WriteInt(len(data))
	m.WriteFull(data)
	ss.SendMessage(m)
	return nil
}

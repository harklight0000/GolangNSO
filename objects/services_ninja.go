package objects

import (
	"nso/constants"
	"nso/core"
	"nso/logging"
	"nso/networking"
	. "nso/utils"
	"path"
)

func SendLoadAll(ss *networking.Session, nj *Ninja) {
	m := core.MessageSubCommand(constants.ME_LOAD_ALL)
	m.WriteInt(nj.GetID())
	m.WriteUTF(nj.clan.ClanName)
	if nj.clan.ClanName != "" {
		m.WriteSByte(nj.clan.TypeClan)
	}
	m.WriteSByte(nj.TaskID)
	m.WriteSByte(nj.Gender)
	m.WriteShort(nj.Head)
	m.WriteByte(nj.GetSpeed())
	m.WriteUTF(nj.Name)
	m.WriteByte(nj.Pk)
	m.WriteByte(nj.GetTypePk())
	m.WriteInt(nj.GetMaxHp())
	m.WriteInt(nj.hp)
	m.WriteInt(nj.GetMaxMP())
	m.WriteInt(nj.mp)
	m.WriteLong(nj.Exp)
	m.WriteLong(nj.ExpDown)
	m.WriteShort(nj.GetEff5BuffHP())
	m.WriteShort(nj.GetEff5BuffMP())
	m.WriteByte(nj.NClass())
	m.WriteShort(nj.GetPPoint())
	m.WriteShort(nj.Potential0)
	m.WriteShort(nj.Potential1)
	m.WriteInt(nj.Potential2)
	m.WriteInt(nj.Potential3)
	m.WriteShort(nj.SPoint)
	m.WriteIByte(len(nj.skills))
	for _, skill := range nj.skills {
		m.WriteShort(skill.template.SkillId)
	}
	m.WriteInt(nj.Xu)
	m.WriteInt(nj.Yen)
	m.WriteInt(nj.User.Luong)
	m.WriteByte(nj.Maxluggage)
	for i := 0; i < int(nj.Maxluggage); i++ {
		item := nj.ItemBag[i]
		if item != nil {
			m.WriteShort(item.ID)
			m.WriteBool(item.IsLock())
			if item.IsTypeBody() || item.IsTypeMount() || item.IsTypeNgocKham() {
				m.WriteByte(item.GetUpgrade())
			}
			m.WriteBool(item.isExpired)
			m.WriteIShort(item.Quantity)
		} else {
			m.WriteShort(-1)
		}
	}
	for i := 0; i < 16; i++ {
		item := nj.ItemBody[i]
		if item != nil {
			m.WriteShort(item.ID)
			m.WriteByte(item.GetUpgrade())
			m.WriteByte(item.Sys)
		} else {
			m.WriteShort(-1)
		}
	}
	m.WriteBool(nj.IsHuman)
	m.WriteBool(!nj.IsHuman)
	m.WriteSByte(-1)
	m.WriteSByte(-1)
	m.WriteSByte(-1)
	m.WriteSByte(-1)
	ss.SendMessage(m)
}

func SetPointPB(ninja *Ninja, point int16) {
	m := core.MessageNotMap(constants.POINT_PB)
	m.WriteShort(point)
	ninja.SendMessage(m)
}

func removePlayer(u *User) {
	ninja := u.Get().(*Ninja)
	m := core.NewMessage(constants.PLAYER_REMOVE)
	m.WriteInt(ninja.GetID())
	u.Area.SendToAll(m)
}

func addEffectMessage(ninja *Ninja, effect Effect) {
	m := core.MessageSubCommand(constants.ME_ADD_EFFECT)
	m.WriteByte(byte(effect.Template.TemplateID))
	m.WriteInt(effect.TimeStart)
	m.WriteInt(int(effect.TimeRemove - CurrentTimeMillis()))
	m.WriteShort(int16(effect.Param))
	effType := effect.Type
	if effType == 2 || effType == 3 || effType == 14 {
		m.WriteShort(ninja.X)
		m.WriteShort(ninja.Y)
	}
	ninja.SendMessage(m)
	m = core.MessageSubCommand(constants.PLAYER_ADD_EFFECT)
	m.WriteInt(ninja.GetID())
	m.WriteByte(byte(effect.Template.TemplateID))
	m.WriteInt(effect.TimeStart)
	m.WriteInt(int(effect.TimeRemove - CurrentTimeMillis()))
	m.WriteShort(int16(effect.Param))
	effType = effect.Type
	if effType == 2 || effType == 3 || effType == 14 {
		m.WriteShort(ninja.X)
		m.WriteShort(ninja.Y)
	}
	ninja.User.Area.SendToAll(m)
}

func updateEffect(ninja *Ninja, eff Effect) {
	m := core.MessageSubCommand(constants.ME_EDIT_EFFECT)
	m.WriteByte(byte(eff.ID))
	m.WriteInt(eff.TimeStart)
	m.WriteInt(eff.TimeLength)
	m.WriteShort(int16(eff.Param))
	ninja.SendMessage(m)

	m = core.MessageSubCommand(constants.PLAYER_EDIT_EFFECT)
	m.WriteInt(ninja.GetID())
	m.WriteByte(byte(eff.ID))
	m.WriteInt(eff.TimeStart)
	m.WriteInt(eff.TimeLength)
	m.WriteShort(int16(eff.Param))
	ninja.User.Area.SendToAll(m)
}

func removeEffectMessage(this *Ninja, ef Effect) {
	m := core.MessageSubCommand(constants.ME_REMOVE_EFFECT)
	m.WriteIByte(ef.ID)
	typeEff := ef.Type
	if typeEff == 0 || typeEff == 12 {
		m.WriteInt(this.hp)
		m.WriteInt(this.mp)
	} else if typeEff == 4 || typeEff == 13 || typeEff == 17 {
		m.WriteInt(this.hp)
	} else if typeEff == 23 {
		m.WriteInt(this.hp)
		m.WriteInt(this.GetMaxHp())
	}
	this.SendMessage(m)
	m = core.MessageSubCommand(constants.PLAYER_REMOVE_EFFECT)
	m.WriteInt(this.GetID())
	m.WriteIByte(ef.ID)
	if typeEff == 0 || typeEff == 12 {
		m.WriteInt(this.hp)
		m.WriteInt(this.mp)
	} else if typeEff == 4 || typeEff == 13 || typeEff == 17 {
		m.WriteInt(this.hp)
	} else if typeEff == 23 {
		m.WriteInt(this.hp)
		m.WriteInt(this.GetMaxHp())
	}
	this.User.Area.SendToAll(m)
}

func SendYellowMessage(this *Ninja, text string) {
	m := core.NewMessage(constants.SERVER_MESSAGE)
	m.WriteUTF(text)
	this.SendMessage(m)
}

const (
	TYPE_ITEM_SELL      = 2
	TYPE_ITEM_BAG       = 3
	TYPE_MENU_CAI_TRANG = 4
	TYPE_ITEM_BODY      = 5
)

func RequestItemInfoMessage(u *User, message *core.Message) {
	typeUI := message.ReadByte()
	index := message.ReadUByte()
	var item *Item
	nj := u.GetNinja().(*Ninja)
	switch typeUI {
	case TYPE_ITEM_BAG:
		if index < 0 || index >= nj.Maxluggage {
			return
		}
		item = nj.ItemBag[index]
	case TYPE_MENU_CAI_TRANG:
		if u.MenuCaiTrang == 1 {
			item = nj.ItemBST[index]
		} else if u.MenuCaiTrang == 2 {
			item = nj.ItemCaiTrang[index]
		} else {
			item = nj.ItemBox[index]
		}
	case TYPE_ITEM_BODY:
		if index < 0 || index >= 32 {
			return
		}
		item = u.Get().(*Ninja).ItemBody[index]
	case 39:
		clan, ok := u.Session.AppCtx.ClansManager().GetClanByName(nj.clan.ClanName).(*ClanManager)
		if ok && clan != nil {
			items := clan.GetItems()
			if index < 0 || int(index) >= len(items) {
				return
			}
			item = (items[index]).(*Item)
		}
	case 41:
		if index < 0 || int(index) >= len(nj.ItemMounts) {
			return
		}
		item = nj.ItemMounts[index]
	}
	if (typeUI >= 14 && typeUI <= 29) || typeUI == 32 || typeUI == 34 || typeUI == 2 || typeUI == 8 || typeUI == 9 {
		item = u.Session.AppCtx.GetItemSellListManager().GetItemSellByType(int(typeUI)).GetItemByIndex(int(index)).(*Item)
	}

	if item == nil {
		return
	}
	m := core.NewMessage(constants.REQUEST_ITEM_INFO)
	m.WriteSByte(typeUI)
	m.WriteByte(index)
	m.WriteLong(item.Expires)
	if IsTypeUIME(typeUI) {
		m.WriteInt(item.Sale)
	}
	if IsTypeUIShop(typeUI) || IsTypeUIShopLock(typeUI) || IsTypeMounts(typeUI) || IsTypeUIStore(typeUI) || IsTypeUIBook(typeUI) || IsTypeUIFashion(typeUI) || IsTypeUIClanShop(typeUI) {
		m.WriteInt(item.BuyCoin)
		m.WriteInt(item.BuyCoinLock)
		m.WriteInt(item.BuyGold)
	}
	if item.IsTypeBody() || item.IsTypeMount() || item.IsTypeNgocKham() {
		m.WriteByte(item.Sys)
		for _, o := range item.Option {
			m.WriteIByte(o.ID)
			m.WriteInt(o.Param)
		}
		itemType := item.GetData().Type
		if len(item.Ngocs) > 0 {
			var option int
			if itemType == constants.TYPE_VU_KHI {
				option = 106
			} else if item.IsTrangSuc() {
				option = 108
			} else if item.IsTrangPhuc() {
				option = 107
			}
			if option != 0 {
				for _, ngoc := range item.Ngocs {
					indx := indexOfOption(ngoc, option)
					op1 := ngoc.Option[indx+1]
					op2 := ngoc.Option[indx+2]
					if ngoc.ID == constants.HUYEN_TINH_NGOC {
						m.WriteByte(109)
					} else if ngoc.ID == constants.HUYET_NGOC {
						m.WriteByte(110)
					} else if ngoc.ID == constants.LAM_TINH_NGOC {
						m.WriteByte(111)
					} else if ngoc.ID == constants.LUC_NGOC {
						m.WriteByte(112)
					}
					m.WriteInt(0)
					m.WriteIByte(op1.ID)
					m.WriteInt(op1.Param)
					m.WriteIByte(op2.ID)
					m.WriteInt(op2.Param)
				}
			}
		}

	}
	if item.ID == 233 || item.ID == 234 || item.ID == 235 {
		img, err := ReadAll(path.Join(u.Session.AppCtx.GetConfig().ResPath, "icon", "1", "diado.png"))
		if err != nil {
			logging.Logger.Info("Không thấy hình địa đồ")
		}
		m.WriteInt(len(img))
		m.WriteFull(img)
	}
	u.SendMessage(m)
}

func indexOfOption(item Item, id int) int {
	for i, o := range item.Option {
		if o.ID == id {
			return i
		}
	}
	return -1
}

package controllers

import (
	"fmt"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"
	. "nso/ainterfaces"
	"nso/cache"
	"nso/config"
	. "nso/constants"
	. "nso/controllers/services"
	"nso/core"
	"nso/entity"
	"nso/errs"
	"nso/logging"
	. "nso/networking"
	. "nso/objects"
	. "nso/sqlplugins"
	. "nso/utils"
	"time"
)

type Controller struct {
	AppCtx  IAppContext
	cfg     *config.AppConfig
	gameCfg *config.GameConfig
	data    IGameData
}

func NewController(c IAppContext) IController {
	return &Controller{
		AppCtx:  c,
		data:    c.GetGameData(),
		cfg:     c.GetConfig(),
		gameCfg: c.GameConfig(),
	}
}

func (this *Controller) OnMessage(session ISession, m *core.Message) (e error) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				logging.Logger.Error("Error occur when run func ", zap.Error(err.(error)))
				e = eris.Wrap(err.(error), "Error occur when OnMessage of Controller")
			case string:
				logging.Logger.Error("Error occur when run  OnMessage of Controller ", zap.String("error", err.(string)))
				e = eris.New(err.(string) + " when run OnMessage of Controller")
			case fmt.Stringer:
				logging.Logger.Error("Error occur when run  OnMessage of Controller ", zap.String("error", err.(fmt.Stringer).String()))
				e = eris.New(err.(fmt.Stringer).String() + " when run OnMessage of Controller")
			}
		}
	}()
	var user *User
	if session.GetUser() != nil {
		user = session.GetUser().(*User)
	}
	command := m.Command
	ss, ok := session.(*Session)
	if !ok {
		return eris.New("Session is not supported *Session")
	}
	switch command {
	case MESSAGE_SUB_COMMAND:
		return this.MessageSubCommand(ss, m)
	case NOT_LOGIN:
		return this.MessageNotLogin(ss, m, user)
	case NOT_MAP:
		return this.MessageNotMap(ss, m)
	case GET_SESSION_ID:
		session.SendKey()
	case CHAT_MAP:
		return this.ChatMap(user, m.ReadUTF())
	case PLAYER_MOVE:
		return this.MoveMessage(m, user)
	case OPEN_UI_MENU:
		return this.openUIMenu(user, m)
	case MENU:
		return this.menu(user, m)
	case MAP_CHANGE:
		return this.mapChange(user)
	case REQUEST_ITEM_INFO:
		RequestItemInfoMessage(user, m)
	case OPEN_UI_ZONE:
		return this.openUIZone(user)
	case ZONE_CHANGE:
		return this.zoneChange(user, m)
	case SKILL_SELECT:
		return this.skillSelect(user, m)
	case PLAYER_ATTACK_NPC:
		return this.playerAttackNPC(user, m)
	default:
		return this.LogNotImplemented(m.Command)
	}

	return nil
}

func (this *Controller) ChatMap(user *User, text string) error {
	if user == nil {
		return eris.New("User is nil")
	}
	m := core.NewMessage(CHAT_MAP)
	m.WriteInt(user.Ninja.ID)
	m.WriteUTF(text)
	user.Area.SendToAll(m)
	return nil
}

func (this *Controller) Login(session *Session, message *core.Message) error {
	ctx := session.AppCtx
	ctx.GetUserManager()
	name := Escape(message.ReadUTF())
	password := Escape(message.ReadUTF())
	session.Version = message.ReadUTF()
	message.ReadUTF()
	message.ReadUTF()
	message.ReadUTF()
	message.ReadByte()
	var user *entity.PlayerEntity
	db := this.AppCtx.GetDatabase()
	if err := db.FindOne(&user, And(Eq("username", name), Eq("password", password))); err != nil {
		return eris.Wrap(err, "Error occur when login with player username "+name)
	} else {
		u := NewUser(session, user)
		er := u.ParseData()
		if u.Lock == 1 {
			session.SendServerDialog("Tài khoản của bạn đã bị khoá liên hệ ADMIN để biết thêm chi tiết")
			return nil
		}
		if u.KichHoat != 0 {
			session.SendServerDialog("Tài khoản của bạn chưa được kích hoạt, vui lòng kiểm tra lại")
			return nil
		}

		if er != nil {
			return eris.Wrap(er, "Error occur when parse data")
		}
		session.User = u
		session.Login = true
		session.UpdateVersion()
	}
	return nil
}

func (this *Controller) SetClientInfo(ss *Session, m *core.Message) error {
	ss.Type = m.ReadUByte()
	ss.ZoomLevel = m.ReadUByte()
	ss.IsGps = m.ReadBoolean()
	ss.Width = m.ReadInt()
	ss.Height = m.ReadInt()
	ss.IsQwerty = m.ReadBoolean()
	ss.IsTouch = m.ReadBoolean()
	ss.Platform = m.ReadUTF()
	m.ReadInt()
	m.ReadByte()
	ss.LanguageID = m.ReadUByte()
	ss.Provider = m.ReadInt()
	ss.Agent = m.ReadUTF()
	this.ClearRMS(ss)
	return nil
}

func (this *Controller) ClearRMS(ss *Session) {
	//m := core.MessageNotLogin(CLEAR_RMS)
	//ss.SendText(m)
}

func (this *Controller) LogNotImplemented(command Command) error {
	logging.Logger.Info("Command not implemented ", zap.String("command", command.String()))
	return nil
}

func (this *Controller) MoveMessage(m *core.Message, user *User) error {
	nj := user.Get().(*Ninja)
	if nj.HasEffect(EFF_LOCKED) {
		logging.Logger.Info("Player is locked ", zap.String("player", user.Username))
		return nil
	}
	x, y := m.ReadShort(), m.ReadShort()
	user.LastX = x
	user.LastY = y
	nj.X = x
	nj.Y = y

	m = MoveMessage(nj)
	user.Area.LoopAll(func(user *User) {
		user.SendMessage(m)
	})
	clone, ok := user.GetClone().(*Ninja)
	if user.IsHuman && ok && clone != nil {
		clone.X = Next[int16](int(nj.X-35), int(nj.X+35))
		clone.Y = Next[int16](int(nj.Y-35), int(nj.Y+35))
		cloneMoveMsg := MoveMessage(clone)
		user.Area.LoopAll(func(user *User) {
			user.SendMessage(cloneMoveMsg)
		})
	}
	return nil
}

func (this *Controller) openUIMenu(user *User, m *core.Message) error {
	idNpc := m.ReadShort()

	logging.Logger.Info("Open UI menu of npc ", zap.Int("id", int(idNpc)))
	user.CurrentMenuProcessor = this.AppCtx.GetMenuFactory().GetMenuProcessor(idNpc)
	customMenu := user.CurrentMenuProcessor.GetMenu(user)
	if len(customMenu) == 0 {
		customMenu = user.CurrentMenuProcessor.DefaultMenu(user)
	}
	SendMenuArray(user, customMenu)
	return nil
}

const VGO_RANGE = 100

func (this *Controller) mapChange(user *User) error {
	startTime := time.Now()
	nj := user.Get().(*Ninja)
	x := nj.X
	y := nj.Y

	var mapID int
	var vgo *cache.Vgo
	for _, v := range user.Area.Map.Vgos {
		if v.MinX <= int(x)+VGO_RANGE &&
			v.MaxX >= int(x)-VGO_RANGE &&
			v.MinY <= int(y)+VGO_RANGE &&
			v.MaxY >= int(y)-VGO_RANGE {
			mapID = v.MapID
			vgo = &v
			break
		}
	}
	if vgo == nil {
		if x < 100 {
			nj.X = user.LastX + 50
			nj.Y = user.LastY
		} else {
			nj.X = user.LastX - 100
			nj.Y = user.LastY
		}
		logging.Logger.Debug("Reset point ", zap.String("player", user.Username))
		ResetPoint(nj)
		return nil
	}
	ma := this.AppCtx.GetMapManager().GetMapByID(mapID)
	cave, ok := ma.GetCave().(*Cave)
	if ok && cave != nil {
		for _, _map := range ma.GetCave().GetMaps() {
			_map := _map.(*Map)
			if _map.ID == mapID {
				ma = _map
				break
			}
		}
	}
	for _, item := range nj.ItemMounts {
		if item != nil && item.IsExpired() {
			return errs.NewErrNextMap("Trang bị thú cưới đã hết hạn. Vui lòng tháo ra để di chuyển")
		}
	}
	if user.Area.Map.IsLdgtMap() {
		if user.ClanTeritoryID == 0 {
			return errs.NewErrNextMap("Không thể đi tiếp")
		}
	}
	// TODO: Check has party enter the same area
	user.Leave()
	area := ma.GetFreeArea().(*Area)
	if area == nil {
		return errs.NewErrNextMap("Không thể đi tiếp")
	}
	nj.X = int16(vgo.GoX)
	nj.Y = int16(vgo.GoY)
	clone, ok := user.GetClone().(*Ninja)
	if ok && clone != nil {
		clone.X = int16(NextInt1(int(nj.X-35), int(nj.X+35)))
		clone.Y = nj.Y
	}
	area.Enter(user)
	logging.Logger.Info(fmt.Sprintf("Elapsed %d", time.Now().Sub(startTime).Nanoseconds()))
	return nil
}

func (this *Controller) menu(user *User, m *core.Message) error {
	var typeClose int8 = -1
	if user.Session.IsNew() {
		typeClose = m.ReadByte()
		logging.Logger.Info("Close menu type ", zap.Int("type", int(typeClose)))
	}
	npcId := m.ReadByte()
	menuID := m.ReadByte()
	optionID := m.ReadByte()
	this.AppCtx.GetMenuFactory().
		GetMenuProcessor(int16(npcId)).
		Process(user, *NewMenuOptions(menuID, optionID, typeClose))
	return nil
}

func (this *Controller) openUIZone(user *User) error {
	nj := user.Get().(*Ninja)
	if !nj.IsAlive() {
		return errs.NewErrNextMap("Bạn đang chết")
	}
	isNearNpcChangeMap := false
	for _, npc := range user.Area.Npcs() {
		if npc.ID == 13 && npc.IsNear(nj) {
			isNearNpcChangeMap = true
			break
		}
	}
	if !isNearNpcChangeMap {
		user.Session.SendServerDialog("Bạn phải gần cột chuyển khu mới có thể chuyển được")
	}
	if isNearNpcChangeMap || nj.QuantityItemTotal(37) > 0 || nj.QuantityItemTotal(35) > 0 {
		m := core.NewMessage(OPEN_UI_ZONE)
		m.WriteByte(user.Area.Map.Size())
		var areas = user.Area.Map.GetAreas()
		for i := 0; i < len(areas); i++ {
			m.WriteByte(areas[i].(*Area).NumPlayers())
			m.WriteByte(areas[i].(*Area).NumParties())
		}
		user.SendMessage(m)
	} else {
		user.Session.SendServerDialog("Bạn không thể chuyển được")
	}
	return nil
}

func (this *Controller) zoneChange(user *User, m *core.Message) error {
	defer EndLoad(user, true)
	zoneID := m.ReadByte()
	itemIndex := m.ReadByte()
	var item *Item
	if itemIndex >= 0 && int(itemIndex) < len(user.Ninja.ItemBag) {
		item = user.Ninja.ItemBag[itemIndex]
	}
	isNearNpcChangeMap := false
	for _, npc := range user.Area.Npcs() {
		if npc.ID == 13 && npc.IsNear(user.Get()) {
			isNearNpcChangeMap = true
			break
		}
	}
	areas := user.Area.Map.GetAreas()
	if (isNearNpcChangeMap || (item != nil && (item.ID == 35 || item.ID == 37))) &&
		zoneID >= 0 && int(zoneID) < len(areas) {
		if areas[zoneID].IsFree() {
			areas[zoneID].Enter(user)
			if item != nil && item.ID == 35 {
				user.RemoveItemBag(byte(itemIndex))
			}
		} else {
			user.Session.SendServerDialog("Khu vực đã đầy")
		}
	}
	CancelTrade(user)
	return nil
}

func (this *Controller) skillSelect(user *User, m *core.Message) error {
	return nil
}

func (this *Controller) playerAttackNPC(user *User, m *core.Message) error {

	return nil
}

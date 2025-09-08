package objects

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	. "nso/ainterfaces"
	. "nso/constants"
	"nso/core"
	"nso/logging"
	"nso/utils"
	"sync"
	"time"
)

const MAX_PARTY = 6

type Area struct {
	ID           int
	Map          *Map
	appCtx       IAppContext
	numTa        int
	numTL        int
	numMobDie    int
	ctx          context.Context
	cancel       context.CancelFunc
	users        map[int]*User
	lUser        sync.RWMutex
	events       chan Event
	mobs         []*Mob
	levelToMobId map[int]int
	lockItemMap  sync.RWMutex
	itemMaps     []*ItemMap
	parties      map[int]IParty
	LParties     sync.RWMutex
}

func (this *Area) IsFree() bool {
	return len(this.users) <= this.Map.MaxPlayer
}

func (this *Area) SendMessage(message *core.Message) {
	for _, m := range this.users {
		if m != nil {
			m.SendMessage(message)
		}
	}
}

func NewArea(appCtx IAppContext, ID int, gameMap IMap) *Area {
	this := &Area{ID: ID, Map: gameMap.(*Map), appCtx: appCtx}
	ctx, cancel := context.WithCancel(appCtx.GetContext())
	this.ctx = ctx
	this.cancel = cancel
	this.levelToMobId = make(map[int]int)
	this.users = make(map[int]*User)
	return this
}

func (this *Area) Awake() error {
	this.initMobs()
	return nil
}

func (this *Area) Start() error {
	return nil
}

func (this *Area) UpdateAsync() error {
	timeUpdate := this.appCtx.GetConfig().UpdateMapDelayMillis
	go func() {
		for {
			select {
			case event := <-this.events:
				if event.Func != nil {
					utils.REFunc(event.Func(this), "Error occur when run func")
				}
				if event.Func2 != nil {
					utils.RFunc(event.Func2(this))
				}
			case <-this.ctx.Done():
				return
			default:
				lastTimeUpdate := CurrentTimeMillis()
				utils.REFunc(this.Update, "Error occur when update map")
				elapsedTime := CurrentTimeMillis() - lastTimeUpdate
				if elapsedTime < timeUpdate {
					time.Sleep(time.Duration(timeUpdate-elapsedTime) * time.Millisecond)
				}
			}
		}
	}()
	return nil
}

func (this *Area) Close() error {
	this.cancel()
	logging.Logger.Info("Closed area with ID: ", zap.Int("id", this.ID))
	return nil
}

func (this *Area) Update() error {
	return nil
}

func (this *Area) initMobs() {
	MapTemplateMobs := this.Map.GetTemplateMob()
	mobTemplates := this.appCtx.GetGameData().Mobs()
	this.mobs = make([]*Mob, len(MapTemplateMobs))
	for i, mobInfo := range MapTemplateMobs {
		m := NewMob(i, mobTemplates[mobInfo.ID], mobInfo.Level, this.appCtx.GameConfig(), this)
		this.mobs[i] = m
		m.X = mobInfo.X
		m.Y = mobInfo.Y
		m.Status = mobInfo.Status
		m.LevelBoss = mobInfo.LevelBoss
		m.MapID = this.Map.GetID()
		m.PlaceID = this.ID
		if m.LevelBoss == TA_THU {
			if i%5 == 0 {
				n := m.HpMax * 200
				m.HpMax = n
				m.Hp = n
				m.TaThu = TaThuInfo{
					ID:    m.Template.ID,
					Level: m.Level,
				}
			}
		} else if m.LevelBoss == THU_LINH {
			hp := m.HpMax * 100
			m.HpMax = hp
			m.Hp = hp
		} else if m.LevelBoss == TINH_ANH {
			hp := m.HpMax * 10
			m.HpMax = hp
			m.Hp = hp
		}
		if isLdgtMap(this.Map.GetID()) {
			m.IsRefresh = false
			m.IncHp(-m.HpMax)
		}
		m.IsBoss = mobInfo.IsBoss
		if !m.IsBoss {
			this.levelToMobId[m.Level] = m.Template.ID
		}
		err := m.Awake()
		if err != nil {
			logging.Logger.Error(fmt.Sprintf("Error occur when awake mobInfo of map id = %d", this.Map.GetID()), zap.Error(err))
			return
		}
	}
}

func (this *Area) RefreshBoss() {
	for _, m := range this.mobs {
		if m.Status == STATUS_DIE && m.IsBoss {
			m.Refresh()
		}
	}
}

func (this *Area) checkCleanMob(id int) bool {
	for _, m := range this.mobs {
		if m != nil && !m.IsDie && m.TemplateId == id {
			return false
		}
	}
	return true
}

func (this *Area) getMobByID(id int) *Mob {
	low := 0
	high := len(this.mobs) - 1
	for low <= high {
		mid := (low + high) / 2
		if this.mobs[mid].ID == id {
			return this.mobs[mid]
		} else if this.mobs[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

func (this *Area) Leave(u *User) {
	this.RemoveUser(u)
	removePlayer(u)
}

func (this *Area) AddUser(u *User) {
	this.lUser.Lock()
	defer this.lUser.Unlock()
	this.users[u.ID] = u
}

func (this *Area) Enter(usr IUser) {
	// TODO: check user is in area
	me := usr.(*User)
	if me.Area != nil {
		me.Area.Leave(me)
	}
	this.AddUser(me)
	me.Area = this
	me.MobAtk = -1
	SendMapInfo(me, this)
	this.LoopNotMe(func(other *User) {
		other.SendCharInfo(other.Get().(*Ninja), me)
		me.SendCharInfo(me.Get().(*Ninja), other)
		other.SendCoat(me)
		me.SendCoat(other)
		other.SendGlove(me)
		me.SendGlove(other)
		if other.CloneNinja != nil && other.IsHuman {
			other.SendCharInfo(other.CloneNinja, me)
		}
		if me.CloneNinja != nil && me.IsHuman {
			me.SendCharInfo(me.CloneNinja, other)
		}

	}, me.ID)
}

func (this *Area) LoopAll(f func(user *User)) {
	this.lUser.RLock()
	defer this.lUser.RUnlock()
	for _, u := range this.users {
		if u != nil {
			f(u)
		}
	}
}

func (this *Area) LoopNotMe(f func(other *User), myID int) {
	this.lUser.RLock()
	defer this.lUser.RUnlock()
	for _, u := range this.users {
		if u != nil && u.ID != myID {
			f(u)
		}
	}
}

func (this *Area) Enter0(usr IUser) {
	u := usr.(*User)
	_map := this.Map
	ninja := u.Get().(*Ninja)
	ninja.X = int16(_map.X0)
	ninja.Y = int16(_map.Y0)
	ninja.MapID = byte(_map.ID)
	cloneNinja, ok := u.GetClone().(*Ninja)
	if ok && cloneNinja != nil {
		cloneNinja.X = int16(_map.X0)
		cloneNinja.Y = int16(_map.Y0)
	}
	this.Enter(u)
}

func (this *Area) RemoveUser(u *User) {
	this.lUser.Lock()
	defer this.lUser.Unlock()
	delete(this.users, u.ID)
}

func (this *Area) GetNpcMap(id int16) NpcInfo {
	return *this.Map.npcs[id]
}

func (this *Area) SendToAll(m *core.Message) {
	this.LoopAll(func(u *User) {
		u.SendMessage(m)
	})
}

func (this *Area) addParty(party *Party) {
	this.LParties.Lock()
	defer this.LParties.Unlock()
	this.parties[party.GetID()] = party
}

func (this *Area) removeParty(party *Party) {
	this.LParties.Lock()
	defer this.LParties.Unlock()
	delete(this.parties, party.GetID())
}

func (this *Area) GetParties() []IParty {
	this.LParties.RLock()
	defer this.LParties.RUnlock()
	var parties []IParty
	for _, party := range this.parties {
		parties = append(parties, party)
	}
	return parties
}

func (this *Area) Npcs() []*NpcInfo {
	return this.Map.Npcs()
}

func (this *Area) NumPlayers() byte {
	this.lUser.RLock()
	defer this.lUser.RUnlock()
	return byte(len(this.users))
}

func (this *Area) NumParties() byte {
	this.LParties.Lock()
	defer this.LParties.Unlock()
	return byte(len(this.parties))
}

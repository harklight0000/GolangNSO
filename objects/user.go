package objects

import (
	. "nso/ainterfaces"
	"nso/constants"
	"nso/core"
	"nso/entity"
	"nso/logging"
	"nso/networking"
	"nso/sqlplugins"
	"nso/utils"
	"time"

	"github.com/rotisserie/eris"
	"go.uber.org/zap"
)

func NewUser(ISession ISession, entity *entity.PlayerEntity) *User {
	this := &User{Session: ISession.(*networking.Session), PlayerEntity: entity}
	this.Ninja = nil
	this.IsHuman = true
	return this
}

type User struct {
	*entity.PlayerEntity
	Session              *networking.Session
	ChatKtgDelay         int64
	TypeMenu             int
	Ninja                *Ninja
	NinjaObjects         []*Ninja
	IsHuman              bool
	Area                 *Area
	CloneNinja           *Ninja
	MobAtk               int
	CurrentMenuProcessor IMenu
	LastX                int16
	LastY                int16
	MenuCaiTrang         int
	lastHash             []byte
}

func (this *User) GetNinja() INinja {
	return this.Ninja
}

func (this *User) Close() error {
	return this.Sync()
}

func (this *User) UpdateAsync() error {
	if this.Ninja == nil {
		return eris.New("Ninja must be selected  before running update")
	}
	go func() {
		err := this.Update()
		if err != nil {
			logging.Logger.Error("Update user error", zap.Error(err))
		}
	}()
	return nil
}

func (this *User) Update() error {
	for {
		select {
		case <-this.Session.Ctx.Done():
			logging.Logger.Info("User disconnected", zap.String("name", this.Username))
			err := this.Sync()
			if err != nil {
				logging.Logger.Error("Sync user error", zap.Error(err))
			}
			return nil
		case <-time.After(time.Second * 15):
			utils.REFunc(this.Sync, "Error sync user")
		}
	}
}

func (this *User) Sync() error {
	var err []error
	if this.Ninja != nil {
		err = append(err, this.Ninja.Sync())
	}
	if this.CloneNinja != nil {
		err = append(err, this.CloneNinja.Sync())
	}
	newHash := utils.Hash(this.PlayerEntity)
	db := this.Session.AppCtx.GetDatabase()
	if !utils.Compare(newHash, this.lastHash) {
		this.lastHash = newHash
		logging.Logger.Info("Sync player", zap.String("name", this.Username))
		err = append(err, db.Update(this.PlayerEntity.TableName(), this.PlayerEntity, sqlplugins.Eq("id", this.ID)))
	}
	return combineErrors(err)
}

func (this *User) SendMessage(message *core.Message) {
	this.Session.SendMessage(message)
}

func (this *User) GetID() int {
	return this.ID
}

func (this *User) RequestItem(itemType int) {
	// TODO: Request Item UI
}

func (this *User) Leave() {
	if this.Area != nil {
		this.Area.Leave(this)
	}
}

func (this *User) GetClone() INinja {
	if this.CloneNinja == nil {
		// Load ninja
		db := this.Session.AppCtx.GetDatabase()
		var ninja *entity.NinjaEntity
		err := db.FindOneOfTable("clone_ninja", &ninja, sqlplugins.Eq("name", this.Ninja.Name))
		if err != nil {
			logging.Logger.Panic("Can not load clone ninja ", zap.String("name", this.Ninja.Name))
		}
		this.CloneNinja = NewNinja(ninja, this)
		this.CloneNinja.IsHuman = false
		err = this.CloneNinja.ParseData()
		if err != nil {
			logging.Logger.Panic("Can not parse clone ninja ", zap.String("name", this.Ninja.Name))
		}
	}
	return this.CloneNinja
}

func (this *User) GetHuman() INinja {
	return this.Ninja
}

func (this *User) Get() INinja {
	if this.IsHuman {
		return this.Ninja
	} else {
		panic("Clone not supported")
	}
}

func (this *User) ParseData() error {
	if len(this.NinjasArray) == 0 {
		this.NinjasArray = make([]string, 3)
	}
	this.ChatKtgDelay = 0
	this.Luong = 0
	this.TypeMenu = -1
	return nil
}

func (this *User) RemoveItemBag(index byte) {
	removedItem := this.GetItemBag(index)
	if removedItem != nil {
		m := core.NewMessage(constants.ITEM_USE_UPTOUP)
		m.WriteByte(index)
		m.WriteIShort(removedItem.Quantity)
		this.SendMessage(m)
		this.Ninja.ItemBag[index] = nil
	}
}

func (this *User) GetItemBag(index byte) *Item {
	if index < 0 || int(index) >= len(this.Ninja.ItemBag) {
		return nil
	}
	return this.Ninja.ItemBag[index]
}

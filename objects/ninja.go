package objects

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rotisserie/eris"
	. "nso/ainterfaces"
	"nso/cache"
	"nso/core"
	"nso/entity"
	"nso/logging"
	"nso/networking"
	"nso/sqlplugins"
	"nso/utils"
	"time"
)

func NewNinja(ninja *entity.NinjaEntity, user *User) *Ninja {
	this := &Ninja{}
	this.NinjaEntity = ninja
	this.User = user
	ss := this.User.Session
	if ss != nil {
		this.Session = ss
		this.AppCtx = ss.AppCtx
		this.data = ss.AppCtx.GetGameData().(*GameData)
	}
	this.effects = make(map[int]*Effect)
	this.ctx, this.cancel = context.WithCancel(user.Session.Ctx)
	return this
}

type Ninja struct {
	*entity.NinjaEntity
	User         *User
	mapLTD       byte
	clan         ClanMember
	Session      *networking.Session
	AppCtx       IAppContext
	IsHuman      bool
	data         *GameData
	skills       []*Skill
	kSkill       []*Skill
	hp           int
	mp           int
	ItemBag      []*Item
	ItemBox      []*Item
	ItemBody     []*Item
	ItemMounts   []*Item
	ItemCaiTrang []*Item
	ItemBST      []*Item
	ItemBodyHide []*Item

	isTaskDanhVong    bool
	countTaskDanhVong int
	useDanhVongPhu    int
	pointUyDanh       int
	pointNon          int
	pointVuKhi        int
	pointAo           int
	PointLien         int
	pointGangTay      int
	pointNhan         int
	pointQuan         int
	pointNgocBoi      int
	pointGiay         int
	pointPhu          int
	pointTinhTu       int
	MapID             byte
	X                 int16
	Y                 int16
	mapType           int16
	mobMe             *Mob
	Die               bool
	NjEffects
	cSkill      *Skill
	cacheSkills map[byte]*Skill
	party       *Party
	ctx         context.Context
	cancel      context.CancelFunc
	lastHash    []byte
}

func (this *Ninja) GetX() int16 {
	return this.X
}

func (this *Ninja) GetY() int16 {
	return this.Y
}

func (this *Ninja) Awake() error {
	return this.ParseData()
}

func (this *Ninja) Start() error {
	return this.UpdateAsync()
}

func (this *Ninja) UpdateAsync() error {
	go func() {
		utils.REFunc(this.Update, "Error update ninja")
	}()
	return nil
}

func (this *Ninja) Update() error {
	for {
		select {
		case <-this.ctx.Done():
			return nil
		case <-time.After(time.Millisecond * 500):
			utils.REFunc(this.update, fmt.Sprintf("Error running update ninja name = %s", this.Name))
		}
	}
}

func (this *Ninja) update() error {
	//logging.Logger.Info("Update ninja")
	return nil
}

func (this *Ninja) Close() error {
	return this.Sync()
}

func (this *Ninja) IsAlive() bool {
	return !this.Die
}

func (this *Ninja) GetParty() IParty {
	return this.party
}

func (this *Ninja) ParseData() error {
	var skillsJSON []SkillJSON

	err := json.Unmarshal([]byte(this.Skill_s), &skillsJSON)
	if err != nil {
		logging.Logger.Panic("Error when parse skill json " + this.Skill_s)
	}
	this.skills = make([]*Skill, len(skillsJSON))
	this.cacheSkills = make(map[byte]*Skill)
	for i, skillJSON := range skillsJSON {
		s := NewSkill(this.data.skills[skillJSON.ID], skillJSON.Point)
		this.skills[i] = s
		this.cacheSkills[s.ID] = this.skills[i]
	}
	logging.Logger.Info("Parse skill success")
	itemFactory := this.data.GetItemFactory()
	parseItem(itemFactory, this.Item_Bag, &this.ItemBag, int(this.Maxluggage))
	parseItem(itemFactory, this.Item_Box, &this.ItemBox, 30)
	parseItem(itemFactory, this.Item_Body, &this.ItemBody, 32)
	parseItem(itemFactory, this.Item_Mounts, &this.ItemMounts, 5)
	parseItem(itemFactory, this.Item_CaiTrang, &this.ItemCaiTrang, 18)
	parseItem(itemFactory, this.Item_BST, &this.ItemBST, 18)
	parseItem(itemFactory, this.Item_BodyHide, &this.ItemBodyHide, 10)
	logging.Logger.Info("Parse item success")
	var effects []*Effect
	err = json.Unmarshal([]byte(this.Effect), &effects)
	if err != nil {
		logging.Logger.Panic("Error when parse effect json " + this.Effect)
	}
	this.effects = make(map[int]*Effect, len(effects))
	for _, effect := range effects {
		this.effects[effect.Template.TemplateID] = effect
		this.effects[effect.Template.TemplateID].EffectEntity = this.data.effects[effect.Template.TemplateID]
		effect.ID = effect.Template.TemplateID
	}
	if this.IsHuman {

		if len(this.TaskDanhVong) == 0 {
			this.TaskDanhVong = []int{
				-1, -1, -1, -1, -1, -1,
			}
		}
		this.useDanhVongPhu = this.TaskDanhVong[5]
		if this.TaskDanhVong[3] == 0 {
			this.isTaskDanhVong = false
		} else {
			this.isTaskDanhVong = true
		}
		this.countTaskDanhVong = this.TaskDanhVong[4]
		this.pointUyDanh = this.CharInfo[0]
		this.pointNon = this.CharInfo[1]
		this.pointVuKhi = this.CharInfo[2]
		this.pointAo = this.CharInfo[3]
		this.PointLien = this.CharInfo[4]
		this.pointGangTay = this.CharInfo[5]
		this.pointNhan = this.CharInfo[6]
		this.pointQuan = this.CharInfo[7]
		this.pointNgocBoi = this.CharInfo[8]
		this.pointGiay = this.CharInfo[9]
		this.pointPhu = this.CharInfo[10]
		this.pointTinhTu = this.CharInfo[11]
		this.MapID = byte(this.Site[0])
		this.X = this.Site[1]
		this.Y = this.Site[2]
		this.mapLTD = byte(this.Site[3])
		if this.mapLTD == 138 {
			this.mapLTD = 22
		}
		this.mapType = this.Site[4]
		if len(this.Tasks) != 2 {
			this.Tasks = make([]*entity.TaskOrder, 2)
		}
		if len(this.Clan) != 2 {
			var data []interface{}
			err := json.Unmarshal([]byte(this.Clan), &data)
			if err != nil {
				logging.Logger.Info("Error when parse clan json " + this.Clan)
			}
			this.clan = *NewClanMember(data[0].(string), this)
			this.clan.PointClan = cast[int](data[1])
		}
	}

	return nil
}

func (this *Ninja) Sync() error {

	db := this.AppCtx.GetDatabase()
	if db == nil {
		return eris.New("db is nil")
	}

	var skillsJSON []SkillJSON
	for _, skill := range this.skills {
		skillsJSON = append(skillsJSON, SkillJSON{
			ID:    skill.ID,
			Point: skill.Point,
		})
	}

	skillsJSONBytes, err := json.Marshal(skillsJSON)
	if err != nil {
		return eris.Wrap(err, "Error when marshal skill json")
	}
	this.Skill_s = string(skillsJSONBytes)
	itemFactory := this.data.GetItemFactory()
	this.Item_Bag = toJSONItems(itemFactory, this.ItemBag)
	this.Item_Box = toJSONItems(itemFactory, this.ItemBox)
	this.Item_Body = toJSONItems(itemFactory, this.ItemBody)
	this.Item_Mounts = toJSONItems(itemFactory, this.ItemMounts)
	this.Item_CaiTrang = toJSONItems(itemFactory, this.ItemCaiTrang)
	this.Item_BST = toJSONItems(itemFactory, this.ItemBST)
	this.Item_BodyHide = toJSONItems(itemFactory, this.ItemBodyHide)

	var effects []*Effect
	for _, effect := range this.effects {
		if effect.IsPermanentEffect() && !effect.IsExpired() {
			effects = append(effects, effect)
		}
	}
	effectsJSON, err := json.Marshal(effects)
	if err != nil {
		return eris.Wrap(err, "Error when marshal effect json")
	}
	this.Effect = string(effectsJSON)
	this.Clan = fmt.Sprintf("[\"%s\",%d]", this.clan.ClanName, this.clan.PointClan)

	if this.IsHuman {
		if len(this.TaskDanhVong) == 0 {
			this.TaskDanhVong = []int{
				-1, -1, -1, -1, -1, -1,
			}
		}
		this.TaskDanhVong[5] = this.useDanhVongPhu
		if this.isTaskDanhVong {
			this.TaskDanhVong[3] = 1
		} else {
			this.TaskDanhVong[3] = 0
		}
		this.TaskDanhVong[4] = this.countTaskDanhVong
		this.CharInfo[0] = this.pointUyDanh
		this.CharInfo[1] = this.pointNon
		this.CharInfo[2] = this.pointVuKhi
		this.CharInfo[3] = this.pointAo
		this.CharInfo[4] = this.PointLien
		this.CharInfo[5] = this.pointGangTay
		this.CharInfo[6] = this.pointNhan
		this.CharInfo[7] = this.pointQuan
		this.CharInfo[8] = this.pointNgocBoi
		this.CharInfo[9] = this.pointGiay
		this.CharInfo[10] = this.pointPhu
		this.CharInfo[11] = this.pointTinhTu
		if this.User.Area != nil && this.User.Area.Map != nil {
			this.Site[0] = int16(this.User.Area.Map.ID)
		} else {
			this.Site[0] = int16(this.MapID)
		}
		this.Site[1] = this.X
		this.Site[2] = this.Y
		this.Site[3] = int16(this.mapLTD)
		this.Site[4] = this.mapType
	}
	var newHash []byte

	if this.IsHuman {
		newHash = utils.Hash(this.NinjaEntity)
	} else {
		newHash = utils.Hash(this.NinjaEntity.BodyEntity)
	}

	if utils.Compare(this.lastHash, newHash) {
		logging.Logger.Info("Not changed")
		return nil
	} else {
		this.lastHash = newHash
	}
	var err2 error
	utils.Bench(func() {
		if this.IsHuman {
			err := db.Update("ninja", this.NinjaEntity, sqlplugins.Eq("id", this.NinjaEntity.ID))
			if err != nil {
				err2 = eris.Wrap(err, "Error when update ninja")
			}
			logging.Logger.Info("Sync ninja success")
		} else {
			err := db.Update("clone_ninja", this.NinjaEntity.BodyEntity, sqlplugins.Eq("id", this.NinjaEntity.ID))
			if err != nil {
				err2 = eris.Wrap(err, "Error when update ninja")
			}
			logging.Logger.Info("Sync clone success")
		}
	})
	return err2
}

func toJSONItems(factory IITemFactory, items []*Item) []cache.ItemJSON {
	var itemsJSON []cache.ItemJSON
	for i, item := range items {
		if item != nil {
			itemsJSON = append(itemsJSON, factory.ToJSON(item, i))
		}
	}
	return itemsJSON
}

func parseItem(itemFactory IITemFactory, input []cache.ItemJSON, output *[]*Item, size int) {
	*output = make([]*Item, size)
	for _, item := range input {
		(*output)[item.Index] = itemFactory.FromJSON(item).(*Item)
	}
}

func (this *Ninja) GetName() string {
	return this.Name
}

func (this *Ninja) GetID() int {
	return this.ID
}

func (this *Ninja) NClass() byte {
	return this.Class
}

func (this *Ninja) GetLevel() int {
	return this.Level
}

func (this *Ninja) SendMessage(message *core.Message) {
	this.Session.SendMessage(message)
}

func (this *Ninja) UpdatePBPoint(point int16) {
	this.PointCave += point
	SetPointPB(this, this.PointCave)
}

func (this *Ninja) UpdateExpNormal(exp int64) {
	// TODO
	logging.Logger.Info("Update exp normal")
}

func (this *Ninja) UpdateExpUseMulti(exp int64) {
	// TODO
	logging.Logger.Info("Update exp use multi")
}

func (this *Ninja) SendInfo() {
	if this.IsHuman {
		this.checkAndResetPPoint()
	}
	this.hp = this.GetMaxHp()
	this.mp = this.GetMaxMP()
	SendLoadAll(this.Session, this)
}

func (this *Ninja) checkAndResetPPoint() {
	totalPpoint := int(this.Potential0) + int(this.Potential1) + this.Potential2 + this.Potential3 + this.PPoint
	if totalPpoint > int(this.data.TotalPPoint(this.GetLevel()))+this.TiemNangSo*10+this.BangHoa*10+25 {
		// TODO :reset point
		this.PPoint = int(this.data.TotalPPoint(this.GetLevel())) + 10*(this.BangHoa+this.TiemNangSo)
	}

}

func (this *Ninja) GetMaxHp() int {
	return 100000
}

func (this *Ninja) GetMaxMP() int {
	return 100000
}

func (this *Ninja) GetSpeed() byte {
	// TODO:
	return this.Speed * 3
}

func (this *Ninja) GetTypePk() byte {
	// TODO:
	return 0
}

func (this *Ninja) GetEff5BuffHP() int16 {
	return 20
}

func (this *Ninja) GetEff5BuffMP() int16 {
	return 20
}

func (this *Ninja) GetPPoint() int16 {
	return int16(this.PPoint)
}

func (this *Ninja) GetCSkill() *Skill {
	if this.cSkill == nil || int(this.cSkill.ID) != this.CSkill {
		for _, skill := range this.skills {
			if int(skill.ID) == this.CSkill {
				this.cSkill = skill
			}
		}
	}
	return this.cSkill
}

func (this *Ninja) GetSkills() []*Skill {
	return this.skills
}

func (this *Ninja) GetSkill(id byte) *Skill {
	return this.cacheSkills[id]
}

func (this *Ninja) AddEffect(id int, timeStart int, timeLength int, param int) {
	effFactory := this.AppCtx.GetEffectFactory()
	effEntity := effFactory.GetTemplateByID(id)
	var eff = this.NjEffects.GetEffectByType(effEntity.Type)
	if eff == nil {
		this.NjEffects.AddEffect(effFactory.CreateEffectFull(id, timeStart, timeLength, param).(*Effect))
		addEffectMessage(this, *eff)
	} else {
		eff.Template = NewTemplate(effEntity)
		eff.TimeLength = timeLength
		eff.TimeStart = timeStart
		eff.Param = param
		eff.TimeRemove = CurrentTimeMillis() - int64(eff.TimeStart+eff.TimeLength)
	}
}

func (this *Ninja) RemoveEffect(id int) {
	eff := this.NjEffects.GetEffectByType(id)
	if eff != nil {
		this.NjEffects.RemoveEffect(id)
		removeEffectMessage(this, *eff)
	}
}

func (this *Ninja) SendYellowMessage(text string) {
	SendYellowMessage(this, text)
}

func (this *Ninja) QuantityItemTotal(id int16) int {
	count := 0
	for _, item := range this.ItemBag {
		if item != nil && item.ID == id {
			count += item.Quantity
		}
	}
	return count
}

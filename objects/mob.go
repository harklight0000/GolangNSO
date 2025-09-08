package objects

import (
	. "nso/ainterfaces"
	"nso/config"
	. "nso/constants"
	"nso/entity"
	"nso/logging"
	. "nso/utils"
	"time"
)

type Mob struct {
	Template       *entity.MobEntity
	TemplateId     int
	IsDisable      bool
	TimeDisable    int64
	IsDontMove     bool
	TimeDontMove   int64
	IsFire         bool
	TimeFire       int64
	IsWind         bool
	TimeWind       int64
	IsThieuDot     bool
	Name           string
	ID             int
	Sys            byte
	Hp             int
	Level          int
	HpMax          int
	X              int16
	Y              int16
	Status         byte
	LvBoss         int
	IsDie          bool
	IsRefresh      bool
	XPUp           int64
	TimeRefresh    int64
	TimeFight      int64
	PlaceID        int
	MapID          int
	NFight         map[int]int
	MasterThieuDot INinja
	Area           *Area
	Map            IMap
	// TODO Gio keo
	AttackCount AtomicInteger
	cfg         *config.GameConfig
	IsBoss      bool
	LevelBoss   byte
	TaThu       TaThuInfo
	IsIce       bool
}

func (this *Mob) Update() error {
	return nil
}

func (this *Mob) CheckFight(ninjaId int) bool {
	return false
}

func NewMob(id int, template *entity.MobEntity, level int, gameConfig *config.GameConfig, area *Area) *Mob {
	this := &Mob{}
	this.ID = id
	this.Template = template
	this.Level = level
	this.Hp = template.Hp
	this.HpMax = template.Hp
	this.XPUp = 100000
	this.cfg = gameConfig
	this.TemplateId = template.ID
	this.IsRefresh = true
	this.IsDie = false
	this.Area = area
	this.Map = area.Map
	this.MapID = area.Map.GetID()
	return this
}

func (this *Mob) Awake() error {
	this.NFight = make(map[int]int)
	return nil
}

func (this *Mob) Start() error {
	return nil
}

func (this *Mob) UpdateAsync() error {
	return nil
}

func (this *Mob) Close() error {
	return nil
}

func (this *Mob) IncHp(num int) {
	this.AttackCount.IncAndGet()
	this.Hp += num
	if this.Hp <= 0 {
		if this.Template.ID == GIO_KEO_DEN_ID || this.Template.ID == GIO_KEO_TRANG_ID {
			this.Hp = this.HpMax
			this.IsDie = false
			return
		}
		this.Hp = this.HpMax
		this.Status = 0
		this.IsDie = true
		if this.IsRefresh {
			this.TimeRefresh = nextTime(time.Duration(this.cfg.TimeRefreshMob) * time.Second)
		}
		if this.IsBoss {
			if this.TemplateId != LAO_DAI && this.TemplateId != LAO_TAM && this.TemplateId != LAO_NHI {
				this.IsRefresh = false
				this.TimeRefresh = -1
			} else if this.TemplateId == HOP_BI_AN {
				this.IsRefresh = true
				this.Hp = 100
				this.TimeRefresh = nextTime(time.Second)
			} else if this.TemplateId == KING_HEO {
				this.IsRefresh = true
				this.Hp = 300
				this.TimeRefresh = nextTime(time.Second)
			} else if this.TemplateId == CHUOT_CANH_TI {
				this.IsRefresh = true
				this.Hp = 10000
				this.TimeRefresh = nextTime(time.Hour)
			} else {
				this.TimeRefresh = nextTime(time.Second)
			}
		}
	}
}

func (this *Mob) ClearFight() {
	this.NFight = make(map[int]int)
}

func (this *Mob) Fight(ninjaId int, dame int) {
	if dame <= 0 {
		logging.Logger.Info("Mob fight but dame <= 0")
		return
	}
	oldDame, ok := this.NFight[ninjaId]
	if !ok {
		this.NFight[ninjaId] = dame
	} else {
		this.NFight[ninjaId] = oldDame + dame
	}
}

func (this *Mob) Refresh() {

	this.ClearFight()
	this.Sys = Next[byte](1, 3)
	if !isCaveMap(this.MapID) &&
		this.LevelBoss != TA_THU &&
		!this.IsBoss &&
		this.Map.GetID() != 74 &&
		this.Map.GetID() != 78 {

		if this.LevelBoss > QUAI_THUONG {
			this.LevelBoss = QUAI_THUONG
		}
		if isLdgtMap(this.MapID) {
			if this.Template.ID != 81 && this.Area.checkCleanMob(this.TemplateId) {
				this.Level = 1
			}
		} else if this.Level > 10 && this.cfg.PercentTaTl >= NextInt2(MAX_PERCENT) && this.Area.numTa < 2 && this.Area.numTL < 1 {
			this.LevelBoss = Next[byte](1, 2)
		}
	}

	cave := this.Map.GetCave().(*Cave)
	if isCaveMap(this.MapID) && cave.finish > 0 && this.Map.GetXHD() == 6 {
		hpUp := this.Template.Hp * (10*cave.finish + MAX_PERCENT) / MAX_PERCENT
		this.HpMax = hpUp
		this.Hp = hpUp
	} else {
		this.HpMax = this.Template.Hp
		this.Hp = this.Template.Hp
	}
	if this.LevelBoss == TA_THU {
		this.Hp = this.HpMax * MAX_PERCENT * 2
		this.HpMax = this.Hp
	} else if this.LevelBoss == THU_LINH {
		this.Hp = this.HpMax * MAX_PERCENT
		this.HpMax = this.Hp
	} else if this.LevelBoss == TINH_ANH {
		this.Hp = this.HpMax * MAX_PERCENT / 10
		this.HpMax = this.Hp
	}
	this.Status = STATUS_ALIVE
	this.IsDie = false
	this.TimeRefresh = 0
	RefreshMob(this.Area, this)
}

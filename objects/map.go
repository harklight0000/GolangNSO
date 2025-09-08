package objects

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"
	. "nso/ainterfaces"
	"nso/cache"
	"nso/entity"
	"nso/logging"
	"nso/utils"
	"time"
)

func NewMap(app IAppContext, id int) *Map {
	ctx, cancel := context.WithCancel(app.GetContext())
	mapEntity := app.GetGameData().MapEntities()[id]
	this := &Map{
		appCtx:    app,
		ctx:       ctx,
		cancel:    cancel,
		MapEntity: mapEntity,
		nArea:     mapEntity.NumZone,
	}
	return this
}

type Map struct {
	*entity.MapEntity
	appCtx  IAppContext
	ctx     context.Context
	timeMap time.Time
	cave    *Cave
	nArea   int
	cancel  context.CancelFunc
	areas   []IArea
	mobs    []MobInfo
	npcs    []*NpcInfo
	Vgos    []cache.Vgo
}

func (this *Map) Npcs() []*NpcInfo {
	return this.npcs
}

func (this *Map) SetNpcs(npcs []*NpcInfo) {
	this.npcs = npcs
}

func (this *Map) IsGtcMap() bool {
	return this.ID >= 118 && this.ID <= 124
}

func (this *Map) GetFreeArea() IArea {
	for _, a := range this.areas {
		if a.IsFree() {
			return a
		}
	}
	return nil
}

func (this *Map) IsVdmq() bool {
	return this.ID >= 139 && this.ID <= 148
}

func (this *Map) GetCave() ICave {
	return this.cave
}

func (this *Map) GetTemplateMob() []MobInfo {
	return this.mobs
}

func (this *Map) SetTimeMap(timeMap time.Time) {
	this.timeMap = timeMap
}

func NewCaveMap(app IAppContext, id int, cave *Cave) *Map {
	this := NewMap(app, id)
	this.cave = cave
	this.nArea = 1
	return this
}

func (this *Map) Awake() error {
	err := this.ParseData()
	if err != nil {
		return eris.Wrap(err, "Error occur when parse data at awake")
	}
	return nil
}

func (this *Map) Start() error {
	//logging.Logger.Info("Awake map name", zap.String("name", this.Name))
	for i := 0; i < this.nArea; i++ {
		area := NewArea(this.appCtx, i, this)
		this.areas = append(this.areas, area)
		err := area.Awake()
		if err != nil {
			return eris.Wrap(err, "Error occur when awake area")
		}
		err = area.Start()
		if err != nil {
			return eris.Wrap(err, "Error occur when start area")
		}

	}
	return nil
}

func (this *Map) UpdateAsync() (e error) {
	appCtx := this.appCtx
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(fmt.Sprintf("Error occur when update async: %s", err))
			}
		}()
		for {
			select {
			case <-this.ctx.Done():
				return
			case <-time.After(time.Duration(appCtx.GetConfig().UpdateMapDelayMillis) * time.Second):
				utils.REFunc(this.Update, "Error occur when update map")
			}
		}
	}()
	return nil
}

func (this *Map) Close() error {
	this.cancel()
	var errs []error
	for _, a := range this.areas {
		err := a.Close()
		if err != nil {
			logging.Logger.Error(fmt.Sprintf("Error close map with id = %d", this.ID), zap.Error(err))
			errs = append(errs, err)
		}
	}
	return combineErrors(errs)
}

func (this *Map) ParseData() (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = eris.New(fmt.Sprintf("Error occur when parse data: %s", err))
		}
	}()
	var mobs [][]interface{}
	err := json.Unmarshal([]byte(this.Mob), &mobs)
	if err != nil {
		return eris.Wrap(err, "Error occur when parse mob")
	}
	for _, m := range mobs {
		this.mobs = append(this.mobs, MobInfo{
			ID:        cast[int16](m[0]),
			Level:     cast[int](m[1]),
			X:         cast[int16](m[2]),
			Y:         cast[int16](m[3]),
			Status:    cast[byte](m[4]),
			LevelBoss: cast[byte](m[5]),
			IsBoss:    m[6].(bool),
		})
	}
	var npcs [][]interface{}
	data := this.appCtx.GetGameData().(*GameData)
	err = json.Unmarshal([]byte(this.NPC), &npcs)
	if err != nil {
		return eris.Wrap(err, "Error occur when parse npc")
	}
	for _, n := range npcs {
		_type := cast[int](n[0])
		x := cast[int16](n[1])
		y := cast[int16](n[2])
		id := cast[int16](n[3])
		npc := data.GNpcs[id]
		npc.ID = byte(id)
		npc.X = x
		npc.Y = y
		npc.Type = _type
		this.npcs = append(this.npcs, &npc)
	}
	for _, v := range this.Vgo {
		v := cache.Vgo{
			MinX:  v[0],
			MinY:  v[1],
			MaxX:  v[2],
			MaxY:  v[3],
			MapID: v[4],
			GoX:   v[5],
			GoY:   v[6],
		}
		if v.MaxX == -1 {
			v.MaxX = v.MinX + 24
		}
		if v.MaxY == -1 {
			v.MaxY = v.MinY + 24
		}
		this.Vgos = append(this.Vgos, v)
	}
	return nil
}

func (this *Map) Update() error {
	return nil
}

func (this *Map) GetID() int {
	return this.ID
}

func (this *Map) GetXHD() int {
	if this.ID == 157 || this.ID == 158 || this.ID == 159 {
		return 9
	}
	if this.ID == 125 || this.ID == 126 || this.ID == 127 || this.ID == 128 {
		return 7
	}
	if this.ID == 114 || this.ID == 115 || this.ID == 116 {
		return 6
	}
	if this.ID == 105 || this.ID == 106 || this.ID == 107 || this.ID == 108 || this.ID == 109 {
		return 5
	}
	if this.ID == 94 || this.ID == 95 || this.ID == 96 || this.ID == 97 {
		return 4
	}
	if this.ID == 91 || this.ID == 92 || this.ID == 93 {
		return 3
	}
	return -1
}

func (this *Map) IsLdgtMap() bool {
	return this.ID >= 80 && this.ID <= 90
}

func (this *Map) Size() byte {
	return byte(len(this.areas))
}

func (this *Map) GetAreas() []IArea {
	return this.areas
}

package objects

import (
	. "nso/ainterfaces"
	"nso/logging"
	"nso/utils"
	"sync"
	"time"
)

var caves sync.Map

var baseCaveID utils.AtomicInteger

type Cave struct {
	CaveID int
	Time   time.Time
	Level  int
	finish int
	Rest   bool
	maps   []IMap
	ninjas []INinja
	X      int
}

func (c *Cave) GetMaps() []IMap {
	return c.maps
}

func (c *Cave) UpdatePoint(point int) {
	logging.Logger.Info("Update point in cave")
	for _, nj := range c.ninjas {
		if nj != nil {
			nj.UpdatePBPoint(int16(point))
		}
	}
}

func (c *Cave) Reset() {
	//TODO implement me
	panic("implement me")
}

func (c *Cave) Finish() {
	//TODO implement me
	panic("implement me")
}

func (c *Cave) UpdateXP(xp int64) {
	logging.Logger.Info("Update XP in cave")
	for _, nj := range c.ninjas {
		nj.UpdateExpUseMulti(xp)
	}
}

func NewCave(ctx IAppContext, x int) *Cave {
	this := &Cave{}
	this.Level = 0
	this.finish = 0
	this.X = x
	this.Rest = false
	this.CaveID = baseCaveID.IncAndGet()
	this.Time = time.Now().Add(time.Hour)
	switch x {
	case 3:
		this.maps = []IMap{
			NewCaveMap(ctx, 91, this),
			NewCaveMap(ctx, 92, this),
			NewCaveMap(ctx, 93, this),
		}
	case 4:
		this.maps = []IMap{
			NewCaveMap(ctx, 94, this),
			NewCaveMap(ctx, 95, this),
			NewCaveMap(ctx, 96, this),
			NewCaveMap(ctx, 97, this),
		}
	case 5:
		this.maps = []IMap{
			NewCaveMap(ctx, 105, this),
			NewCaveMap(ctx, 106, this),
			NewCaveMap(ctx, 107, this),
			NewCaveMap(ctx, 108, this),
			NewCaveMap(ctx, 109, this),
		}
	case 6:
		this.maps = []IMap{
			NewCaveMap(ctx, 114, this),
			NewCaveMap(ctx, 115, this),
			NewCaveMap(ctx, 116, this),
		}
	case 7:
		this.maps = []IMap{
			NewCaveMap(ctx, 125, this),
			NewCaveMap(ctx, 126, this),
			NewCaveMap(ctx, 127, this),
			NewCaveMap(ctx, 128, this),
		}
	case 9:
		this.maps = []IMap{
			NewCaveMap(ctx, 157, this),
			NewCaveMap(ctx, 158, this),
			NewCaveMap(ctx, 159, this),
		}
	default:
		logging.Logger.Panic("Unknow type of x ")
	}
	for _, mapGame := range this.maps {
		mapGame.SetTimeMap(this.Time)
	}
	caves.Store(this.CaveID, this)
	return this
}

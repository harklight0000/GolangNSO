package objects

import (
	"go.uber.org/zap"
	. "nso/ainterfaces"
	"nso/entity"
	"nso/logging"
	. "nso/utils"
)

type MapManager struct {
	maps []*Map
}

func NewMapManager(ctx IAppContext, maps []*entity.MapEntity) *MapManager {
	this := &MapManager{}
	for i := 0; i < len(maps); i++ {
		m := NewMap(ctx, i)
		err := m.Awake()
		if err != nil {
			logging.Logger.Info("Error occur when awake map: ", zap.Int("id", maps[i].ID))
		}
		err = m.Start()
		if err != nil {
			logging.Logger.Info("Error occur when start map: ", zap.Int("id", maps[i].ID))
		}
		this.maps = append(this.maps, m)
	}
	return this
}

func (m *MapManager) GetMapByID(id int) IMap {
	return m.maps[id]
}

func (m *MapManager) Enter(user IUser, dst IMap) error {
	u := user.(*User)
	var isTurn = false
	if dst.GetXHD() != -1 || dst.IsVdmq() {
		isTurn = true
		dst = m.GetMapByID(int(u.GetHuman().(*Ninja).mapLTD))
	}
	area := dst.GetFreeArea()
	if u.Area != nil {
		u.Area.Leave(u)
	}
	if area == nil {
		_map := dst.(*Map)
		_map.areas[NextInt2(len(_map.areas))].(*Area).Enter(u)
	} else {
		u.Area = area.(*Area)
		if !isTurn {
			u.Area.Enter(u)
		} else {
			u.Area.Enter0(u)
		}
	}
	return nil
}

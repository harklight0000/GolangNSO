package ainterfaces

import (
	"encoding/json"
	"fmt"
	"math"
	"nso/entity"

	"github.com/rotisserie/eris"
)

type MobInfo struct {
	ID        int16 `json:"id"`
	X         int16 `json:"x"`
	Y         int16 `json:"y"`
	Status    byte  `json:"status"`
	Level     int   `json:"level"`
	LevelBoss byte  `json:"levelBoss"`
	IsBoss    bool  `json:"isBoss"`
}

type NpcInfo struct {
	ID byte
	X  int16
	Y  int16
	entity.NPCEntity
	Talks [][]string
}

func (n *NpcInfo) ParseData() (er error) {
	defer func() {
		if err := recover(); err != nil {
			er = eris.New(fmt.Sprintf("Error occur when parse data: %s", err))
		}
	}()
	var talks [][]string
	err := json.Unmarshal([]byte(n.NPCEntity.Talk), &talks)
	if err != nil {
		return eris.Wrap(err, "Error when parse talks")
	}
	n.Talks = talks
	return nil
}

func (n *NpcInfo) IsNear(nj INinja) bool {
	return int(math.Abs(float64(n.X-nj.GetX()))) <= 50 && int(math.Abs(float64(n.Y-nj.GetY()))) <= 50
}

type TaThuInfo struct {
	ID    int
	Level int
}

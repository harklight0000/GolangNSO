package constants

import (
	"nso/config"
	"path"
)

var (
	ResCacheMap          string
	ResCacheSkill        string
	ResCacheItem         string
	ResCacheSkillNhanBan string
	ResNjArrow           string
	ResNjEffect          string
	ResNjImage           string
	ResNjPart            string
	ResNjSkill           string

	VsItem        byte
	VsMap         byte
	VsSkill       byte
	VsData        byte
	IsRefreshBoss = []bool{false, false, false, false, false, false}
	MapBossVDMQ   = []int16{141, 142, 143}
	MapBoss45     = []int16{14, 15, 16, 34, 35, 52, 68}
	MapBoss55     = []int16{44, 67}
	MapBoss65     = []int16{24, 41, 45, 59}
	MapBoss75     = []int16{18, 36, 54}
	MapBossMoi    = []int16{161, 162, 163}
	MapBossSK     = []int16{2, 28}
	MapBossLC     = []int16{134, 135, 136, 137}
)

func init() {
	cfg := config.GetAppConfig()
	resPath := cfg.ResPath
	ResCacheMap = path.Join(resPath, "cache", "map")
	ResCacheSkill = path.Join(resPath, "cache", "skill")
	ResCacheItem = path.Join(resPath, "cache", "item")
	ResCacheSkillNhanBan = path.Join(resPath, "cache", "skillnhanban")
	VsItem = cfg.VsItem
	VsMap = cfg.VsMap
	VsSkill = cfg.VsSkill
	VsData = cfg.VsData

	ResNjArrow = path.Join(resPath, "cache", "data", "nj_arrow")
	ResNjEffect = path.Join(resPath, "cache", "data", "nj_effect")
	ResNjImage = path.Join(resPath, "cache", "data", "nj_image")
	ResNjPart = path.Join(resPath, "cache", "data", "nj_part")
	ResNjSkill = path.Join(resPath, "cache", "data", "nj_skill")
}

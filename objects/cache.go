package objects

import (
	"go.uber.org/zap"
	"io/ioutil"
	config2 "nso/config"
	"nso/constants"
	"nso/core"
	"nso/logging"
	. "nso/utils"
	"os"
	"path"
)

var Cache [][]byte

const (
	DATA_CACHE = iota
	MAP_CACHE
	MAP_CACHE_NEW
	SKILL_CACHE
	ITEM_CACHE
	SKILL_NHAN_BAN_CACHE
	size
)

var (
	CacheMapPath     = ""
	CacheMapPathNew  = ""
	CacheSkillPath   = ""
	CacheItemPath    = ""
	CacheDataPath    = ""
	SkillNhanBanPath = ""
)

func init() {
	CacheItemPath = path.Join(config2.GetAppConfig().ResPath, "cache", "item")
	CacheSkillPath = path.Join(config2.GetAppConfig().ResPath, "cache", "skill")
	CacheMapPath = path.Join(config2.GetAppConfig().ResPath, "cache", "map")
	CacheMapPathNew = path.Join(config2.GetAppConfig().ResPath, "cache", "map_new")
	CacheDataPath = path.Join(config2.GetAppConfig().ResPath, "cache", "data.bin")
	SkillNhanBanPath = path.Join(config2.GetAppConfig().ResPath, "cache", "skillnhanban")
}

func createItem(data *GameData) {
	file, err := os.OpenFile(CacheItemPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logging.Logger.Info("Error when creating item cache")
		return
	}
	defer file.Close()
	m := core.NewDataOutputStream(file)
	m.WriteByte(constants.VsItem)
	m.WriteIByte(len(data.optionItems))
	for _, optionItem := range data.optionItems {
		m.WriteUTF(optionItem.Name)
		m.WriteIByte(optionItem.Type)
	}
	m.WriteIShort(len(data.items))
	for _, item := range data.items {
		m.WriteIByte(item.Type)
		m.WriteIByte(item.Gender)
		m.WriteUTF(item.Name)
		m.WriteUTF(item.Description)
		m.WriteByte(item.Level)
		m.WriteShort(item.IconID)
		m.WriteShort(item.Part)
		m.WriteByte(item.UpToUp)
	}
}

func createData(data *GameData) {
	file, err := os.OpenFile(CacheDataPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logging.Logger.Info("Error when creating item cache ", zap.Error(err))
		return
	}
	defer file.Close()
	m := core.NewDataOutputStream(file)
	m.WriteByte(constants.VsData)
	a := ReadAllBytes(constants.ResNjArrow)
	m.WriteInt(len(a))
	m.WriteFull(a)
	a = ReadAllBytes(constants.ResNjEffect)
	m.WriteInt(len(a))
	m.WriteFull(a)
	a = ReadAllBytes(constants.ResNjImage)
	m.WriteInt(len(a))
	m.WriteFull(a)
	a = ReadAllBytes(constants.ResNjPart)
	m.WriteInt(len(a))
	m.WriteFull(a)
	a = ReadAllBytes(constants.ResNjSkill)
	m.WriteInt(len(a))
	m.WriteFull(a)
	m.WriteIByte(len(data.tasks))
	var tasks = data.tasks
	for i, task := range tasks {
		m.WriteIByte(len(task.Tasks))
		for j := range task.Tasks {
			m.WriteSByte(tasks[i].Tasks[j])
			m.WriteSByte(tasks[i].MapTasks[j])
		}
	}
	levels := data.levels
	m.WriteIByte(len(levels))
	for i := range levels {
		m.WriteLong(levels[i].Exps)
	}

	m.WriteIByte(len(data.crystals))
	for i := 0; i < len(data.crystals); i++ {
		m.WriteInt(data.crystals[i])
	}

	m.WriteIByte(len(data.upClothe))
	for i := 0; i < len(data.upClothe); i++ {
		m.WriteInt(data.upClothe[i])
	}

	m.WriteIByte(len(data.upAdorn))
	for i := 0; i < len(data.upAdorn); i++ {
		m.WriteInt(data.upAdorn[i])
	}
	m.WriteIByte(len(data.upWeapon))
	for i := 0; i < len(data.upWeapon); i++ {
		m.WriteInt(data.upWeapon[i])
	}
	m.WriteIByte(len(data.coinUpCrystals))
	for i := 0; i < len(data.coinUpCrystals); i++ {
		m.WriteInt(data.coinUpCrystals[i])
	}
	m.WriteIByte(len(data.coinUpClothes))
	for i := 0; i < len(data.coinUpClothes); i++ {
		m.WriteInt(data.coinUpClothes[i])
	}
	m.WriteIByte(len(data.coinUpAdorns))
	for i := 0; i < len(data.coinUpAdorns); i++ {
		m.WriteInt(data.coinUpAdorns[i])
	}
	m.WriteIByte(len(data.coinUpWeapons))
	for i := 0; i < len(data.coinUpWeapons); i++ {
		m.WriteInt(data.coinUpWeapons[i])
	}
	m.WriteIByte(len(data.goldUps))
	for i := 0; i < len(data.goldUps); i++ {
		m.WriteInt(data.goldUps[i])
	}
	m.WriteIByte(len(data.maxPercents))
	for i := 0; i < len(data.maxPercents); i++ {
		m.WriteInt(data.maxPercents[i])
	}
	effs := data.effects
	m.WriteIByte(len(effs))
	for _, eff := range effs {
		m.WriteIByte(eff.ID)
		m.WriteIByte(eff.Type)
		m.WriteUTF(eff.Name)
		m.WriteShort(eff.IConId)
	}
}

func createMap(data *GameData) {
	file, err := os.OpenFile(CacheMapPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logging.Logger.Info("Error when creating map cache")
		return
	}
	defer file.Close()
	dos := core.NewDataOutputStream(file)
	dos.WriteByte(constants.VsMap)
	dos.WriteIByte(len(data.maps))
	for _, mapData := range data.maps {
		dos.WriteUTF(mapData.Name)
	}
	dos.WriteIByte(len(data.GNpcs))
	for _, npc := range data.GNpcs {
		dos.WriteUTF(npc.Name)
		dos.WriteShort(npc.Head)
		dos.WriteShort(npc.Body)
		dos.WriteShort(npc.Leg)
		dos.WriteIByte(len(npc.Talks))
		for _, talk := range npc.Talks {
			dos.WriteIByte(len(talk))
			for _, menu := range talk {
				dos.WriteUTF(menu)
			}
		}
	}
	// New version using short
	dos.WriteIByte(len(data.mobs))
	for _, mob := range data.mobs {
		dos.WriteByte(mob.Type)
		dos.WriteUTF(mob.Name)
		dos.WriteInt(mob.Hp)
		dos.WriteByte(mob.RangeMove)
		dos.WriteByte(mob.Speed)
	}
}

func createMapNew(data *GameData) {
	file, err := os.OpenFile(CacheMapPathNew, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logging.Logger.Info("Error when creating map cache")
		return
	}
	defer file.Close()
	dos := core.NewDataOutputStream(file)
	dos.WriteByte(constants.VsMap)
	dos.WriteIByte(len(data.maps))
	for _, mapData := range data.maps {
		dos.WriteUTF(mapData.Name)
	}
	dos.WriteIByte(len(data.GNpcs))
	for _, npc := range data.GNpcs {
		dos.WriteUTF(npc.Name)
		dos.WriteShort(npc.Head)
		dos.WriteShort(npc.Body)
		dos.WriteShort(npc.Leg)
		dos.WriteIByte(len(npc.Talks))
		for _, talk := range npc.Talks {
			dos.WriteIByte(len(talk))
			for _, menu := range talk {
				dos.WriteUTF(menu)
			}
		}
	}
	// New version using short
	dos.WriteIShort(len(data.mobs))
	for _, mob := range data.mobs {
		dos.WriteByte(mob.Type)
		dos.WriteUTF(mob.Name)
		dos.WriteInt(mob.Hp)
		dos.WriteByte(mob.RangeMove)
		dos.WriteByte(mob.Speed)
	}
}

func createCahe(data *GameData) {
	createData(data)
	createItem(data)
	createMap(data)
	createMapNew(data)
}

func LoadCache() {
	Cache = make([][]byte, size+1)
	var err error
	Cache[MAP_CACHE], err = ioutil.ReadFile(CacheMapPath)
	if err != nil {
		logging.Logger.Info("Error when reading map cache")
	}
	Cache[MAP_CACHE_NEW], err = ioutil.ReadFile(CacheMapPathNew)
	if err != nil {
		logging.Logger.Info("Error when reading map cache")
	}
	Cache[SKILL_CACHE], err = ioutil.ReadFile(CacheSkillPath)
	if err != nil {
		logging.Logger.Info("Error when reading skill cache")
	}
	Cache[ITEM_CACHE], err = ioutil.ReadFile(CacheItemPath)
	if err != nil {
		logging.Logger.Info("Error when reading item cache")
	}
	Cache[SKILL_NHAN_BAN_CACHE], err = ioutil.ReadFile(SkillNhanBanPath)
	if err != nil {
		logging.Logger.Info("Error when reading skill_nhan_ban cache")
	}
	Cache[DATA_CACHE], err = ioutil.ReadFile(CacheDataPath)
	if err != nil {
		logging.Logger.Info("Error when reading data cache")
	}
}

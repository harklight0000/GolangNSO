package objects

import (
	. "fmt"
	. "nso/ainterfaces"
	"nso/entity"
	. "nso/logging"
	. "nso/sqlplugins"
)

func (this *GameData) loadMap(db IDatabase) {
	// Load maps
	var maps []*entity.MapEntity
	//err := db.FindManyWithOption(&maps, Asc("id"))
	//if err != nil {
	//	Logger.Panic("Error loading maps: " + err.Error())
	//}
	for i := 0; i < 163; i++ {
		var m entity.MapEntity
		err := db.FindOne(&m, Eq("id", i))
		if err != nil {
			Logger.Panic(Sprintf("Error loading map: %d", i))
		}
		maps = append(maps, &m)
	}
	this.maps = maps
	Logger.Info(Sprintf("Loaded %d maps", len(maps)))
}

func (this *GameData) loadNpc(db IDatabase) {
	// Load npcs
	var npcs []*entity.NPCEntity
	err := db.FindManyWithOption(&npcs, Asc("id"))
	if err != nil {
		Logger.Panic("Error loading npcs: " + err.Error())
	}
	this.npcs = npcs
	Logger.Info(Sprintf("Loaded %d npcs", len(npcs)))
}

func (this *GameData) loadMob(db IDatabase) {
	// Load mobs
	var mobs []*entity.MobEntity
	err := db.FindManyWithOption(&mobs, Asc("id"))
	if err != nil {
		Logger.Panic("Error loading mobs: " + err.Error())
	}
	this.mobs = mobs
	Logger.Info(Sprintf("Loaded %d mobs", len(mobs)))
}

func (this *GameData) loadOptionItem(db IDatabase) {
	// Load option item option
	var optionItems []*entity.OptionItemEntity
	err := db.FindManyWithOption(&optionItems, Asc("id"))
	if err != nil {
		Logger.Panic("Error loading option items: " + err.Error())
	}
	Logger.Info(Sprintf("Loaded %d option items", len(optionItems)))
	this.optionItems = optionItems
}

func (this *GameData) loadItem(db IDatabase) {
	// Load item
	var items []*entity.ItemEntity
	err := db.FindManyWithOption(&items, Asc("id"))
	if err != nil {
		Logger.Panic("Error loading items: " + err.Error())
	}
	items = append(items)
	Logger.Info(Sprintf("Loaded %d items", len(items)))
	for _, i := range items {
		err := db.Update(i.TableName(), i, Eq("id", i.ID))
		if err != nil {
			Logger.Panic("Error updating item: id=" + Sprintf("%d", i.ID) + " " + err.Error())
		}
	}
	this.items = items
}

func (this *GameData) loadSkills(db IDatabase) {
	// Load skills
	var skills []*entity.SkillEntity
	err := db.FindManyWithOption(&skills, Asc("id"))
	if err != nil {
		Logger.Panic("Error loading skills: " + err.Error())
	}
	Logger.Info(Sprintf("Loaded %d skills", len(skills)))
	this.skills = skills
}

func (this *GameData) loadNjPart(db IDatabase) {
	var njParts []*entity.NjPartEntity
	err := db.FindManyWithOption(&njParts, Asc("id"))
	if err != nil {
		Logger.Panic("Error loading nj parts: " + err.Error())
	}
	Logger.Info(Sprintf("Loaded %d nj parts", len(njParts)))
	this.njParts = njParts
}

func (this *GameData) loadTask(db IDatabase) {
	//Load tasks
	var tasks []*entity.TasksEntity
	err := db.FindManyWithOption(&tasks, Asc("id"))
	if err != nil {
		Logger.Panic("Error loading tasks: " + err.Error())
	}
	Logger.Info(Sprintf("Loaded %d tasks", len(tasks)))
	this.tasks = tasks
}

func (this *GameData) loadLevel(db IDatabase) {
	//Load level
	var levels []*entity.LevelEntity
	err := db.FindManyWithOption(&levels, Asc("level"))
	if err != nil {
		Logger.Panic("Error loading levels: " + err.Error())
	}
	Logger.Info(Sprintf("Loaded %d levels", len(levels)))
	this.levels = levels

}

func (this *GameData) loadEffect(db IDatabase) {
	// Load effects
	var effects []*entity.EffectEntity
	err := db.FindManyWithOption(&effects, Asc("id"))
	if err != nil {
		Logger.Panic("Error loading effects: " + err.Error())
	}
	Logger.Info(Sprintf("Loaded %d effects", len(effects)))
	this.effects = effects
}

func (this *GameData) loadItemSell(db IDatabase) {
	// Load item sell
	var itemSells []*entity.ItemSellEntity
	err := db.FindManyWithOption(&itemSells, Asc("id"))
	if err != nil {
		Logger.Panic("Error loading item sells: " + err.Error())
	}
	Logger.Info(Sprintf("Loaded %d item sells", len(itemSells)))
	this.itemSells = make(map[int]*entity.ItemSellEntity, len(itemSells))
	for _, itemSell := range itemSells {
		this.itemSells[itemSell.Type] = itemSell
	}
}

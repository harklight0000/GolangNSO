package entity

type NpcDailyEntity struct {
}

func (n NpcDailyEntity) TableName() string {
	return "npc_daily"
}

package cache

type Mob struct {
	ID        int  `json:"id" bson:"id"`
	X         int  `json:"x" bson:"x"`
	Y         int  `json:"y" bson:"y"`
	Status    int  `json:"status" bson:"status"`
	Level     int  `json:"level" bson:"level"`
	LevelBoss int  `json:"levelBoss" bson:"levelBoss"`
	IsBoss    bool `json:"isBoss" bson:"isBoss"`
}

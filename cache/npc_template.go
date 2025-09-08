package cache

type NPC struct {
	Type int8  `json:"type" bson:"type"`
	X    int16 `json:"x" bson:"x"`
	Y    int16 `json:"y" bson:"y"`
	ID   int16 `json:"id" bson:"id"`
}

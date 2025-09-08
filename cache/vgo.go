package cache

type Vgo struct {
	MinX  int `json:"minX" bson:"minX"`
	MinY  int `json:"minY" bson:"minY"`
	MaxX  int `json:"maxX" bson:"maxX"`
	MaxY  int `json:"maxY" bson:"maxY"`
	MapID int `json:"mapID" bson:"mapID"`
	GoX   int `json:"goX" bson:"goX"`
	GoY   int `json:"goY" bson:"goY"`
}

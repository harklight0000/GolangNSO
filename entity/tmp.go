package entity

type TaskOrder struct {
	Count       int    `json:"count" bson:"count"`
	MaxCount    int    `json:"maxCount" bson:"maxCount"`
	TaskId      int    `json:"taskId" bson:"taskId"`
	KillId      int    `json:"killId" bson:"killId"`
	MapID       int    `json:"mapId" bson:"mapId"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type BattleData struct {
	Point int   `json:"point" bson:"point"`
	Phe   int16 `json:"phe" bson:"phe"`
}

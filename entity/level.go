package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type LevelEntity struct {
	_ID    primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	Level  int                `json:"level" bson:"level" gorm:"primary_key;column:level;type:tinyint;default:0"`
	Exps   int64              `json:"exps" bson:"exps" gorm:"column:exps;type:bigint;default:0"`
	Ppoint int                `json:"ppoint" bson:"ppoint" gorm:"column:ppoint;type:tinyint;default:0"`
	Spoint int                `json:"spoint" bson:"spoint" gorm:"column:spoint;type:tinyint;default:0"`
}

func (l LevelEntity) TableName() string {
	return "level"
}

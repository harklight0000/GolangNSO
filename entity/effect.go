package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type EffectEntity struct {
	_ID    primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID     int                `json:"id" gorm:"primary_key;" bson:"id"`
	Type   int                `json:"type" bson:"type" gorm:"column:type;type:tinyint;default:0"`
	Name   string             `json:"name" bson:"name" gorm:"column:name;type:varchar(100);default:''"`
	IConId int16              `json:"iconId" bson:"iconId" gorm:"column:iconId;type:tinyint(6);default:0"`
}

func (e EffectEntity) TableName() string {
	return "effect"
}

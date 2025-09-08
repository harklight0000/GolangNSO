package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nso/cache"
)

type NjPartEntity struct {
	_ID  primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID   int                `json:"id" gorm:"primary_key;" bson:"id"`
	Type int                `json:"type" bson:"type" gorm:"column:type;type:tinyint;default:0"`
	Pi   []cache.PartImage  `json:"pi" bson:"pi" gorm:"column:pi;type:longtext;default:'[]';serializer:my_json"`
}

func (n NjPartEntity) TableName() string {
	return "nj_part"
}

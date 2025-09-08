package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type OptionItemEntity struct {
	_ID  primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID   int                `json:"id" gorm:"primary_key;" bson:"id"`
	Name string             `json:"name" bson:"name" gorm:"column:name;type:varchar(500);default:''"`
	Type int                `json:"type" bson:"type" gorm:"column:type;type:tinyint(6);default:0"`
}

func (o OptionItemEntity) TableName() string {
	return "optionitem"
}

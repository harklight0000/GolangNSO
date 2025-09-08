package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ItemShinwaEntity struct {
	_ID  primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID   int                `json:"id" gorm:"primary_key;" bson:"id"`
	Item string             `json:"item" bson:"item" gorm:"column:item;type:longtext;default:'[]'"`
}

func (i ItemShinwaEntity) TableName() string {
	return "itemshinwa"
}

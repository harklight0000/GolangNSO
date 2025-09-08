package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ItemSellEntity struct {
	_ID      primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID       int                `json:"id" gorm:"primary_key;" bson:"id"`
	Type     int                `json:"type" bson:"type" gorm:"column:type;type:tinyint;default:0"`
	ListItem string             `json:"ListItem" bson:"ListItem" gorm:"column:ListItem;type:mediumtext;default:'[]'"`
}

func (i ItemSellEntity) TableName() string {
	return "itemsell"
}

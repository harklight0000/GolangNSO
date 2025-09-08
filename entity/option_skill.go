package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type OptionSkillEntity struct {
	_ID  primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID   int                `json:"id" gorm:"primary_key;" bson:"id"`
	Name string             `json:"name" bson:"name" gorm:"column:name;type:varchar(500);default:''"`
}

func (o OptionSkillEntity) TableName() string {
	return "optionskill"
}

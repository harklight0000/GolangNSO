package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type NPCEntity struct {
	_ID    primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID     int                `json:"id" gorm:"primary_key;" bson:"id"`
	Name   string             `json:"name" bson:"name" gorm:"column:name;type:varchar(500);default:''"`
	Head   int16              `json:"head" bson:"head" gorm:"column:head;type:int;default:0"`
	Body   int16              `json:"body" bson:"body" gorm:"column:body;type:int;default:0"`
	Leg    int16              `json:"leg" bson:"leg" gorm:"column:leg;type:int;default:0"`
	Type   int                `json:"type" bson:"type" gorm:"column:type;type:tinyint(4);default:0"`
	TaskID int                `json:"taskid" bson:"taskid" gorm:"column:taskid;type:tinyint(6);default:0"`
	Talk   string             `json:"talk" bson:"talk" gorm:"column:talk;type:varchar(5000);default:'[]'"`
}

func (N NPCEntity) TableName() string {
	return "npc"
}

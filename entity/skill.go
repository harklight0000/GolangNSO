package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nso/cache"
)

type SkillEntity struct {
	_ID            primitive.ObjectID     `json:"_id" bson:"_id" gorm:"-"`
	ID             byte                   `json:"id" gorm:"primary_key;" bson:"id"`
	Name           string                 `json:"name" bson:"name" gorm:"column:name;type:varchar(500);default:''"`
	MaxPoint       int                    `json:"maxPoint" bson:"maxPoint" gorm:"column:maxPoint;type:int;default:0"`
	Type           byte                   `json:"type" bson:"type" gorm:"column:type;type:tinyint;default:0"`
	IconID         int                    `json:"iconId" bson:"iconId" gorm:"column:iconId;type:int;default:0"`
	Desc           string                 `json:"desc" bson:"desc" gorm:"column:desc;type:varchar(5000);default:''"`
	SkillTemplates []cache.SkillTemplates `json:"skillTemplates" bson:"skillTemplates" gorm:"column:skillTemplates;type:longtext;default:'';serializer:my_json"`
}

func (s SkillEntity) TableName() string {
	return "skill"
}

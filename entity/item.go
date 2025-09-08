package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nso/cache"
)

type ItemEntity struct {
	_ID           primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID            int16              `json:"id" gorm:"primary_key;" bson:"id"`
	Type          int                `json:"type" bson:"type" gorm:"column:type;type:tinyint(6);default:0"`
	Class         int                `json:"class" bson:"class" gorm:"column:class;type:tinyint(6);default:0"`
	Skill         int                `json:"skill" bson:"skill" gorm:"column:skill;type:tinyint(6);default:0"`
	Name          string             `json:"name" bson:"name" gorm:"column:name;type:varchar(500);default:''"`
	Description   string             `json:"description" bson:"description" gorm:"column:description;type:varchar(5000);default:''"`
	Level         byte               `json:"level" bson:"level" gorm:"column:level;type:tinyint(6);default:0"`
	IconID        int16              `json:"iconID" bson:"iconID" gorm:"column:iconID;type:int;default:0"`
	Part          int16              `json:"part" bson:"part" gorm:"column:part;type:int;default:0"`
	UpToUp        byte               `json:"uptoup" bson:"uptoup" gorm:"column:uptoup;type:tinyint(6);default:0"`
	IsExpires     int                `json:"isExpires" bson:"isExpires" gorm:"column:isExpires;type:tinyint(4);default:0"`
	SecondExpires int64              `json:"secondExpires" bson:"secondExpires" gorm:"column:secondExpires;type:bigint;default:0"`
	SaleCoinLock  int                `json:"saleCoinLock" bson:"saleCoinLock" gorm:"column:saleCoinLock;type:tinyint(6);default:5"`
	ItemOption    []cache.Option     `json:"ItemOption" bson:"ItemOption" gorm:"column:ItemOption;type:varchar(5000);default:'[]';serializer:my_json"`
	Option1       []cache.Option     `json:"Option1" bson:"Option1" gorm:"column:Option1;type:varchar(500);default:'[]';serializer:my_json"`
	Option2       []cache.Option     `json:"Option2" bson:"Option2" gorm:"column:Option2;type:varchar(500);default:'[]';serializer:my_json"`
	Option3       []cache.Option     `json:"Option3" bson:"Option3" gorm:"column:Option3;type:varchar(500);default:'[]';serializer:my_json"`
	Gender        int                `json:"gender" bson:"gender" gorm:"column:gender;type:tinyint;default:0"`
}

func (i ItemEntity) TableName() string {
	return "item"
}

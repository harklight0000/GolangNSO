package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate go run ../generate/migrate.go
type MapEntity struct {
	_ID       primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID        int                `json:"id" gorm:"primary_key;" bson:"id"`
	TileID    int                `json:"tileID" bson:"tileID" gorm:"column:tileID;type:tinyint;default:0"`
	BgID      int                `json:"bgID" bson:"bgID" gorm:"column:bgID;type:tinyint;default:0"`
	TypeMap   int                `json:"typeMap" bson:"typeMap" gorm:"column:typeMap;type:tinyint;default:0"`
	Name      string             `json:"name" bson:"name" gorm:"column:name;type:varchar(500);default:''"`
	Vgo       [][]int            `json:"Vgo" bson:"Vgo" gorm:"column:Vgo;type:varchar(5000);default:'[]';serializer:my_json"`
	Mob       string             `json:"Mob" bson:"Mob" gorm:"column:Mob;type:varchar(5000);default:'[]'"`
	NPC       string             `json:"NPC" bson:"NPC" gorm:"column:NPC;type:varchar(5000);default:'[]'"`
	MaxPlayer int                `json:"maxplayer" bson:"maxplayer" gorm:"column:maxplayer;type:tinyint;default:20"`
	NumZone   int                `json:"numzone" bson:"numzone" gorm:"column:numzone;type:tinyint;default:30"`
	X0        int                `json:"x0" bson:"x0" gorm:"column:x0;type:int;default:0"`
	Y0        int                `json:"y0" bson:"y0" gorm:"column:y0;type:int;default:0"`
}

func (m MapEntity) TableName() string {
	return "map"
}

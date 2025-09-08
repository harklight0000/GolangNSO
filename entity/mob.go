package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type MobEntity struct {
	_ID             primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID              int                `json:"id" gorm:"primary_key;" bson:"id"`
	Name            string             `json:"name" bson:"name" gorm:"column:name;type:varchar(100);default:''"`
	Hp              int                `json:"hp" bson:"hp" gorm:"column:hp;type:int;default:0"`
	RangeMove       byte               `json:"rangeMove" bson:"rangeMove" gorm:"column:rangeMove;type:int;default:0"`
	Speed           byte               `json:"speed" bson:"speed" gorm:"column:speed;type:int;default:0"`
	TypeFly         int                `json:"typeFly" bson:"typeFly" gorm:"column:typeFly;type:int;default:0"`
	NImage          int                `json:"nImage" bson:"nImage" gorm:"column:nImage;type:int;default:0"`
	Flag            int                `json:"flag" bson:"flag" gorm:"column:flag;type:int;default:0"`
	FrameBossMove   string             `json:"frameBossMove" bson:"frameBossMove" gorm:"column:frameBossMove;type:varchar(500);default:'[]'"`
	FrameBossAttack string             `json:"frameBossAttack" bson:"frameBossAttack" gorm:"column:frameBossAttack;type:varchar(500);default:'[]'"`
	Info            int                `json:"info" bson:"info" gorm:"column:info;type:int;default:0"`
	Img1            string             `json:"Img1" bson:"Img1" gorm:"column:Img1;type:varchar(2500);default:'[]'"`
	Img2            string             `json:"Img2" bson:"Img2" gorm:"column:Img2;type:varchar(2500);default:'[]'"`
	Img3            string             `json:"Img3" bson:"Img3" gorm:"column:Img3;type:varchar(2500);default:'[]'"`
	Img4            string             `json:"Img4" bson:"Img4" gorm:"column:Img4;type:varchar(2500);default:'[]'"`
	FrameBoss       string             `json:"frameBoss" bson:"frameBoss" gorm:"column:frameBoss;type:varchar(8000);default:'[]'"`
	Type            byte               `json:"type" bson:"type" gorm:"column:type;type:tinyint(4);default:0"`
}

func (m MobEntity) TableName() string {
	return "mob"
}

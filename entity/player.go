package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlayerEntity struct {
	_ID            primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID             int                `json:"id" gorm:"primary_key;" bson:"id"`
	Username       string             `json:"username" bson:"username" gorm:"column:username;type:char(30);default:''"`
	Password       string             `json:"password" bson:"password" gorm:"column:password;type:char(30);default:''"`
	Luong          int                `json:"luong" bson:"luong" gorm:"column:luong;type:int;default:0"`
	NinjasArray    []string           `json:"ninja" bson:"ninja" gorm:"column:ninja;type:varchar(500);default:'[]';serializer:my_json"`
	Coin           int                `json:"coin" bson:"coin" gorm:"column:coin;type:int;default:0"`
	TopSK          int                `json:"topSK" bson:"topSK" gorm:"column:topSK;type:int;default:0"`
	Lock           int                `json:"lock" bson:"lock" gorm:"column:lock;type:int;default:0"`
	KichHoat       int                `json:"kichhoat" bson:"kichhoat" gorm:"column:kichhoat;type:int;default:1"`
	Nap            int                `json:"nap" bson:"nap" gorm:"column:nap;type:int;default:0"`
	ClanTeritoryID int                `json:"clanTeritoryID" bson:"clanTeritoryID" gorm:"column:clanTeritoryID;type:int;default:0"`
	TanThu         int                `json:"tanThu" bson:"tanThu" gorm:"column:tanThu;type:int;default:0"`
	DiemDanh       int                `json:"diemDanh" bson:"diemDanh" gorm:"column:diemDanh;type:int;default:0"`
}

func (p PlayerEntity) TableName() string {
	return "player"
}

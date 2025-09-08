package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type GiftCodeEntity struct {
	_ID           primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID            int                `json:"idgift" gorm:"column:idgift;primary_key;auto_increment" bson:"idgift"`
	GiftCode      string             `json:"giftcode" bson:"giftcode" gorm:"column:giftcode;type:varchar(999);default:''"`
	Yen           int                `json:"yen" bson:"yen" gorm:"column:yen;type:int;default:0"`
	Xu            int                `json:"xu" bson:"xu" gorm:"column:xu;type:int;default:0"`
	Long          int                `json:"long" bson:"long" gorm:"column:long;type:int;default:0"`
	ItemID        int                `json:"itemId" bson:"itemId" gorm:"column:itemId;type:int;default:-1"`
	ItemQuantity  int                `json:"itemQuantity" bson:"itemQuantity" gorm:"column:itemQuantity;type:int;default:0"`
	ItemID1       int                `json:"itemId1" bson:"itemId1" gorm:"column:itemId1;type:int;default:-1"`
	ItemQuantity1 int                `json:"itemQuantity1" bson:"itemQuantity1" gorm:"column:itemQuantity1;type:int;default:0"`
	ItemID2       int                `json:"itemId2" bson:"itemId2" gorm:"column:itemId2;type:int;default:-1"`
	ItemQuantity2 int                `json:"itemQuantity2" bson:"itemQuantity2" gorm:"column:itemQuantity2;type:int;default:0"`
	ItemID3       int                `json:"itemId3" bson:"itemId3" gorm:"column:itemId3;type:int;default:-1"`
	ItemQuantity3 int                `json:"itemQuantity3" bson:"itemQuantity3" gorm:"column:itemQuantity3;type:int;default:0"`
	ItemID4       int                `json:"itemId4" bson:"itemId4" gorm:"column:itemId4;type:int;default:-1"`
	ItemQuantity4 int                `json:"itemQuantity4" bson:"itemQuantity4" gorm:"column:itemQuantity4;type:int;default:0"`
	ItemID5       int                `json:"itemId5" bson:"itemId5" gorm:"column:itemId5;type:int;default:-1"`
	ItemQuantity5 int                `json:"itemQuantity5" bson:"itemQuantity5" gorm:"column:itemQuantity5;type:int;default:0"`
	Status        int                `json:"status" bson:"status" gorm:"column:status;type:int;default:0"`
	UserName      string             `json:"username" bson:"username" gorm:"column:username;type:text;default:''"`
	LuotNhap      int                `json:"luotnhap" bson:"luotnhap" gorm:"column:luotnhap;type:int;default:0"`
	GioiHan       int                `json:"gioihan" bson:"gioihan" gorm:"column:gioihan;type:int;default:0"`
	MessTB        string             `json:"mess_tb" bson:"mess_tb" gorm:"column:mess_tb;type:varchar(999);default:''"`
}

func (g GiftCodeEntity) TableName() string {
	return "giftcode"
}

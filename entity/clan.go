package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ClanEntity struct {
	_ID            primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID             int                `gorm:"primary_key;auto_increment" bson:"id"`
	Name           string             `json:"name" bson:"name;default:''"`
	Exp            int64              `json:"exp" bson:"exp" gorm:"column:exp;type:bigint;default:0"`
	Level          int                `json:"level" bson:"level" gorm:"column:level;type:int;default:1"`
	ItemLevel      int                `json:"item_level" bson:"item_level" gorm:"column:item_level;type:int;default:0"`
	Coin           int                `json:"coin" bson:"coin" gorm:"column:coin;type:int;default:1000000"`
	RegDate        string             `json:"reg_date" bson:"reg_date" gorm:"column:reg_date;type:varchar(100);default:'28/05/2003 16:05:22'"`
	Log            string             `json:"log" bson:"log" gorm:"column:log;type:longtext;default:''"`
	Alert          string             `json:"alert" bson:"alert" gorm:"column:alert;type:varchar(200);default:''"`
	UseCard        int                `json:"use_card" bson:"use_card" gorm:"column:use_card;type:tinyint;default:4"`
	OpenDun        int                `json:"open_dun" bson:"open_dun" gorm:"column:open_dun;type:tinyint;default:3"`
	Debt           int                `json:"debt" bson:"debt" gorm:"column:debt;type:tinyint;default:0"`
	Members        string             `json:"members" bson:"members" gorm:"column:members;type:longtext;default:''"`
	Items          string             `json:"items" bson:"items" gorm:"column:items;type:varchar(5000);default:'[]'"`
	Week           string             `json:"week" bson:"week" gorm:"column:week;type:varchar(100);default:'2003-05-28 22:22:1'"`
	ClanBattleData string             `json:"clan_battle_data" bson:"clan_battle_data" gorm:"column:clan_battle_data;type:varchar(2000);default:''"`
	ClanThanThu    string             `json:"clan_than_thu" bson:"clan_than_thu" gorm:"column:clan_than_thu;type:varchar(5000);default:'[]'"`
}

func (c *ClanEntity) TableName() string {
	return "clan"
}

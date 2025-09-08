package entity

import "nso/cache"

type NinjaEntity struct {
	BodyEntity
	Head          int16   `json:"head" bson:"head" gorm:"column:head;type:tinyint(4);default:-1"`
	Maxluggage    byte    `json:"maxluggage" bson:"maxluggage" gorm:"column:maxluggage;type:tinyint(4);default:0"`
	Site          []int16 `json:"site" bson:"site" gorm:"column:site;type:varchar(100);default:'[22,1678,264,22,0]';serializer:my_json"`
	Clan          string  `json:"clan" bson:"clan" gorm:"column:clan;type:text;default:'[]'"`
	DenBu         int     `json:"denbu" bson:"denbu" gorm:"column:denbu;type:tinyint(4);default:0"`
	NewLogin      string  `json:"newlogin" bson:"newlogin" gorm:"column:newlogin;type:text;default:''"`
	DDClan        int     `json:"ddClan" bson:"ddClan" gorm:"column:ddClan;type:tinyint(4);default:0"`
	CaveID        int     `json:"caveID" bson:"caveID" gorm:"column:caveID;type:tinyint(6);default:-1"`
	NCave         int     `json:"nCave" bson:"nCave" gorm:"column:nCave;type:tinyint(4);default:1"`
	PointCave     int16   `json:"pointCave" bson:"pointCave" gorm:"column:pointCave;type:tinyint(4);default:0"`
	UseCave       int     `json:"useCave" bson:"useCave" gorm:"column:useCave;type:tinyint(4);default:5"`
	BagCaveMax    int     `json:"bagCaveMax" bson:"bagCaveMax" gorm:"column:bagCaveMax;type:tinyint(4);default:0"`
	ItemIDCaveMax int     `json:"itemIDCaveMax" bson:"itemIDCaveMax" gorm:"column:itemIDCaveMax;type:tinyint(6);default:0"`
	ExpType       int     `json:"exptype" bson:"exptype" gorm:"column:exptype;type:tinyint(6);default:0"`
	NvhnCount     int     `json:"nvhncount" bson:"nvhncount" gorm:"column:nvhncount;type:tinyint(6);default:20"`
	TaThuCount    int     `json:"tathucount" bson:"tathucount" gorm:"column:tathucount;type:tinyint(6);default:2"`

	Friends      []cache.Friend `json:"friend" bson:"friend" gorm:"column:friend;type:text;default:'[]';serializer:my_json"`
	ChienTruong  BattleData     `json:"chientruong" bson:"chientruong" gorm:"column:chientruong;type:varchar(100);default:'';serializer:my_json"`
	Tasks        []*TaskOrder   `json:"tasks" bson:"tasks" gorm:"column:tasks;type:varchar(1500);default:'[]';serializer:my_json"`
	TaskIndex    int            `json:"taskIndex" bson:"taskIndex" gorm:"column:taskIndex;type:tinyint(6);default:0"`
	TaskCount    int            `json:"taskCount" bson:"taskCount" gorm:"column:taskCount;type:tinyint(6);default:0"`
	QuaCap10     int            `json:"quacap10" bson:"quacap10" gorm:"column:quacap10;type:int;default:0"`
	QuaCap20     int            `json:"quacap20" bson:"quacap20" gorm:"column:quacap20;type:int;default:0"`
	QuaCap30     int            `json:"quacap30" bson:"quacap30" gorm:"column:quacap30;type:int;default:0"`
	QuaCap40     int            `json:"quacap40" bson:"quacap40" gorm:"column:quacap40;type:int;default:0"`
	QuaCap50     int            `json:"quacap50" bson:"quacap50" gorm:"column:quacap50;type:int;default:0"`
	TopSK        int            `json:"topSK" bson:"topSK" gorm:"column:topSK;type:int;default:0"`
	TopSK1       int            `json:"topSK1" bson:"topSK1" gorm:"column:topSK1;type:int;default:0"`
	TopSK2       int            `json:"topSK2" bson:"topSK2" gorm:"column:topSK2;type:int;default:0"`
	VuiXuan      int            `json:"vuixuan" bson:"vuixuan" gorm:"column:vuixuan;type:int;default:0"`
	KichHoatKm   int            `json:"kichhoatkm" bson:"kichhoatkm" gorm:"column:kichhoatkm;type:int;default:0"`
	LvKm         int            `json:"lvkm" bson:"lvkm" gorm:"column:lvkm;type:int;default:0"`
	ExpKm        int64          `json:"expkm" bson:"expkm" gorm:"column:expkm;type:bigint;default:0"`
	TaskDanhVong []int          `json:"taskdanhvong" bson:"taskdanhvong" gorm:"column:taskdanhvong;type:varchar(999);default:'[-1,-1,-1,-1,-1,-1]';serializer:my_json"`
	CharInfo     []int          `json:"char_info" bson:"char_info" gorm:"column:char_info;type:varchar(999);default:'[0,0,0,0,0,0,0,0,0,0,0,0]';serializer:my_json"`
	HaiLoc       int            `json:"hailoc" bson:"hailoc" gorm:"column:hailoc;type:int;default:2"`
	TaThuLenh    int            `json:"Tathulenh" bson:"Tathulenh" gorm:"column:Tathulenh;type:int;default:5"`
	TayTn        int            `json:"taytn" bson:"taytn" gorm:"column:taytn;type:int;default:0"`
	TayKn        int            `json:"taykn" bson:"taykn" gorm:"column:taykn;type:int;default:0"`
	Gender       int8           `json:"gender" bson:"gender" gorm:"type:int;default:-1"`
	TaskID       int8           `json:"taskId" bson:"taskId" gorm:"column:taskId;type:tinyint(6);default:1"`
}

func (n NinjaEntity) TableName() string {
	return "ninja"
}

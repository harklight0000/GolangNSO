package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nso/cache"
)

type BodyEntity struct {
	_ID           primitive.ObjectID `json:"_id" gorm:"-" bson:"_id"`
	ID            int                `json:"id" gorm:"primary_key" bson:"id"`
	Name          string             `json:"name" bson:"name" gorm:"column:name;type:varchar(100);default:''"`
	Speed         byte               `json:"speed" bson:"speed" gorm:"column:speed;type:tinyint;default:4"`
	Class         byte               `json:"class" bson:"class" gorm:"column:class;type:tinyint;default:0"`
	KSkill        []int8             `json:"KSkill" bson:"KSkill" gorm:"column:KSkill;type:varchar(100);default:'[-1,-1,-1]';serializer:my_json"`
	OSkill        []int8             `json:"OSkill" bson:"OSkill" gorm:"column:OSkill;type:varchar(100);default:'[-1,-1,-1,-1,-1]';serializer:my_json"`
	CSkill        int                `json:"CSkill" bson:"CSkill" gorm:"column:CSkill;type:tinyint(4);default:0"`
	Level         int                `json:"level" bson:"level" gorm:"column:level;type:tinyint(6);default:1"`
	Exp           int64              `json:"exp" bson:"exp" gorm:"column:exp;type:bigint;default:0"`
	ExpDown       int64              `json:"expdown" bson:"expdown" gorm:"column:exp_down;type:bigint;default:0"`
	Pk            byte               `json:"pk" bson:"pk" gorm:"column:pk;type:tinyint;default:0"`
	Xu            int                `json:"xu" bson:"xu" gorm:"column:xu;type:int;default:0"`
	XuBox         int                `json:"xuBox" bson:"xuBox" gorm:"column:xuBox;type:int;default:0"`
	Yen           int                `json:"yen" bson:"yen" gorm:"column:yen;type:int;default:0"`
	MaxLuggage    int                `json:"maxluggage" bson:"maxluggage" gorm:"column:maxluggage;type:tinyint;default:30"`
	LevelBag      int                `json:"levelBag" bson:"levelBag" gorm:"column:levelBag;type:tinyint;default:0"`
	Item_Bag      []cache.ItemJSON   `json:"ItemBag" bson:"ItemBag" gorm:"column:ItemBag;type:text;default:'[]';serializer:my_json"`
	Item_Box      []cache.ItemJSON   `json:"ItemBox" bson:"ItemBox" gorm:"column:ItemBox;type:text;default:'[]';serializer:my_json"`
	Item_CaiTrang []cache.ItemJSON   `json:"ItemCaiTrang" bson:"ItemCaiTrang" gorm:"column:ItemCaiTrang;type:text;default:'[]';serializer:my_json"`
	Item_BST      []cache.ItemJSON   `json:"ItemBST" bson:"ItemBST" gorm:"column:ItemBST;type:varchar(5000);default:'[]';serializer:my_json"`
	Item_BodyHide []cache.ItemJSON   `json:"ItemBodyHide" bson:"ItemBodyHide" gorm:"column:ItemBodyHide;type:varchar(5000);default:'[]';serializer:my_json"`
	Item_Body     []cache.ItemJSON   `json:"ItemBody" bson:"ItemBody" gorm:"column:ItemBody;type:text;default:'[]';serializer:my_json"`
	Item_Mounts   []cache.ItemJSON   `json:"ItemMounts" bson:"ItemMounts" gorm:"column:ItemMounts;type:text;default:'[]';serializer:my_json"`
	Effect        string             `json:"effect" bson:"effect" gorm:"column:effect;type:longtext;default:'[]'"`
	Skill_s       string             `json:"skill" bson:"skill" gorm:"column:skill;type:varchar(5000);default:'[]'"`
	TiemNangSo    int                `json:"tiemnangso" bson:"tiemnangso" gorm:"column:tiemnangso;type:tinyint;default:0"`
	KyNangSo      int                `json:"kynangso" bson:"kynangso" gorm:"column:kynangso;type:tinyint;default:0"`
	PhongLoi      int                `json:"phongloi" bson:"phongloi" gorm:"column:phongloi;type:tinyint;default:0"`
	BangHoa       int                `json:"banghoa" bson:"banghoa" gorm:"column:banghoa;type:tinyint;default:0"`
	PPoint        int                `json:"ppoint" bson:"ppoint" gorm:"column:ppoint;type:int;default:0"`
	Potential0    int16              `json:"potential0" bson:"potential0" gorm:"column:potential0;type:int;default:15"`
	Potential1    int16              `json:"potential1" bson:"potential1" gorm:"column:potential1;type:int;default:5"`
	Potential2    int                `json:"potential2" bson:"potential2" gorm:"column:potential2;type:int;default:5"`
	Potential3    int                `json:"potential3" bson:"potential3" gorm:"column:potential3;type:int;default:5"`
	SPoint        int16              `json:"spoint" bson:"spoint" gorm:"column:spoint;type:int;default:0"`
}

func (c BodyEntity) TableName() string {
	return "clone_ninja"
}

package ainterfaces

import (
	"nso/cache"
	"nso/entity"
)

type IItem interface {
	Clone() IItem
	IsExpired() bool
	IsExpiredEggDaemon() bool
	GetUpgradeMax() int
	UpgradeNext(next byte)
	IsTypeBody() bool
	IsTypeNgocKham() bool
	GetData() *entity.ItemEntity
	IsPrecious() bool
	PercentAppear() int
	FindParamById(id int) int
	IsLock() bool
	SetLock(lock bool)
	GetUpgrade() byte
	SetUpgrade(upgrade byte)
	IsTrangSuc() bool
	IsYoroi() bool
	IsTrangPhuc() bool
	IsTypeTask() bool
	IsVuKhi() bool
	GetIDJiraiNam(_type int) int
	GetIDJiraiNu(_type int) int
	IsTypeMount() bool
}

type IITemShop interface {
	IItem
	GetOptionShopMin(optionID int, param int) int
}

type IITemFactory interface {
	ItemDefault1(id int) IItem
	ItemDefaultLock(id int, isLock bool) IItem
	ItemDefaultSys(id int, sys byte) IItem
	FromJSON(json cache.ItemJSON) IItem
	FromString(string string) IItem
	ToJSON(item IItem, index int) cache.ItemJSON
	GetItemIDByLevel(maxLevel int, _type byte, gender byte) []int16
	IsUpgradeHide(id int, upgrade byte) bool
	ItemNgocDefault(id int, upgrade int, random bool) IItem
	ItemDefaultMat(id int) IItem
}

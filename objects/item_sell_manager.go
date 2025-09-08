package objects

import (
	"encoding/json"
	"go.uber.org/zap"
	. "nso/ainterfaces"
	"nso/cache"
	"nso/entity"
	"nso/logging"
)

type ItemSell struct {
	*entity.ItemSellEntity
	items []*Item
}

func NewItemSell(itemSellEntity *entity.ItemSellEntity, factory IITemFactory) *ItemSell {
	this := &ItemSell{ItemSellEntity: itemSellEntity}
	var jsons []cache.ItemJSON
	err := json.Unmarshal([]byte(itemSellEntity.ListItem), &jsons)
	if err != nil {
		logging.Logger.Error("Unmarshal item sell error", zap.Error(err))
	}
	this.items = make([]*Item, len(jsons))
	for i, itJson := range jsons {
		this.items[i] = factory.FromJSON(itJson).(*Item)
	}
	return this
}

func (this ItemSell) GetID() int {
	return this.ID
}

func (this ItemSell) GetType() int {
	return this.Type
}

func (this ItemSell) GetItems() []IItem {
	var items []IItem
	for _, item := range this.items {
		items = append(items, item)
	}
	return items
}

func (this ItemSell) GetItemByIndex(index int) IItem {
	if index < 0 || index >= len(this.items) {
		return nil
	}
	return this.items[index]
}

type ItemSellListManager struct {
	AppCtx    IAppContext
	itemSells map[int]IItemSell
}

func (this ItemSellListManager) GetItemSellByType(_type int) IItemSell {
	return this.itemSells[_type]
}

func NewItemSellListManager(ctx IAppContext) IItemSellListManager {
	this := &ItemSellListManager{}
	this.AppCtx = ctx
	db := this.AppCtx.GetDatabase()
	var itemsSells []*entity.ItemSellEntity
	err := db.FindAll(&itemsSells)
	if err != nil {
		logging.Logger.Error("Can not load item sells", zap.Error(err))
	}
	this.itemSells = make(map[int]IItemSell)
	for _, itemSell := range itemsSells {
		this.itemSells[itemSell.Type] = NewItemSell(itemSell, ctx.GetGameData().GetItemFactory())
	}
	return this
}

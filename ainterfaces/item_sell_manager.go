package ainterfaces

type IItemSell interface {
	GetID() int
	GetType() int
	GetItems() []IItem
	GetItemByIndex(index int) IItem
}

type IItemSellListManager interface {
	GetItemSellByType(_type int) IItemSell
}

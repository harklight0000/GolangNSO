package ainterfaces

type IUser interface {
	ISender
	IParseData
	IUpdate
	GetID() int
	Get() INinja
	GetHuman() INinja
	GetClone() INinja
	Leave()
	RequestItem(itemType int)
	GetNinja() INinja
}

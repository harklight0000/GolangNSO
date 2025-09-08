package ainterfaces

type IMenu interface {
	Process(user IUser, options MenuOptions)
	GetMenu(user IUser) []string
	DefaultMenu(user IUser) []string
	SendWrite(user IUser, title string)
}

type IMenuFactory interface {
	GetMenuProcessor(npcId int16) IMenu
}

type MenuOptions struct {
	Option1 int8
	Option2 int8
	Option3 int8
}

func NewMenuOptions(option1 int8, option2 int8, option3 int8) *MenuOptions {
	return &MenuOptions{Option1: option1, Option2: option2, Option3: option3}
}

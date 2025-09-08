package menu

import (
	. "nso/ainterfaces"
)

type GooshoMenu struct {
	MenuInfo
}

func (g GooshoMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (g GooshoMenu) GetMenu(user IUser) []string {
	return nil
}

func (g GooshoMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

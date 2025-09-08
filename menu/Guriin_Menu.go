package menu

import (
	. "nso/ainterfaces"
)

type GuriinMenu struct {
	MenuInfo
}

func (g GuriinMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (g GuriinMenu) GetMenu(user IUser) []string {
	return nil
}

func (g GuriinMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

package menu

import (
	. "nso/ainterfaces"
)

type OnggiaNoelMenu struct {
	MenuInfo
}

func (o *OnggiaNoelMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (o *OnggiaNoelMenu) GetMenu(user IUser) []string {
	return nil
}

func (o *OnggiaNoelMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

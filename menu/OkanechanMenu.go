package menu

import (
	. "nso/ainterfaces"
)

type OkanechanMenu struct {
	MenuInfo
}

func (o *OkanechanMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (o *OkanechanMenu) GetMenu(user IUser) []string {
	return nil
}

func (o *OkanechanMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

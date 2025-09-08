package menu

import (
	. "nso/ainterfaces"
)

type FujiwaraMenu struct {
	MenuInfo
}

func (f *FujiwaraMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (f *FujiwaraMenu) GetMenu(user IUser) []string {
	return nil
}

func (f *FujiwaraMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

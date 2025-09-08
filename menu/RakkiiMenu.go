package menu

import (
	. "nso/ainterfaces"
)

type RakkiiMenu struct {
	MenuInfo
}

func (r *RakkiiMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (r *RakkiiMenu) GetMenu(user IUser) []string {
	return nil
}

func (r *RakkiiMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

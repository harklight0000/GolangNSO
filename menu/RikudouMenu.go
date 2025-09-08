package menu

import (
	. "nso/ainterfaces"
)

type RikudouMenu struct {
	MenuInfo
}

func (r *RikudouMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (r *RikudouMenu) GetMenu(user IUser) []string {
	return nil
}

func (r *RikudouMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

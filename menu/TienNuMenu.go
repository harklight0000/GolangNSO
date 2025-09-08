package menu

import (
	. "nso/ainterfaces"
)

type TienNuMenu struct {
	MenuInfo
}

func (t TienNuMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (t TienNuMenu) GetMenu(user IUser) []string {
	return []string{"Làm diều", "Chung sức xây dựng", "Top"}
}

func (t TienNuMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

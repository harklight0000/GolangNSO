package menu

import (
	. "nso/ainterfaces"
)

type VuaHungMenu struct {
	MenuInfo
}

func (v *VuaHungMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (v *VuaHungMenu) GetMenu(user IUser) []string {
	return nil
}

func (v *VuaHungMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

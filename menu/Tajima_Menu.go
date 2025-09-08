package menu

import (
	. "nso/ainterfaces"
)

type TajimaMenu struct {
	MenuInfo
}

func (t *TajimaMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (t *TajimaMenu) GetMenu(user IUser) []string {
	return nil
}

func (t *TajimaMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

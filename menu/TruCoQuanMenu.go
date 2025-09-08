package menu

import (
	. "nso/ainterfaces"
)

type TruCoQuanMenu struct {
	MenuInfo
}

func (t *TruCoQuanMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (t *TruCoQuanMenu) GetMenu(user IUser) []string {
	return nil
}

func (t *TruCoQuanMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

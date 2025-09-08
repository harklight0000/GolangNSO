package menu

import (
	. "nso/ainterfaces"
)

type ChiHangMenu struct {
	MenuInfo
}

func (c *ChiHangMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (c *ChiHangMenu) GetMenu(user IUser) []string {
	return nil
}

func (c *ChiHangMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

package menu

import (
	. "nso/ainterfaces"
)

type CasinoMenu struct {
	MenuInfo
}

func (c *CasinoMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (c *CasinoMenu) GetMenu(user IUser) []string {
	return nil
}

func (c *CasinoMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

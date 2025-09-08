package menu

import (
	. "nso/ainterfaces"
)

type CaythongMenu struct {
	MenuInfo
}

func (c *CaythongMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (c *CaythongMenu) GetMenu(user IUser) []string {
	return nil
}

func (c *CaythongMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

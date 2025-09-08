package menu

import (
	. "nso/ainterfaces"
)

type CaymaiMenu struct {
	MenuInfo
}

func (c *CaymaiMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (c *CaymaiMenu) GetMenu(user IUser) []string {
	return nil
}

func (c *CaymaiMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

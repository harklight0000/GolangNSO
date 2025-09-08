package menu

import (
	. "nso/ainterfaces"
)

type CayDaoMenu struct {
	MenuInfo
}

func (c *CayDaoMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (c *CayDaoMenu) GetMenu(user IUser) []string {
	return nil
}

func (c *CayDaoMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

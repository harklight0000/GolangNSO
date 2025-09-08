package menu

import (
	. "nso/ainterfaces"
)

type NaoMenu struct {
	MenuInfo
}

func (n *NaoMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (n *NaoMenu) GetMenu(user IUser) []string {
	return nil
}

func (n *NaoMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

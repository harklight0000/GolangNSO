package menu

import (
	. "nso/ainterfaces"
)

type KagaiMenu struct {
	MenuInfo
}

func (k *KagaiMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (k *KagaiMenu) GetMenu(user IUser) []string {
	return nil
}

func (k *KagaiMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

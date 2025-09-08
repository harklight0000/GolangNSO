package menu

import (
	. "nso/ainterfaces"
)

type KirinMenu struct {
	MenuInfo
}

func (k *KirinMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (k *KirinMenu) GetMenu(user IUser) []string {
	return nil
}

func (k *KirinMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

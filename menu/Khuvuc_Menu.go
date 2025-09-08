package menu

import (
	. "nso/ainterfaces"
)

type KhuVucMenu struct {
	MenuInfo
}

func (k *KhuVucMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (k *KhuVucMenu) GetMenu(user IUser) []string {
	return nil
}

func (k *KhuVucMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

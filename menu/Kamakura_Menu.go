package menu

import (
	. "nso/ainterfaces"
)

type KamakuraMenu struct {
	MenuInfo
}

func (this *KamakuraMenu) Process(user IUser, options MenuOptions) {
}

func (this *KamakuraMenu) GetMenu(user IUser) []string {
	return nil
}

func (this *KamakuraMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

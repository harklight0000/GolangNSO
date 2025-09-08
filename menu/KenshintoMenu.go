package menu

import (
	. "nso/ainterfaces"
)

type KenshintoMenu struct {
	MenuInfo
}

func (this *KenshintoMenu) Process(user IUser, options MenuOptions) {
}

func (this *KenshintoMenu) GetMenu(user IUser) []string {
	return nil
}

func (this *KenshintoMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

package menu

import (
	. "nso/ainterfaces"
)

type TabemonoMenu struct {
	MenuInfo
}

func (this *TabemonoMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (this *TabemonoMenu) GetMenu(user IUser) []string {
	return nil
}

func (this *TabemonoMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

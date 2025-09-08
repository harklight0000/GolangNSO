package menu

import (
	. "nso/ainterfaces"
)

type BaReiMenu struct {
	MenuInfo
}

func (b *BaReiMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (b *BaReiMenu) GetMenu(user IUser) []string {
	return nil
}

func (b *BaReiMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

package menu

import (
	. "nso/ainterfaces"
)

type LongDenMenu struct {
	MenuInfo
}

func (l *LongDenMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (l *LongDenMenu) GetMenu(user IUser) []string {
	return nil
}

func (l *LongDenMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

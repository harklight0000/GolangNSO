package menu

import (
	. "nso/ainterfaces"
)

type DaiThanhMenu struct {
	MenuInfo
}

func (d *DaiThanhMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (d *DaiThanhMenu) GetMenu(user IUser) []string {
	return nil
}

func (d *DaiThanhMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

package menu

import (
	. "nso/ainterfaces"
)

type MatsurugiMenu struct {
	MenuInfo
}

func (m *MatsurugiMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (m *MatsurugiMenu) GetMenu(user IUser) []string {
	return nil
}

func (m *MatsurugiMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

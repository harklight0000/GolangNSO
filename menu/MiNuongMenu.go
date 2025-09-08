package menu

import (
	. "nso/ainterfaces"
)

type MiNuongMenu struct {
	MenuInfo
}

func (m *MiNuongMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (m *MiNuongMenu) GetMenu(user IUser) []string {
	return nil
}

func (m *MiNuongMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

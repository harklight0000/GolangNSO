package menu

import (
	. "nso/ainterfaces"
)

type TashinoMenu struct {
	MenuInfo
}

func (t *TashinoMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (t *TashinoMenu) GetMenu(user IUser) []string {
	return nil
}

func (t *TashinoMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

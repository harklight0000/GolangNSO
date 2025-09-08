package menu

import (
	. "nso/ainterfaces"
)

type AdminMenu struct {
	MenuInfo
}

func (a *AdminMenu) Process(user IUser, options MenuOptions) {
	menuId := options.Option1
	switch menuId {
	case 0:
		user.RequestItem(2)
	case 1:

	}
}

func (a *AdminMenu) GetMenu(user IUser) []string {
	return nil
}

func (a *AdminMenu) SendWrite(user IUser, title string) {

}

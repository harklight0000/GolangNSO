package menu

import (
	. "nso/ainterfaces"
)

type ThayOokamesamaMenu struct {
	MenuInfo
}

func (t *ThayOokamesamaMenu) Process(user IUser, options MenuOptions) {
}

func (t *ThayOokamesamaMenu) GetMenu(user IUser) []string {
	return nil
}

func (t *ThayOokamesamaMenu) SendWrite(user IUser, title string) {
}

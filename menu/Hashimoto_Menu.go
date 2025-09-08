package menu

import (
	. "nso/ainterfaces"
)

type HashimotoMenu struct {
	MenuInfo
}

func (h *HashimotoMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (h *HashimotoMenu) GetMenu(user IUser) []string {
	return nil
}

func (h *HashimotoMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

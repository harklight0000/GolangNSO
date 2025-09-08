package menu

import (
	. "nso/ainterfaces"
)

type KirikoMenu struct {
	MenuInfo
}

func (k *KirikoMenu) Process(user IUser, options MenuOptions) {

}

func (k *KirikoMenu) GetMenu(user IUser) []string {
	return nil
}

func (k *KirikoMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

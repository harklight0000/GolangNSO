package menu

import (
	. "nso/ainterfaces"
)

type AmejiMenu struct {
	MenuInfo
}

func (this *AmejiMenu) Process(user IUser, options MenuOptions) {

}

func (this *AmejiMenu) GetMenu(user IUser) []string {
	return nil
}

func (this *AmejiMenu) SendWrite(user IUser, title string) {
}

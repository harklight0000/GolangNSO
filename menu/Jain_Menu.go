package menu

import (
	. "nso/ainterfaces"
)

type JaianMenu struct {
	MenuInfo
}

func (j *JaianMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (j *JaianMenu) GetMenu(user IUser) []string {
	return nil
}

func (j *JaianMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

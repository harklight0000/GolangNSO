package menu

import (
	. "nso/ainterfaces"
)

type SunooMenu struct {
	MenuInfo
}

func (s *SunooMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (s *SunooMenu) GetMenu(user IUser) []string {
	return nil
}

func (s *SunooMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

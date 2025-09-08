package menu

import (
	. "nso/ainterfaces"
)

type SobaMenu struct {
	MenuInfo
}

func (s *SobaMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (s *SobaMenu) GetMenu(user IUser) []string {
	return nil
}

func (s *SobaMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

package menu

import (
	. "nso/ainterfaces"
)

type ShinwaMenu struct {
	MenuInfo
}

func (s *ShinwaMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (s *ShinwaMenu) GetMenu(user IUser) []string {
	return nil
}

func (s *ShinwaMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

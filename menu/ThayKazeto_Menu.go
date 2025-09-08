package menu

import (
	. "nso/ainterfaces"
)

type ThayKazetoMenu struct {
	MenuInfo
}

func (t *ThayKazetoMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (t *ThayKazetoMenu) GetMenu(user IUser) []string {
	return nil
}

func (t *ThayKazetoMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

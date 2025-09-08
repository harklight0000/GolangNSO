package menu

import . "nso/ainterfaces"

type KatanaMenu struct {
	MenuInfo
}

func (this *KatanaMenu) Process(user IUser, options MenuOptions) {
	//TODO implement me
	panic("implement me")
}

func (this *KatanaMenu) GetMenu(user IUser) []string {
	return nil
}

func (this *KatanaMenu) SendWrite(user IUser, title string) {
	//TODO implement me
	panic("implement me")
}

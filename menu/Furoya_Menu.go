package menu

import (
	. "nso/ainterfaces"
	"nso/logging"
)

type FuroyaMenu struct {
	MenuInfo
}

func (f *FuroyaMenu) Process(user IUser, options MenuOptions) {
	logging.Logger.Debug("FuroyaMenu.Process()")
}

func (f *FuroyaMenu) GetMenu(user IUser) []string {
	return nil
}

func (f *FuroyaMenu) SendWrite(user IUser, title string) {
}

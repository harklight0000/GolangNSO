package menu

import (
	. "nso/ainterfaces"
)

type CoToyotomiMenu struct {
	MenuInfo
}

func (c *CoToyotomiMenu) Process(user IUser, options MenuOptions) {

}

func (c *CoToyotomiMenu) GetMenu(user IUser) []string {
	return nil
}

func (c *CoToyotomiMenu) SendWrite(user IUser, title string) {
}

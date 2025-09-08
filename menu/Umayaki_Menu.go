package menu

import (
	. "nso/ainterfaces"
)

type UmayakiMenuLang struct {
	MenuInfo
}

func (this *UmayakiMenuLang) Process(user IUser, options MenuOptions) {
	switch options.Option1 {
	case 0:
		this.Chat(user, "Không chê ta nghèo lên ngựa ta đèo")
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		fallthrough
	case 7:
		ma := this.App.GetMapManager().GetMapByID(arrLang[options.Option1-1])
		area := ma.GetFreeArea()
		if area == nil {
			this.Chat(user, "Tất cả các khu đều đầy")
			return
		}
		area.Enter0(user)
	default:
		this.Chat(user, "Chức năng chưa có")
	}
}

func (this *UmayakiMenuLang) GetMenu(user IUser) []string {
	return nil
}

func (this *UmayakiMenuLang) SendWrite(user IUser, title string) {

}

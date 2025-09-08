package menu

import (
	"nso/ainterfaces"
	"nso/objects"
)

type UmayakiMenuTruong struct {
	ainterfaces.MenuInfo
}

var arrLang = []int{10, 17, 22, 32, 38, 43, 48}

var arrTruong = []int{1, 27, 72}

func (this *UmayakiMenuTruong) Process(user ainterfaces.IUser, options ainterfaces.MenuOptions) {
	switch options.Option1 {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		ma := this.App.GetMapManager().GetMapByID(arrTruong[options.Option1])
		area := ma.GetFreeArea().(*objects.Area)
		if area == nil {
			this.Chat(user, "Tất cả các khu đều đầy")
			return
		}
		area.Enter0(user)
	case 3:
		this.Chat(user, "Ta kéo xe qua làng và trường")
	default:
		this.Chat(user, "Chức năng chưa có")
	}
}

func (this *UmayakiMenuTruong) GetMenu(user ainterfaces.IUser) []string {
	return nil
}

func (this *UmayakiMenuTruong) SendWrite(user ainterfaces.IUser, title string) {

}

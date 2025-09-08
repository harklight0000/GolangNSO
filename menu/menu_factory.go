package menu

import (
	"go.uber.org/zap"
	. "nso/ainterfaces"
	. "nso/constants"
	"nso/logging"
)

type ProcessorFactory struct {
	npcs   []NpcInfo
	appCtx IAppContext
}

func NewProcessorFactory(appCtx IAppContext) IMenuFactory {
	this := &ProcessorFactory{}
	this.appCtx = appCtx
	this.npcs = appCtx.GetGameData().Npcs()
	return this
}

func (m *ProcessorFactory) GetMenuProcessor(npcId int16) IMenu {
	var menu IMenu
	var menuInfo MenuInfo
	menuInfo.NpcInfo = m.npcs[npcId]
	menuInfo.App = m.appCtx
	switch npcId {
	case Kanata:
		menu = &KatanaMenu{menuInfo}
	case Furoya:
		menu = &FuroyaMenu{menuInfo}
	case Ameji:
		menu = &AmejiMenu{menuInfo}
	case Kiriko:
		menu = &KirikoMenu{menuInfo}
	case Tabemono:
		menu = &TabemonoMenu{menuInfo}
	case Kamakura:
		menu = &KamakuraMenu{menuInfo}
	case Kenshinto:
		menu = &KenshintoMenu{menuInfo}
	case Umayaki_Lang:
		menu = &UmayakiMenuLang{menuInfo}
	case Umayaki_Truong:
		menu = &UmayakiMenuTruong{menuInfo}
	case CoToyotomi:
		menu = &CoToyotomiMenu{menuInfo}
	case ThayOokamesama:
		menu = &ThayOokamesamaMenu{menuInfo}
	case ThayKazeto:
		menu = &ThayKazetoMenu{menuInfo}
	case Tajima:
		menu = &TajimaMenu{menuInfo}
	case KhuVuc:
		menu = &KhuVucMenu{menuInfo}
	case Hashimoto:
		menu = &HashimotoMenu{menuInfo}
	case Fujiwara:
		menu = &FujiwaraMenu{menuInfo}
	case Nao:
		menu = &NaoMenu{menuInfo}
	case Jaian:
		menu = &JaianMenu{menuInfo}
	case BaRei:
		menu = &BaReiMenu{menuInfo}
	case Kirin:
		menu = &KirinMenu{menuInfo}
	case Soba:
		menu = &SobaMenu{menuInfo}
	case Sunoo:
		menu = &SunooMenu{menuInfo}
	case Guriin:
		menu = &GuriinMenu{menuInfo}
	case Matsurugi:
		menu = &MatsurugiMenu{menuInfo}
	case Okanechan:
		menu = &OkanechanMenu{menuInfo}
	case Rikudou:
		menu = &RikudouMenu{menuInfo}
	case Goosho:
		menu = &GooshoMenu{menuInfo}
	case TruCoQuan:
		menu = &TruCoQuanMenu{menuInfo}
	case Shinwa:
		menu = &ShinwaMenu{menuInfo}
	case ChiHang:
		menu = &ChiHangMenu{menuInfo}
	case Rakkii:
		menu = &RakkiiMenu{menuInfo}
	case LongDen:
		menu = &LongDenMenu{menuInfo}
	case Kagai:
		menu = &KagaiMenu{menuInfo}
	case TienNu:
		menu = &TienNuMenu{menuInfo}
	case Caythong:
		menu = &CaythongMenu{menuInfo}
	case OnggiaNoel:
		menu = &OnggiaNoelMenu{menuInfo}
	case VuaHung:
		menu = &VuaHungMenu{menuInfo}
	case Admin:
		menu = &AdminMenu{menuInfo}
	case Caymai:
		menu = &CaymaiMenu{menuInfo}
	case CayDao:
		menu = &CayDaoMenu{menuInfo}
	case Tashino:
		menu = &TashinoMenu{menuInfo}
	case Casino:
		menu = &CasinoMenu{menuInfo}
	case DaiThanh:
		menu = &DaiThanhMenu{menuInfo}
	case MiNuong:
		menu = &MiNuongMenu{menuInfo}
	default:
		logging.Logger.Info("Unknown npcId: ", zap.Int("npcID", int(npcId)))
	}
	return menu
}

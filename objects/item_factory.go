package objects

import (
	"encoding/json"
	"github.com/rotisserie/eris"
	. "nso/ainterfaces"
	"nso/cache"
	. "nso/constants"
	"nso/entity"
	. "nso/utils"
)

type ItemFactory struct {
	data []*entity.ItemEntity
}

func NewItemFactory(data []*entity.ItemEntity) *ItemFactory {
	this := &ItemFactory{}
	this.data = data
	return this
}

func (this *ItemFactory) ItemDefault1(id int) IItem {
	var item *Item
	if id >= 652 && id <= 655 {
		item = this.ItemNgocDefault(id, 1, true).(*Item)
	} else if id >= 685 && id <= 695 {
		item = this.ItemDefaultMat(id).(*Item)
	} else if id >= HAI_MA_1_ID && id <= DI_LONG_3_ID {
		item = this.ItemDefaultSys(id, byte(0)).(*Item)
		if id == HAI_MA_1_ID || id == DI_LONG_1_ID {
			item.Option = append(item.Option, cache.NewOption(ST_NGUOI_ID, 1000))
			item.Option = append(item.Option, cache.NewOption(ST_QUAI_ID, 5000))
		} else if id == HAI_MA_2_ID || id == DI_LONG_2_ID {
			item.Option = append(item.Option, cache.NewOption(ST_NGUOI_ID, 3000))
			item.Option = append(item.Option, cache.NewOption(ST_QUAI_ID, 15000))
		} else if id == HAI_MA_3_ID || id == DI_LONG_3_ID {
			item.Option = append(item.Option, cache.NewOption(ST_NGUOI_ID, 8000))
			item.Option = append(item.Option, cache.NewOption(ST_QUAI_ID, 30000))
		}
		return item
	} else if id == HOA_LONG_ID {
		item = this.ItemDefaultSys(id, 0).(*Item)
		item.Option = append(item.Option, cache.NewOption(ST_NGUOI_ID, 8000))
		item.Option = append(item.Option, cache.NewOption(ST_QUAI_ID, 30000))
	} else if id == CAN_CAU_CA {
		item = this.ItemDefaultLock(id, true).(*Item)
		item.Expires = TimeDay(3)
		return item
	} else if id == THIEN_BIEN_LENH {
		item = this.ItemDefaultLock(id, true).(*Item)
		item.Expires = TimeHour(5)
		return item
	} else {
		item = this.ItemDefaultSys(id, 0).(*Item)
	}
	return item
}

func (this *ItemFactory) ItemDefaultLock(id int, isLock bool) IItem {
	item := this.ItemDefaultSys(id, 0)
	item.SetLock(isLock)
	return item
}

func (this *ItemFactory) ItemDefaultSys(id int, sys byte) IItem {
	if id == -1 {
		return nil
	}
	data := *this.data[id]
	item := NewItem(data)
	item.Sys = sys
	if data.IsExpires == 1 {
		item.isExpired = true
		item.Expires = TimeSeconds(data.SecondExpires)
	}
	var options []cache.Option
	switch sys {
	case 0:
		options = data.ItemOption
	case 1:
		options = data.Option1
	case 2:
		options = data.Option2
	case 3:
		options = data.Option3
	}
	for _, option := range options {
		item.Option = append(item.Option, cache.NewOption(option.ID, option.Param))
	}

	return item
}

func (this *ItemFactory) FromJSON(json cache.ItemJSON) IItem {
	item := NewItem(*this.data[json.ID])
	item.SetLock(json.IsLock)
	item.Sale = json.Sale
	item.Quantity = json.Quantity
	item.Upgrade = json.Upgrade
	item.Index = json.Index
	item.ID = json.ID
	item.isExpired = json.IsExpires
	item.Expires = json.Expires
	item.BuyGold = json.BuyGold
	item.Sys = json.Sys
	item.TimeBuy = json.TimeBuy
	for _, ngoc := range json.Ngocs {
		item.Ngocs = append(item.Ngocs, *this.FromJSON(ngoc).(*Item))
	}
	item.Option = make([]cache.Option, len(json.Option))
	for i, option := range json.Option {
		item.Option[i] = cache.NewOption(option.ID, option.Param)
	}
	return item
}

func (this *ItemFactory) FromString(string string) IItem {
	var item cache.ItemJSON
	err := json.Unmarshal(([]byte)(string), &item)
	if err != nil {
		return nil
	}
	return this.FromJSON(item)
}

func (this *ItemFactory) ToJSON(item IItem, index int) cache.ItemJSON {
	it := item.(*Item)
	it.Index = index
	var json = cache.ItemJSON{}
	json.IsLock = it.IsLock()
	json.Sale = it.Sale
	json.Quantity = it.Quantity
	json.Upgrade = it.Upgrade
	json.Index = it.Index
	json.ID = it.ID
	json.IsExpires = it.IsExpired()
	json.Expires = it.Expires
	json.BuyGold = it.BuyGold
	json.Sys = it.Sys
	json.TimeBuy = it.TimeBuy
	var ngocs = make([]cache.ItemJSON, len(it.Ngocs))
	for i, ngoc := range it.Ngocs {
		ngocs = append(ngocs, this.ToJSON(&ngoc, i))
	}
	json.Ngocs = ngocs
	json.Option = it.Option
	return json
}

func (this *ItemFactory) GetItemIDByLevel(maxLevel int, _type byte, gender byte) []int16 {
	//TODO implement me
	panic("implement me")
}

func (this *ItemFactory) IsUpgradeHide(id int, upgrade byte) bool {
	return ((id == 27 || id == 30 || id == 60) && upgrade < 4) || ((id == 28 || id == 31 || id == 37 || id == 61) && upgrade < 8) || ((id == 29 || id == 32 || id == 38 || id == 62) && upgrade < 12) || ((id == 33 || id == 34 || id == 35 || id == 36 || id == 39) && upgrade < 14) || (((id >= 40 && id <= 46) || (id >= 48 && id <= 56)) && upgrade < 16)
}

func (this *ItemFactory) ItemNgocDefault(id int, upgrade int, ran bool) IItem {
	it := this.ItemDefaultSys(id, 0).(*Item)
	it.Upgrade = byte(upgrade)
	vuKhi := cache.NewOption(VU_KHI_OPTION_ID, 0)
	trangBi := cache.NewOption(TRANG_BI_OPTION_ID, 0)
	trangSuc := cache.NewOption(TRANG_SUC_OPTION_ID, 0)
	if id == HUYET_NGOC {
		// Vu khi
		it.Option = append(it.Option, vuKhi)
		it.Option = append(it.Option, cache.NewOption(OPTION_TAN_CONG_ID, random(MAX_TAN_CONG, 0.8, ran)))
		it.Option = append(it.Option, cache.NewOption(OPTION_CHI_MANG_ID, -random(MAX_CHI_MANG, 0.8, ran)))

		// Trang bi
		it.Option = append(it.Option, trangBi)
		it.Option = append(it.Option, cache.NewOption(OPTION_GIAM_TRU_ST_ID, random(MAX_GIAM_TRU_ST, 0.8, ran)))
		it.Option = append(it.Option, cache.NewOption(73, -random(NE_DON, 0.8, ran)))

		// Trang suc
		it.Option = append(it.Option, trangSuc)
		it.Option = append(it.Option, cache.NewOption(115, random(NE_DON, 0.8, ran)))
		it.Option = append(it.Option, cache.NewOption(119, random(MOI_GIAY_HOI_PHUC_MP, 0.8, ran)))
	} else if id == HUYEN_TINH_NGOC {
		// Vu khi
		it.Option = append(it.Option, vuKhi)
		it.Option = append(it.Option, cache.NewOption(102, random(ST_CHI_MANG, 0.8, ran)))
		it.Option = append(it.Option, cache.NewOption(115, -random(NE_DON, 0.8, ran)))

		// Trang bi
		it.Option = append(it.Option, trangBi)
		it.Option = append(it.Option, cache.NewOption(126, random(PHAN_DON, 0.8, ran)))
		it.Option = append(it.Option, cache.NewOption(105, -random(NE_DON, 0.6, ran)))

		// Trang suc
		it.Option = append(it.Option, trangSuc)
		it.Option = append(it.Option, cache.NewOption(114, random(MAX_CHI_MANG, 0.8, ran)))
		it.Option = append(it.Option, cache.NewOption(118, random(KHANG_TAT_CA, 0.8, ran)))
	} else if id == LAM_TINH_NGOC {
		// Vu khi
		it.Option = append(it.Option, vuKhi)
		it.Option = append(it.Option, cache.NewOption(103, random(MAX_ST_NGUOI, 0.8, ran)))
		it.Option = append(it.Option, cache.NewOption(125, -random(HP_TOI_DA, 0.8, ran)))

		// Trang bi
		it.Option = append(it.Option, trangBi)
		it.Option = append(it.Option, cache.NewOption(121, random(KHANG_ST_CHI_MANG, 0.8, ran)))
		it.Option = append(it.Option, cache.NewOption(120, -random(MOI_GIAY_HOI_PHUC_HP, 0.6, ran)))

		// Trang suc
		it.Option = append(it.Option, trangSuc)
		it.Option = append(it.Option, cache.NewOption(116, random(CHINH_XAC, 0.8, ran)))
		it.Option = append(it.Option, cache.NewOption(126, random(PHAN_DON, 0.8, ran)))
	} else if id == LUC_NGOC {
		// Vu khi
		it.Option = append(it.Option, vuKhi)
		it.Option = append(it.Option, cache.NewOption(105, random(ST_CHI_MANG, 0.8, ran)))
		it.Option = append(it.Option, cache.NewOption(116, -random(CHINH_XAC, 0.8, ran)))

		// Trang bi
		it.Option = append(it.Option, trangBi)
		it.Option = append(it.Option, cache.NewOption(125, random(HP_TOI_DA, 0.8, ran)))
		it.Option = append(it.Option, cache.NewOption(117, -random(MP_TOI_DA, 0.6, ran)))

		// Trang suc
		it.Option = append(it.Option, trangSuc)
		it.Option = append(it.Option, cache.NewOption(117, random(MP_TOI_DA, 0.8, ran)))
		it.Option = append(it.Option, cache.NewOption(124, random(MAX_GIAM_TRU_ST, 0.8, ran)))
	}
	it.SetUpgrade(byte(upgrade))
	it.Option = append(it.Option, cache.NewOption(EXP_ID, 0))
	it.Option = append(it.Option, cache.NewOption(GIA_KHAM_OPTION_ID, 800000))
	return it
}

func random(max int, percent float32, ran bool) int {
	if !ran {
		return max
	}
	return NextInt1(int(percent*float32(max)), max)
}

func (this *ItemFactory) ItemDefaultMat(id int) IItem {
	if id < 685 || id > 694 {
		panic(eris.New("Id is not eye item"))
	}
	item := this.ItemDefaultLock(id, true).(*Item)
	item.Sale = 5
	if id == Geningan {
		item.Option = append(item.Option, cache.NewOption(6, 1000))
		item.Option = append(item.Option, cache.NewOption(87, 500))
		item.SetUpgrade(1)
	} else if id == Chuuningan {
		item.Option = append(item.Option, cache.NewOption(6, 2000))
		item.Option = append(item.Option, cache.NewOption(87, 750))
		item.SetUpgrade(2)
	} else if id == Jougan {
		item.Option = append(item.Option, cache.NewOption(6, 3000))
		item.Option = append(item.Option, cache.NewOption(87, 1000))
		item.Option = append(item.Option, cache.NewOption(79, 25))
		item.SetUpgrade(3)
	} else if id == Seningan {
		item.Option = append(item.Option, cache.NewOption(6, 4000))
		item.Option = append(item.Option, cache.NewOption(87, 1250))
		item.Option = append(item.Option, cache.NewOption(79, 25))
		item.SetUpgrade(4)
	} else if id == Kyubigan {
		item.Option = append(item.Option, cache.NewOption(6, 5000))
		item.Option = append(item.Option, cache.NewOption(87, 1500))
		item.Option = append(item.Option, cache.NewOption(79, 25))
		item.SetUpgrade(5)
	} else if id == Rinnegan {
		item.Option = append(item.Option, cache.NewOption(6, 6000))
		item.Option = append(item.Option, cache.NewOption(87, 1750))
		item.Option = append(item.Option, cache.NewOption(79, 25))
		item.SetUpgrade(6)
	} else if id == Sharingan {
		item.Option = append(item.Option, cache.NewOption(6, 7000))
		item.Option = append(item.Option, cache.NewOption(87, 2250))
		item.Option = append(item.Option, cache.NewOption(79, 25))
		item.Option = append(item.Option, cache.NewOption(64, 0))
		item.SetUpgrade(7)
	} else if id == Tenseigan {
		item.Option = append(item.Option, cache.NewOption(6, 8000))
		item.Option = append(item.Option, cache.NewOption(87, 2250))
		item.Option = append(item.Option, cache.NewOption(79, 25))
		item.Option = append(item.Option, cache.NewOption(64, 0))
		item.SetUpgrade(8)
	} else if id == Ketsuryugan {
		item.Option = append(item.Option, cache.NewOption(6, 9000))
		item.Option = append(item.Option, cache.NewOption(87, 2500))
		item.Option = append(item.Option, cache.NewOption(79, 25))
		item.Option = append(item.Option, cache.NewOption(64, 0))
		item.SetUpgrade(9)
	} else if id == Sukaigan {
		item.Option = append(item.Option, cache.NewOption(6, 10000))
		item.Option = append(item.Option, cache.NewOption(87, 2725))
		item.Option = append(item.Option, cache.NewOption(79, 25))
		item.Option = append(item.Option, cache.NewOption(64, 0))
		item.Option = append(item.Option, cache.NewOption(113, 5000))
		item.SetUpgrade(10)
	}
	return item
}

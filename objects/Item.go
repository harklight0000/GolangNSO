package objects

import (
	. "nso/ainterfaces"
	"nso/cache"
	. "nso/constants"
	"nso/entity"
)

func NewItem(entity entity.ItemEntity) *Item {
	this := &Item{data: entity}
	this.ID = this.data.ID
	this.Sale = entity.SaleCoinLock
	this.Quantity = 1
	this.Upgrade = 0
	this.Index = 0
	this.BuyCoin = 0
	this.BuyCoinLock = 0
	this.BuyGold = 0
	this.Sys = 0
	if entity.IsExpires == 1 {
		this.isExpired = true
		this.TimeBuy = CurrentTimeMillis()
	} else {
		this.TimeBuy = 0
	}
	return this
}

type Item struct {
	data        entity.ItemEntity
	json        cache.ItemJSON
	Expires     int64
	ID          int16
	Sale        int
	Quantity    int
	Upgrade     byte
	Index       int
	BuyCoin     int
	BuyCoinLock int
	BuyGold     int
	Sys         byte
	TimeBuy     int64
	Ngocs       []Item
	Option      []cache.Option
	isExpired   bool
	isLock      bool
}

func (this *Item) IsTypeMount() bool {
	return this.data.Type >= 29 && this.data.Type <= 33
}

func (this *Item) Clone() IItem {
	other := NewItem(this.data)
	other.Quantity = this.Quantity
	other.Upgrade = this.Upgrade
	other.Index = this.Index
	other.BuyCoin = this.BuyCoin
	other.BuyCoinLock = this.BuyCoinLock
	other.BuyGold = this.BuyGold
	other.Sys = this.Sys
	other.TimeBuy = this.TimeBuy
	for _, ngoc := range this.Ngocs {
		other.Ngocs = append(other.Ngocs, ngoc)
	}
	for _, option := range this.Option {
		other.Option = append(other.Option, option)
	}
	return other
}

func (this *Item) IsExpired() bool {
	return this.isExpired && CurrentTimeMillis() >= this.Expires
}

func (this *Item) IsExpiredEggDaemon() bool {
	return CurrentTimeMillis()-this.TimeBuy >= this.Expires
}

func (this *Item) GetIDJiraiNam(_type int) int {
	switch _type {
	case 0:
		return 746
	case 1:
		return 747
	case 2:
		return 712
	case 3:
		return 713
	case 4:
		return 748
	case 5:
		return 752
	case 6:
		return 751
	case 7:
		return 750
	case 8:
		return 749
	}
	return -1
}

func (this *Item) GetIDJiraiNu(_type int) int {
	switch _type {
	case 0:
		return 753
	case 1:
		return 754
	case 2:
		return 715
	case 3:
		return 716
	case 4:
		return 755
	case 5:
		return 759
	case 6:
		return 758
	case 7:
		return 757
	case 8:
		return 756
	}
	return -1
}

func (this *Item) GetUpgradeMax() int {
	level := this.data.Level
	switch {
	case level >= 1 && level < 20:
		return 4
	case level >= 20 && level < 40:
		return 8
	case level >= 40 && level < 50:
		return 12
	case level >= 50 && level < 60:
		return 14
	default:
		return 16
	}
}

func (this *Item) UpgradeNext(next byte) {
	this.SetUpgrade(this.GetUpgrade() + next)
	if this.Option == nil {
		return
	}
	for i, option := range this.Option {
		id := option.ID
		switch {
		case id == 6 || id == 7:
			option.Param += 15 * int(next)
		case id == 8 || id == 9 || id == 19:
			option.Param += 10 * int(next)
		case (id >= 10 && id <= 15) || id == 17 || id == 18 || id == 20:
			option.Param += 5 * int(next)
		case id >= 21 && id <= 26:
			option.Param += 150 * int(next)
		case id == 16:
			option.Param += 3 * int(next)
		}
		this.Option[i] = option
	}
}

func (this *Item) IsTypeBody() bool {
	return this.data.Type >= TYPE_NON && this.data.Type <= TYPE_BI_KIP
}

func (this *Item) IsTypeNgocKham() bool {
	return this.data.Type == TYPE_NGOC
}

func (this *Item) GetData() *entity.ItemEntity {
	return &this.data
}

func (this *Item) IsPrecious() bool {
	id := this.ID
	return id == 383 || id == 384 || id == 385 || id == 308 || id == 309 || id == 353 || id == 652 || id == 653 || id == 654 || id == 695 || id >= 685 && id <= 704 || id == 655 || id == 599 || id == 600 || id == 605 || id == 597 || id == 602 || id == 603
}

func (this *Item) PercentAppear() int {
	id := this.ID
	switch id {
	case 599:
		fallthrough
	case 600:
		return 10
	case 605:
		return 2
	case 383:
		return 50
	case 384:
		return 20
	case 385:
		fallthrough
	case 687:
		fallthrough
	case 689:
		fallthrough
	case 686:
		fallthrough
	case 688:
		fallthrough
	case 690:
		fallthrough
	case 691:
		fallthrough
	case 692:
		fallthrough
	case 693:
		fallthrough
	case 694:
		return 2
	case 308:
		fallthrough
	case 309:
		fallthrough
	case 653:
		fallthrough
	case 654:
		fallthrough
	case 655:
		return 40
	case 685:
		return 3
	case 695:
		return 100
	case 696:
		return 80
	case 455:
		return 30
	case 456:
		return 15
	case 457:
		return 3
	case 545:
		return 30
	case 454:
		return 30
	case 697:
		return 20
	case 698:
		return 15
	case 699:
		return 10
	case 700:
		return 7
	case 701:
		return 6
	case 702:
		return 5
	case 703:
		return 4
	case 704:
		return 3
	default:
		return 100
	}
}

func (this *Item) FindParamById(id int) int {
	if this.Option == nil {
		return 0
	}
	for _, option := range this.Option {
		if option.ID == id {
			return option.Param
		}
	}
	return 0
}

func (this *Item) IsLock() bool {
	return this.isLock
}

func (this *Item) SetLock(lock bool) {
	this.isLock = lock
}

func (this *Item) GetUpgrade() byte {
	return this.Upgrade
}

func (this *Item) SetUpgrade(upgrade byte) {
	this.Upgrade = upgrade
}

func (this *Item) IsTrangSuc() bool {
	_type := this.data.Type
	return _type == TYPE_DAY_CHUYEN || _type == TYPE_NHAN || _type == TYPE_BOI || _type == TYPE_BUA
}

func (this *Item) IsYoroi() bool {
	return this.data.Type == TYPE_YOROI
}

func (this *Item) IsTrangPhuc() bool {
	_type := this.data.Type
	return _type == TYPE_NON || _type == TYPE_AO || _type == TYPE_QUAN || _type == TYPE_GIAY || _type == TYPE_GANG
}

func (this *Item) IsTypeTask() bool {
	_type := this.data.Type
	return _type == TYPE_TASK1 || _type == TYPE_TASK2 || _type == TYPE_TASK3
}

func (this *Item) IsVuKhi() bool {
	return this.data.Type == TYPE_VU_KHI
}

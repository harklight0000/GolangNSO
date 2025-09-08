package objects

type GameScr struct {
	crystals       []int
	upClothe       []int
	upAdorn        []int
	upWeapon       []int
	coinUpCrystals []int
	coinUpClothes  []int
	coinUpAdorns   []int
	coinUpWeapons  []int
	goldUps        []int
	maxPercents    []int
	arryenLuck     []int
	arrdayLuck     []byte
	optionBikiep   []int
	paramBikiep    []int
	percentBikiep  []int
}

func NewGameScr() *GameScr {
	this := &GameScr{}
	this.crystals = []int{1, 4, 16, 64, 256, 1024, 4096, 16384, 65536, 262144, 1048576, 3096576}
	this.upClothe = []int{4, 9, 33, 132, 177, 256, 656, 2880, 3968, 6016, 13440, 54144, 71680, 108544, 225280, 1032192}
	this.upAdorn = []int{6, 14, 50, 256, 320, 512, 1024, 5120, 6016, 9088, 19904, 86016, 108544, 166912, 360448, 1589248}
	this.upWeapon = []int{18, 42, 132, 627, 864, 1360, 2816, 13824, 17792, 26880, 54016, 267264, 315392, 489472, 1032192, 4587520}
	this.coinUpCrystals = []int{10, 40, 160, 640, 2560, 10240, 40960, 163840, 655360, 1310720, 3932160, 11796480}
	this.crystals = []int{1, 4, 16, 64, 256, 1024, 4096, 16384, 65536, 262144, 1048576, 3096576}
	this.coinUpClothes = []int{120, 270, 990, 3960, 5310, 7680, 19680, 86400, 119040, 180480, 403200, 1624320, 2150400, 3256320, 6758400, 10137600}
	this.coinUpAdorns = []int{180, 420, 1500, 7680, 9600, 15360, 30720, 153600, 180480, 272640, 597120, 2580480, 3256320, 5007360, 10813440, 16220160}
	this.coinUpWeapons = []int{540, 1260, 3960, 18810, 25920, 40800, 84480, 414720, 533760, 806400, 1620480, 8017920, 9461760, 14684160, 22026240, 33039360}
	this.goldUps = []int{1, 2, 3, 4, 5, 10, 15, 20, 50, 100, 150, 200, 300, 400, 500, 600}
	this.maxPercents = []int{80, 75, 70, 65, 60, 55, 50, 45, 40, 35, 30, 25, 20, 15, 10, 5}
	this.arryenLuck = []int{1_000_000, 2_000_000}
	this.arrdayLuck = []byte{3, 7, 15, 30}
	this.optionBikiep = []int{86, 87, 88, 89, 90, 91, 92, 95, 96, 97, 98, 84, 100}
	this.paramBikiep = []int{50, 500, 250, 250, 250, 20, 10, 10, 100, 100, 100, 10, 50, 10}
	this.percentBikiep = []int{80, 75, 70, 65, 60, 55, 50, 45, 40, 35, 30, 25, 20, 15, 10, 5}
	return this
}

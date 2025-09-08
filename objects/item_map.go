package objects

type ItemMap struct {
	Item        *Item
	X           int16
	Y           int16
	ItemMapID   int16
	RemoveDelay int64
	IDMaster    int
	IsVisible   bool
	NextRefresh bool
}

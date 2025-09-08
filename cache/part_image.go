package cache

type PartImage struct {
	ID int  `json:"id" bson:"id"`
	Dx int8 `json:"dx" bson:"dx"`
	Dy int8 `json:"dy" bson:"dy"`
}

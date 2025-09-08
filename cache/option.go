package cache

type Option struct {
	ID    int `json:"id" bson:"id"`
	Param int `json:"param" bson:"param"`
}

func NewOption(ID int, param int) Option {
	return Option{ID: ID, Param: param}
}

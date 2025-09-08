package cache

type ItemJSON struct {
	IsLock      bool  `json:"is_lock" bson:"is_lock"`
	Sale        int   `json:"sale" bson:"sale"`
	Quantity    int   `json:"quantity" bson:"quantity"`
	Upgrade     byte  `json:"upgrade" bson:"upgrade"`
	Index       int   `json:"index" bson:"index"`
	ID          int16 `json:"id" bson:"id"`
	IsExpires   bool  `json:"isExpires" bson:"isExpires"`
	Expires     int64 `json:"expires" bson:"expires"`
	BuyCoin     int   `json:"buyCoin" bson:"buyCoin"`
	BuyCoinLock int   `json:"buyCoinLock" bson:"buyCoinLock"`
	BuyGold     int   `json:"buyGold" bson:"buyGold"`
	Sys         byte  `json:"sys" bson:"sys"`
	TimeBuy     int64 `json:"timeBuy" bson:"timeBuy"`

	Ngocs  []ItemJSON `json:"ngocs" bson:"ngocs"`
	Option []Option   `json:"option" bson:"option"`
}

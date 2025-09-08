package cache

type Friend struct {
	Name  string `json:"name" bson:"name"`
	Agree bool   `json:"agree" bson:"agree"`
}

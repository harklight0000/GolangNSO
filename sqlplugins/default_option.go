package sqlplugins

import "go.mongodb.org/mongo-driver/bson"

type defaultOption struct {
}

func (d *defaultOption) ToMongoDBOption() bson.M {
	return bson.M{}
}

func (d *defaultOption) ToSQLOption() string {
	return " "
}

func DefaultOption() *defaultOption {
	return &defaultOption{}
}

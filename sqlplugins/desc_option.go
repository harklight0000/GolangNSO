package sqlplugins

import "go.mongodb.org/mongo-driver/bson"

type DescOption struct {
	basicOption
}

func Desc(field string) *DescOption {
	return &DescOption{
		basicOption{
			Name: field,
		},
	}
}

func (d *DescOption) ToMongoDBOption() bson.M {
	return bson.M{
		"sort": bson.M{
			d.Name: -1,
		},
	}
}

func (d DescOption) ToSQLOption() string {
	return " ORDER BY " + d.Name + " DESC "
}

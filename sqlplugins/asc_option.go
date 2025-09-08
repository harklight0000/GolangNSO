package sqlplugins

import "go.mongodb.org/mongo-driver/bson"

type AscOption struct {
	basicOption
}

func Asc(fieldName string) *AscOption {
	return &AscOption{
		basicOption{Name: fieldName},
	}
}

func (a *AscOption) ToMongoDBOption() bson.M {
	return bson.M{
		"sort": bson.M{
			a.Name: 1,
		},
	}
}

func (a *AscOption) ToSQLOption() string {
	return " ORDER BY " + a.Name + " ASC "
}

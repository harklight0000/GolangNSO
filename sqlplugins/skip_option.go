package sqlplugins

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func SkipOption(number int) *skipOption {
	return &skipOption{
		basicOption{
			Value: number,
		},
	}
}

type skipOption struct {
	basicOption
}

func (s *skipOption) ToMongoDBOption() bson.M {
	return bson.M{
		"skip": s.Value,
	}
}

func (s *skipOption) ToSQLOption() string {
	return fmt.Sprintf(" OFFSET %d ", s.Value)
}

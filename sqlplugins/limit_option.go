package sqlplugins

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"nso/ainterfaces"
	"nso/errs"
	"nso/logging"
	"strconv"
)

type LimitOption struct {
	basicOption
}

func Limit(number int) ainterfaces.IOption {
	return &LimitOption{
		basicOption{
			Value: number,
		},
	}
}

func (l *LimitOption) ToMongoDBOption() bson.M {
	return bson.M{
		"limit": l.Value,
	}
}

func (l *LimitOption) ToSQLOption() string {
	limit, err := strconv.Atoi(fmt.Sprintf("%v", l.Value))
	if err != nil {
		logging.Logger.Info("Error converting limit to int!! " + errs.ToString(err))
		return " LIMIT 1 "
	}
	return " limit " + strconv.Itoa(limit) + " "
}

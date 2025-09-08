package sqlplugins

import (
	"go.mongodb.org/mongo-driver/bson"
	. "nso/ainterfaces"
	"strings"
)

type defaultCondition struct {
}

func DefaultCondition() ICondition {
	return &defaultCondition{}
}

func (d *defaultCondition) ToSQLCondition() string {
	return ""
}

func (d *defaultCondition) ToMongoCondition() bson.M {
	return bson.M{}
}

func (a *andCondition) ToSQLCondition() string {
	sb := &strings.Builder{}
	for i, cond := range a.Conditions {
		sb.WriteString(cond.ToSQLCondition())
		if i < len(a.Conditions)-1 {
			sb.WriteString(" AND ")
		}
	}
	return sb.String()
}

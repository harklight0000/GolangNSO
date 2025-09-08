package sqlplugins

import (
	"go.mongodb.org/mongo-driver/bson"
	. "nso/ainterfaces"
	"strings"
)

type orCondition struct {
	compoundCondition
}

func Or(cond ...ICondition) *orCondition {
	return &orCondition{compoundCondition{
		Conditions: cond,
	}}
}

func (o *orCondition) ToSQLCondition() string {
	sb := &strings.Builder{}
	for i, cond := range o.Conditions {
		sb.WriteString(cond.ToSQLCondition())
		if i < len(o.Conditions)-1 {
			sb.WriteString(" OR ")
		}
	}
	return sb.String()
}

func (o *orCondition) ToMongoCondition() bson.M {
	cond := bson.M{}
	conds := make([]bson.M, 0)
	for _, c := range o.Conditions {
		conds = append(conds, c.ToMongoCondition())
	}
	cond["$or"] = conds
	return cond
}

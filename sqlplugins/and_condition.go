package sqlplugins

import (
	"go.mongodb.org/mongo-driver/bson"
	. "nso/ainterfaces"
)

type andCondition struct {
	compoundCondition
}

func And(cond ...ICondition) *andCondition {
	return &andCondition{
		compoundCondition{
			Conditions: cond,
		},
	}
}

func (a *andCondition) ToMongoCondition() bson.M {
	var cond = bson.M{}
	conds := make([]bson.M, 0)
	for _, c := range a.Conditions {
		conds = append(conds, c.ToMongoCondition())
	}
	cond["$and"] = conds
	return cond
}

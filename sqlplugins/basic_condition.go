package sqlplugins

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	. "nso/ainterfaces"
	"nso/logging"
)

type basicCondition struct {
	Field    string
	Value    interface{}
	Operator string
}

func (c *basicCondition) ToSQLCondition() string {
	if c.Operator == "like" {
		return fmt.Sprintf("`%s` LIKE '%v'", c.Field, c.Value)
	}
	// check type of value
	switch c.Value.(type) {
	case string:
		return fmt.Sprintf(" `%s` %s '%v' ", c.Field, c.Operator, c.Value)
	default:
		return fmt.Sprintf(" `%s` %s %v ", c.Field, c.Operator, c.Value)
	}
}

func (c *basicCondition) ToMongoCondition() bson.M {
	if c.Operator == "=" {
		if c.Field == "_id" {
			var err error
			switch c.Value.(type) {
			case string:
				c.Value, err = primitive.ObjectIDFromHex(c.Value.(string))
				if err != nil {
					logging.Logger.Error("Error converting string to ObjectID", zap.Error(err))
				}
			}
		}
		return bson.M{
			c.Field: bson.M{"$eq": c.Value},
		}
	} else if c.Operator == "<>" || c.Operator == "!=" {
		return bson.M{
			c.Field: bson.M{"$ne": c.Value},
		}
	} else if c.Operator == "<" {
		return bson.M{
			c.Field: bson.M{"$lt": c.Value},
		}
	} else if c.Operator == "<=" {
		return bson.M{c.Field: bson.M{"$lte": c.Value}}
	} else if c.Operator == ">" {
		return bson.M{
			c.Field: bson.M{"$gt": c.Value},
		}
	} else if c.Operator == ">=" {
		return bson.M{
			c.Field: bson.M{"$gte": c.Value},
		}
	}
	logging.Logger.Info("Unknown operator: " + c.Operator)
	return bson.M{}
}

func Lte(field string, value interface{}) ICondition {
	return &basicCondition{
		Field:    field,
		Value:    value,
		Operator: "<=",
	}
}

func Eq(field string, value interface{}) ICondition {
	return &basicCondition{
		Field:    field,
		Value:    value,
		Operator: "=",
	}
}

func Neq(field string, value interface{}) ICondition {

	return &basicCondition{
		Field:    field,
		Value:    value,
		Operator: "<>",
	}
}

func Lt(field string, value interface{}) ICondition {
	return &basicCondition{
		Field:    field,
		Value:    value,
		Operator: "<",
	}
}

func Gte(field string, value interface{}) ICondition {
	return &basicCondition{
		Field:    field,
		Value:    value,
		Operator: ">=",
	}
}

func Gt(field string, value interface{}) ICondition {
	return &basicCondition{
		Field:    field,
		Value:    value,
		Operator: ">",
	}
}

func Like(field string, value interface{}) ICondition {
	return &basicCondition{
		Field:    field,
		Value:    value,
		Operator: "like",
	}
}

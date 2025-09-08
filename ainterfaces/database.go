package ainterfaces

import (
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Bson interface {
	bson.M | bson.D | bson.A
}

type ICondition interface {
	ToSQLCondition() string
	ToMongoCondition() bson.M
}

type IOption interface {
	ToMongoDBOption() bson.M
	ToSQLOption() string
}

type IDatabase interface {
	Create(value interface{}) error
	Update(table string, value interface{}, condition ICondition) error
	Delete(table string, condition ICondition) error
	FindOne(result interface{}, condition ICondition) error
	FindOneOfTable(table string, result interface{}, condition ICondition) error
	FindMany(result interface{}, condition ICondition) error
	FindAll(result interface{}) error
	FindManyWithConditionAndOption(result interface{}, condition ICondition, options IOption) error
	FindManyWithOption(result interface{}, options IOption) error
	Unwrap(v interface{})
	io.Closer
}

package database

import (
	"context"
	"github.com/rotisserie/eris"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	. "nso/ainterfaces"
	"nso/logging"
	. "nso/sqlplugins"
	"reflect"
)

type mongoDBAdapter struct {
	*mongo.Database
	ctx context.Context
}

func (m *mongoDBAdapter) FindOneOfTable(table string, result interface{}, condition ICondition) error {
	if condition == nil {
		condition = DefaultCondition()
	}
	var bsonResult bson.M
	err := m.Database.Collection(table).FindOne(m.ctx, condition.ToMongoCondition()).Decode(&bsonResult)
	if err != nil {
		return eris.Wrapf(err, "Fail to find from mongo")
	}
	err = convertBSONToStruct(bsonResult, result)
	if err != nil {
		return eris.Wrapf(err, "Fail to convert from bson object")
	}
	return nil
}

func (m *mongoDBAdapter) FindManyWithOption(result interface{}, option IOption) error {
	return m.FindManyWithConditionAndOption(result, DefaultCondition(), option)
}

func (m *mongoDBAdapter) FindMany(result interface{}, condition ICondition) error {
	return m.FindManyWithConditionAndOption(result, condition, DefaultOption())
}

func (m *mongoDBAdapter) FindAll(result interface{}) error {
	return m.FindManyWithConditionAndOption(result, DefaultCondition(), DefaultOption())
}

func (m *mongoDBAdapter) Close() error {
	return eris.Wrap(m.Database.Client().Disconnect(m.ctx), "Fail to disconnect from mongo")
}

func NewMongoDatabaseAdapter(db *mongo.Database) IDatabase {
	ctx := context.Background()
	return &mongoDBAdapter{Database: db, ctx: ctx}
}

func (m *mongoDBAdapter) Create(value interface{}) error {
	table := getTableName(value)
	bsonObject, err := convertStructToBSON[bson.M](value)
	if err != nil {
		return eris.Wrapf(err, "Fail to convert to bson object")
	}
	bsonObject["_id"] = primitive.NewObjectID()
	result, err := m.Database.Collection(table).InsertOne(m.ctx, bsonObject)
	if err != nil {
		return eris.Wrapf(err, "Fail to insert to mongo")
	}
	if result.InsertedID == nil {
		return eris.Wrapf(err, "Fail to insert to mongo")
	}
	if err = m.FindOne(value, Eq("_id", result.InsertedID)); err != nil {
		return eris.Wrapf(err, "Fail to find inserted value")
	}
	return nil
}

func (m *mongoDBAdapter) Update(table string, value interface{}, condition ICondition) error {
	if condition == nil {
		condition = DefaultCondition()
	}
	_type := reflect.TypeOf(value)
	if _type.Kind() != reflect.Struct {
		if _type.Kind() == reflect.Ptr {
			_type = _type.Elem()
			if _type.Kind() != reflect.Struct {
				return eris.New("Fail to update, value is not a struct")
			}
		} else {
			return eris.New("Invalid value type")
		}
	}
	updateValues, err := convertStructToBSON[bson.M](value)
	delete(updateValues, "_id")
	if err != nil {
		return eris.Wrap(err, "Error convert struct to bson update")
	}
	updated, err := m.Database.Collection(table).UpdateOne(m.ctx, condition.ToMongoCondition(),
		bson.D{{"$set", updateValues}})
	if err != nil {
		return eris.Wrapf(err, "Fail to update to mongo")
	}
	if updated.ModifiedCount == 0 {
		return eris.Wrapf(err, "Fail to update to mongo")
	}
	return nil
}

func (m *mongoDBAdapter) Delete(table string, condition ICondition) error {
	if condition == nil {
		condition = DefaultCondition()
	}
	one, err := m.Database.Collection(table).DeleteOne(m.ctx, condition.ToMongoCondition())
	if err != nil {
		return eris.Wrapf(err, "Fail to delete from mongo")
	}
	if one.DeletedCount == 0 {
		return eris.Wrapf(err, "Not found to delete")
	}
	return nil
}

func (m *mongoDBAdapter) FindOne(result interface{}, condition ICondition) error {
	return m.FindOneOfTable(getTableName(result), result, condition)
}

func (m *mongoDBAdapter) FindManyWithConditionAndOption(result interface{}, condition ICondition, option IOption) error {
	table := getTableName(result)
	if condition == nil {
		condition = DefaultCondition()
	}
	if option == nil {
		option = DefaultOption()
	}

	var findOptions = options.Find()
	for key, v := range option.ToMongoDBOption() {
		switch key {
		case "limit":
			limit := v.(int)
			findOptions.SetLimit(int64(limit))
		case "skip":
			skip := int64(v.(int))
			findOptions.SetSkip(skip)
		case "sort":
			sort := v.(bson.M)
			findOptions.SetSort(sort)
		}
	}

	mongoCursor, err := m.Database.Collection(table).Find(m.ctx, condition.ToMongoCondition(), findOptions)

	if err != nil {
		return eris.Wrapf(err, "Fail to find from mongo")
	}
	defer func(mongCursor *mongo.Cursor, ctx context.Context) {
		err := mongCursor.Close(ctx)
		if err != nil {
			logging.Logger.Info("Fail to close mongo cursor" + err.Error())
		}
	}(mongoCursor, m.ctx)
	var bsonResults []bson.M
	err = mongoCursor.All(m.ctx, &bsonResults)
	err = convertBSONObjectsToObjects(bsonResults, result)
	if err != nil {
		return eris.Wrapf(err, "Fail to convert from bson object")
	}
	return nil
}

func (m *mongoDBAdapter) Unwrap(v interface{}) {
	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		reflect.ValueOf(v).Elem().Set(reflect.ValueOf(m.Database))
	} else {
		logging.Logger.Info("Unwrap failed, v is not a pointer")
	}
}

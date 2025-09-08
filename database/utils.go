package database

import (
	"github.com/rotisserie/eris"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"gorm.io/gorm/schema"
	. "nso/ainterfaces"
	. "nso/logging"
	"reflect"
	"strings"
)

func getTableName(value interface{}) string {
	table, ok := value.(schema.Tabler)
	if ok {
		return table.TableName()
	}
	_type := reflect.TypeOf(value)
	switch _type.Kind() {
	case reflect.Ptr:
		elementType := _type.Elem()
		if elementType.Kind() == reflect.Slice {
			if elementType.Elem().Kind() == reflect.Struct {
				elementInterface := reflect.New(elementType.Elem()).Interface()
				table, ok := elementInterface.(schema.Tabler)
				if ok {
					return table.TableName()
				}
				return strings.ToLower(elementType.Elem().Name())
			} else if elementType.Elem().Kind() == reflect.Ptr {
				elementInterface := reflect.New(elementType.Elem().Elem()).Interface()
				table, ok := elementInterface.(schema.Tabler)
				if ok {
					return table.TableName()
				}
				return strings.ToLower(elementType.Elem().Elem().Name())
			}
		}
		elementIf := reflect.New(_type.Elem()).Interface()
		table, ok := elementIf.(schema.Tabler)
		if ok {
			return table.TableName()
		}
		return strings.ToLower(elementType.Name())
	case reflect.Struct:
		return strings.ToLower(_type.Name())
	case reflect.Slice:
		_typeElement := _type.Elem()
		if _typeElement.Kind() == reflect.Ptr {
			elementInterface := reflect.New(_typeElement.Elem()).Interface()
			table, ok := elementInterface.(schema.Tabler)
			if ok {
				return table.TableName()
			}
			return strings.ToLower(_typeElement.Elem().Name())
		}
		elementInterface := reflect.New(_typeElement).Interface()
		table, ok := elementInterface.(schema.Tabler)
		if ok {
			return table.TableName()
		}
		return strings.ToLower(_typeElement.Name())
	}
	Logger.Info("Unsupported type", zap.String("type", _type.String()))
	return ""
}

func convertBSONObjectsToObjects[B Bson](inputs []B, result interface{}) error {
	var err error
	if err != nil {
		return eris.Wrapf(err, "Fail to find from mongo")
	}
	value := reflect.ValueOf(result)
	if value.Kind() != reflect.Ptr {
		return eris.Wrapf(nil, "Result must be a pointer")
	}
	sliceValue := value.Elem()
	// Get slice type
	sliceType := value.Type().Elem()
	if sliceType.Kind() != reflect.Slice {
		return eris.Wrapf(nil, "result element must be a slice")
	}
	componentType := sliceType.Elem()
	var isPointer = false
	if componentType.Kind() == reflect.Pointer {
		isPointer = true
		componentType = componentType.Elem()
	}
	for _, bsonResult := range inputs {
		r := reflect.New(componentType)
		err = convertBSONToStruct[B](bsonResult, r.Interface())
		if err != nil {
			return eris.Wrapf(err, "Fail to convert bson to struct")
		}
		if isPointer {
			sliceValue = reflect.Append(sliceValue, r)
		} else {
			sliceValue = reflect.Append(sliceValue, r.Elem())
		}
	}
	value.Elem().Set(sliceValue)
	return nil
}

func convertStructToBSON[B Bson](value interface{}) (B, error) {
	var bsonObject B
	typeValue := reflect.TypeOf(value)
	if typeValue.Kind() == reflect.Slice {
		return bsonObject, eris.New("Slice is not supported")
	}
	if typeValue.Kind() == reflect.Ptr {
		typeValue = typeValue.Elem()
		value = reflect.ValueOf(value).Elem().Interface()
	}
	bytes, err := bson.Marshal(value)
	if err != nil {
		return bsonObject, eris.Wrap(err, "failed to marshal value to json")
	}
	err = bson.Unmarshal(bytes, &bsonObject)
	if err != nil {
		return bsonObject, eris.Wrap(err, "failed to unmarshal json to bson")
	}
	return bsonObject, nil
}

func convertBSONToStruct[B Bson](value B, result interface{}) error {
	typeResult := reflect.TypeOf(result)
	if typeResult.Kind() != reflect.Pointer {
		return eris.New("result must be pointer")
	}
	elementType := typeResult.Elem()
	if elementType.Kind() != reflect.Struct {
		return eris.New("element of pointer must be a struct")
	}
	doc, err := bson.Marshal(value)
	if err != nil {
		return eris.Wrap(err, "failed to marshal bson to json")
	}
	err = bson.Unmarshal(doc, result)
	if err != nil {
		return eris.Wrap(err, "failed to unmarshal json to struct")
	}
	return nil
}

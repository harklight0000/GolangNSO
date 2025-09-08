package database

import (
	"context"
	"encoding/json"
	"github.com/rotisserie/eris"
	"gorm.io/gorm/schema"
	"reflect"
	"strings"
)

type JSONSerializer struct {
	schema.JSONSerializer
}

func (j JSONSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	if dbValue == nil {
		fieldValue := reflect.New(field.FieldType)
		err = eris.Wrap(json.Unmarshal([]byte("[]"), fieldValue.Interface()), "Error when unmarshal json of field name 1 "+field.Name)
		field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
		return
	} else {
		str := string(dbValue.([]byte))
		if strings.Contains(str, "\"param\"500") {
			str = strings.Replace(str, "\"param\"500", "\"param\":500", -1)
			dbValue = []byte(str)
		}
		if strings.Contains(str, "}{") {
			str = strings.Replace(str, "}{", "},{", -1)
			dbValue = []byte(str)
		}
		err = eris.Wrap(j.JSONSerializer.Scan(ctx, field, dst, dbValue), "Error when unmarshal json of field name 2 "+field.Name)
	}
	return
}

func (j JSONSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	if fieldValue == nil {
		return "[]", nil
	} else {
		return j.JSONSerializer.Value(ctx, field, dst, fieldValue)
	}
}

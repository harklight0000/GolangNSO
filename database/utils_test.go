package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"nso/errs"
	"testing"
)

func TestConvertBSONMToStruct(t *testing.T) {
	type test struct {
		ID   string `bson:"_id" json:"_id"`
		Name string `bson:"name" json:"name"`
		Age  int    `bson:"age" json:"age"`
	}
	var test1 test
	err := convertBSONToStruct[bson.M](bson.M{"_id": "1", "name": "test", "age": 1}, &test1)
	if err != nil {
		t.Error(errs.ToString(err))
		return
	}
	if test1.ID != "1" {
		t.Error("Cannot convert bson to struct")
	}
	if test1.Name != "test" {
		t.Error("Cannot convert bson to struct")
	}
	if test1.Age != 1 {
		t.Error("Cannot convert bson to struct")
	}
}

func TestConvertBSONDToStruct(t *testing.T) {
	type test struct {
		ID   string `bson:"_id" json:"_id"`
		Name string `bson:"name" json:"name"`
		Age  int    `bson:"age" json:"age"`
	}
	var test1 test
	err := convertBSONToStruct[bson.D](bson.D{{"_id", "1"}, {"name", "test"}, {"age", 1}}, &test1)
	if err != nil {
		t.Error(errs.ToString(err))
		return
	}
	if test1.ID != "1" {
		t.Error("Cannot convert bson to struct")
	}
	if test1.Name != "test" {
		t.Error("Cannot convert bson to struct")
	}
	if test1.Age != 1 {
		t.Error("Cannot convert bson to struct")
	}
}

func TestConvertStructToBSONM(t *testing.T) {
	type test struct {
		ID   string `bson:"_id" json:"_id"`
		Name string `bson:"name" json:"name"`
		Age  int    `bson:"age" json:"age"`
	}
	test1 := test{ID: "1", Name: "test", Age: 1}
	bson1, err := convertStructToBSON[bson.M](test1)
	if err != nil {
		t.Error(errs.ToString(err))
		return
	}
	if bson1["_id"] != "1" {
		t.Error("_id is not present")
	}
	if bson1["name"] != "test" {
		t.Error("name is not present")
	}
	if bson1["age"] != int32(1) {
		t.Error("age is not present")
	}
}

func TestConvertStructToBSOND(t *testing.T) {
	type test struct {
		ID   string `bson:"_id" json:"_id"`
		Name string `bson:"name" json:"name"`
		Age  int    `bson:"age" json:"age"`
	}
	test1 := test{ID: "1", Name: "test", Age: 1}
	bson1, err := convertStructToBSON[bson.D](&test1)
	if err != nil {
		t.Error(errs.ToString(err))
		return
	}
	if bson1.Map()["_id"] != "1" {
		t.Error("_id is not present")
	}
	if bson1.Map()["name"] != "test" {
		t.Error("name is not present")
	}
	if bson1.Map()["age"] != int32(1) {
		t.Error("age is not present")
	}
}

func TestConvertStructToBSON_Fail(t *testing.T) {
	type test struct {
		ID   string `bson:"_id" json:"_id"`
		Name string `bson:"name" json:"name"`
		Age  int    `bson:"age" json:"age"`
	}
	test1 := []test{{ID: "1", Name: "test", Age: 1}}
	_, err := convertStructToBSON[bson.D](test1)
	if err == nil {
		t.Error("Cannot convert struct to bson")
	}
}

func TestConvertBSONToStruct_Fail(t *testing.T) {
	type test struct {
		ID   string `bson:"_id" json:"_id"`
		Name string `bson:"name" json:"name"`
		Age  int    `bson:"age" json:"age"`
	}
	var test1 test
	err := convertBSONToStruct[bson.M](bson.M{"_id": "1", "name": "test", "age": 1}, test1)
	if err == nil {
		t.Error("Not passing pointer should fail")
	} else {
		t.Log(errs.ToString(err))
	}
	err = convertBSONToStruct(bson.M{"_id": "1", "name": "test", "age": 1}, &[]interface{}{})
	if err == nil {
		t.Error("Passing slice should fail")
	} else {
		t.Log(errs.ToString(err))
	}
}

func TestConvertBSONObjectsToObjects(t *testing.T) {

	var test1 []*Test3
	err := convertBSONObjectsToObjects[bson.M]([]bson.M{{"_id": "1", "name": "test", "age": 1}}, &test1)
	if err != nil {
		t.Error(errs.ToString(err))
		return
	}
	if test1[0].ID != "1" {
		t.Error("Id not match")
	}
	if test1[0].Name != "test" {
		t.Error("Name not match")
	}
	if test1[0].Age != 1 {
		t.Error("Age not match")
	}

	var test2 []*Test3
	err = convertBSONObjectsToObjects[bson.M]([]bson.M{{"_id": "1", "name": "test", "age": 1}}, &test2)
	if err != nil {
		t.Error(errs.ToString(err))
		return
	}
	if test2[0].ID != "1" {
		t.Error("Id not match")
	}
	if test2[0].Name != "test" {
		t.Error("Name not match")
	}
	if test2[0].Age != 1 {
		t.Error("Age not match")
	}
}

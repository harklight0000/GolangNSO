package database

import (
	"nso/ainterfaces"
	"nso/logging"
	. "nso/sqlplugins"
	"testing"
)

type Test1 struct {
	ID   int    `gorm:"primary_key"`
	_ID  string `gorm:"-" bson:"_id" json:"_id"`
	Name string `gorm:"name" bson:"name" json:"name"`
	Age  int    `gorm:"age" bson:"age" json:"age"`
}

func TestOptions_MongoDBAdapter(t *testing.T) {
	mongodbb := InitMongoDB()
	db := NewMongoDatabaseAdapter(mongodbb)
	defer tearDownMongoTest(mongodbb, db)
	count := initData(t, db)
	testConditions(t, db, count)
}

func initData(t *testing.T, db ainterfaces.IDatabase) int {
	var ts = []Test1{
		{_ID: "", Name: "John", Age: 30},
		{_ID: "", Name: "Mary", Age: 25},
		{_ID: "", Name: "Bob", Age: 20},
		{_ID: "", Name: "Jane", Age: 35},
		{_ID: "", Name: "Jack", Age: 40},
	}
	for _, tes := range ts {
		err := db.Create(&tes)
		if err != nil {
			t.Error(err)
		}
	}
	return len(ts)
}

func TestOption_SqlAdapter(t *testing.T) {
	sql := InitSQLDB()
	err := sql.Migrator().AutoMigrate(&Test1{})
	if err != nil {
		t.Error(err)
	}
	db := NewSQLDatabaseAdapter(sql)
	defer func(db ainterfaces.IDatabase) {
		err := sql.Migrator().DropTable(&Test1{})
		if err != nil {
			logging.Logger.Info("Fail to drop a table name Test1")
			return
		}
		err = db.Close()
		if err != nil {
			logging.Logger.Info("Fail to disconnect sql db")
		}
	}(db)
	count := initData(t, db)
	testConditions(t, db, count)
}

func testConditions(t *testing.T, db ainterfaces.IDatabase, count int) {

	var err error
	t.Run("NilOption", func(t *testing.T) {
		var ts2 []Test1
		err = db.FindManyWithConditionAndOption(&ts2, nil, nil)
		if err != nil {
			t.Error(err)
			return
		}
		if len(ts2) != count {
			t.Errorf("Expect %d, but got %d", count, len(ts2))
		}
	})

	t.Run("DefaultOption", func(t *testing.T) {
		var ts2 []Test1
		err = db.FindManyWithConditionAndOption(&ts2, DefaultCondition(), nil)
		if err != nil {
			t.Error(err)
			return
		}
		if len(ts2) != count {
			t.Errorf("Expect %d, but got %d", count, len(ts2))
		}
	})

	t.Run("Eq", func(t *testing.T) {
		var t3 []Test1
		err = db.FindManyWithConditionAndOption(&t3, Eq("name", "John"), nil)
		if err != nil {
			t.Error(err)
			return
		}
		if len(t3) != 1 {
			t.Errorf("Expect %d, but got %d", 1, len(t3))
		}
		if t3[0].Name != "John" {
			t.Errorf("Expect %s, but got %s", "John", t3[0].Name)
		}
	})

	t.Run("Gt", func(t *testing.T) {
		var t4 []Test1
		err = db.FindManyWithConditionAndOption(&t4, Gt("age", 30), nil)
		if err != nil {
			t.Error(err)
			return
		}
		if len(t4) != 2 {
			t.Errorf("Expect %d, but got %d", 2, len(t4))
		}
	})

	t.Run("Lt", func(t *testing.T) {
		var t5 []Test1
		err = db.FindMany(&t5, Lt("age", 30))
		if err != nil {
			t.Error(err)
			return
		}
		if len(t5) != 2 {
			t.Errorf("Expect %d, but got %d", 2, len(t5))
		}
	})

	t.Run("Gte", func(t *testing.T) {
		var t6 []Test1
		err = db.FindManyWithConditionAndOption(&t6, Gte("age", 30), nil)
		if err != nil {
			t.Error(err)
			return
		}
		if len(t6) != 3 {
			t.Errorf("Expect %d, but got %d", 3, len(t6))
		}
	})

	t.Run("Lte", func(t *testing.T) {
		var t7 []Test1
		err = db.FindManyWithConditionAndOption(&t7, Lte("age", 30), nil)
		if err != nil {
			t.Error(err)
			return
		}
		if len(t7) != 3 {
			t.Errorf("Expect %d, but got %d", 3, len(t7))
		}
	})
	t.Run("And", func(t *testing.T) {
		var t8 []Test1
		err = db.FindManyWithConditionAndOption(&t8, And(Eq("name", "John"), Gte("age", 30)), nil)
		if err != nil {
			t.Error(err)
			return
		}
		if len(t8) != 1 {
			t.Errorf("Expect %d, but got %d", 1, len(t8))
		}
		if t8[0].Name != "John" {
			t.Errorf("Expect %s, but got %s", "John", t8[0].Name)
		}
		if t8[0].Age != 30 {
			t.Errorf("Expect %d, but got %d", 30, t8[0].Age)
		}
	})
}

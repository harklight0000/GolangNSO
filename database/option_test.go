package database

import (
	"nso/ainterfaces"
	"nso/logging"
	"nso/sqlplugins"
	"testing"
)

func TestOption_AdapterSQL(t *testing.T) {
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
	testOptions(t, db, count)
}

func TestOption_AdapterMongo(t *testing.T) {
	mongodbb := InitMongoDB()
	db := NewMongoDatabaseAdapter(mongodbb)
	defer tearDownMongoTest(mongodbb, db)
	count := initData(t, db)
	testOptions(t, db, count)
}

func testOptions(t *testing.T, db ainterfaces.IDatabase, count int) {
	t.Run("Count", func(t *testing.T) {
		var ts []Test1
		err := db.FindManyWithOption(&ts, sqlplugins.Limit(count-1))
		if err != nil {
			t.Error(err)
			return
		}
		if len(ts) != count-1 {
			t.Errorf("Expect %d, but got %d", count-1, len(ts))
		}
	})
	t.Run("Asc", func(t *testing.T) {
		var ts []Test1
		err := db.FindManyWithOption(&ts, sqlplugins.Asc("age"))
		if err != nil {
			t.Error(err)
			return
		}
		if len(ts) != count {
			t.Errorf("Expect %d, but got %d", count, len(ts))
		}
		for i := 0; i < count-1; i++ {
			if ts[i].Age > ts[i+1].Age {
				t.Errorf("Expect %d, but got %d", ts[i].Age, ts[i+1].Age)
			}
		}
	})
	t.Run("Desc", func(t *testing.T) {
		var ts []Test1
		err := db.FindManyWithOption(&ts, sqlplugins.Desc("age"))
		if err != nil {
			t.Error(err)
			return
		}
		if len(ts) != count {
			t.Errorf("Expect %d, but got %d", count, len(ts))
		}
		for i := 0; i < count-1; i++ {
			if ts[i].Age < ts[i+1].Age {
				t.Errorf("Expect %d, but got %d", ts[i].Age, ts[i+1].Age)
			}
		}
	})
}

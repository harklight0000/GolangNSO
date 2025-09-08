package database

import (
	"context"
	"github.com/rotisserie/eris"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"nso/ainterfaces"
	"nso/errs"
	"nso/logging"
	"nso/sqlplugins"
	"testing"
)

func initMongoTest() (ainterfaces.IDatabase, *mongo.Database, error) {
	db := InitMongoDB()
	if db == nil {
		return nil, nil, eris.New("Cannot init mongo")
	}
	database := NewMongoDatabaseAdapter(db)
	if database == nil {
		return nil, nil, eris.New("Cannot init mongo database adapter")
	}
	return database, db, nil
}

func tearDownMongoTest(db *mongo.Database, adapterDb ainterfaces.IDatabase) {
	defer func() {
		if err := adapterDb.Close(); err != nil {
			logging.Logger.Error("Cannot drop mongo database", zap.Error(err))
		}
	}()
	ctx := context.TODO()
	err := db.Drop(ctx)
	if err != nil {
		logging.Logger.Panic("Fail to drop mongo database", zap.String("Drop mongodb error", errs.ToString(err)))
	}
}

func TestNewMongoDatabaseAdapter(t *testing.T) {
	database, db, err := initMongoTest()
	if err != nil {
		t.Error(err)
	}

	defer tearDownMongoTest(db, database)
	if database == nil {
		t.Error("Cannot init mongo database adapter")
	}
	var mongdoDB *mongo.Database
	database.Unwrap(&mongdoDB)
	if mongdoDB == nil {
		t.Error("Cannot unwrap mongo database adapter")
	}
	t.Run("Create", func(t *testing.T) {
		type Test struct {
			ObjectID string `bson:"_id" json:"_id"`
			Name     string `bson:"name" json:"name"`
		}
		t1 := Test{Name: "test"}
		err := database.Create(&t1)
		if err != nil {
			t.Error(err)
		}
		if t1.ObjectID == "" {
			t.Error("id should set")
		}
	})

	t.Run("FindOne", func(t *testing.T) {
		type Test struct {
			ObjectID string `bson:"_id" json:"_id"`
			Name     string `bson:"name" json:"name"`
		}
		t1 := Test{Name: "test"}
		err := database.Create(&t1)
		if err != nil {
			t.Error(err)
		}
		t1.Name = "test2"
		err = database.Update("test", t1, sqlplugins.Eq("_id", t1.ObjectID))
		if err != nil {
			t.Error(err)
		}
		var t2 Test
		err = database.FindOne(&t2, sqlplugins.Eq("_id", t1.ObjectID))
		if err != nil {
			t.Error(errs.ToString(err))
		}
		if t2.ObjectID == "" {
			t.Error("Cannot find one")
		}
	})

	t.Run("Update", func(t *testing.T) {
		type Test struct {
			ObjectID string `bson:"_id" json:"_id"`
			Name     string `bson:"name" json:"name"`
		}
		t1 := Test{Name: "test"}
		err := database.Create(&t1)
		if err != nil {
			t.Error(err)
		}
		t1.Name = "test2"
		err = database.Update("test", t1, sqlplugins.Eq("_id", t1.ObjectID))
		if err != nil {
			t.Error(err)
		}
		err = database.Update("test", &t1, sqlplugins.Eq("_id", t1.ObjectID))
		if err != nil {
			t.Error(eris.Wrap(err, "Update using pointer"))
		}
		var t2 Test
		err = database.FindOne(&t2, sqlplugins.Eq("_id", t1.ObjectID))
		if err != nil {
			t.Error(eris.Wrap(err, "Error FindOne"))
		}
		if t2.Name != "test2" {
			t.Error("update failed")
		}
	})

}

func TestMongoDBAdapter_FindMany(t *testing.T) {
	database, db, err := initMongoTest()
	if err != nil {
		t.Error(err)
	}

	defer tearDownMongoTest(db, database)
	if database == nil {
		t.Error("Cannot init mongo database adapter")
	}
	t1, t2 := &Test{Name: "test"}, &Test{Name: "test"}
	err = database.Create(t1)
	if err != nil {
		t.Error("Cannot create t1")
	}
	err = database.Create(t2)
	if err != nil {
		t.Error("Cannot create t2")
	}

	var t3 []*Test
	err = database.FindMany(&t3, sqlplugins.Eq("name", "test"))
	if err != nil {
		t.Error(errs.ToString(err))
	}
	if len(t3) != 2 {
		t.Error("FindManyWithConditionAndOption failed expect t3 size 2 but got ", len(t3))
	}
	var t4 []*Test
	err = database.FindMany(&t4, sqlplugins.Eq("name", "test2"))
	if err != nil {
		t.Error(errs.ToString(err))
	}
}

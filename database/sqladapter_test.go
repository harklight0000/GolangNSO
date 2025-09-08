package database

import (
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
	"nso/ainterfaces"
	config2 "nso/config"
	"nso/logging"
	. "nso/sqlplugins"
	"testing"
)

func init() {
	config2.GetAppConfig().UseSQL = true
}

type Test struct {
	ID   int    `gorm:"primary_key""`
	_ID  string `gorm:"-" bson:"_id"`
	Name string `gorm:"column:name" bson:"name"`
	gorm.Model
}

func (t *Test) TableName() string {
	return "test"
}

func initSQLTest() (*gorm.DB, ainterfaces.IDatabase, error) {
	db := InitSQLDB()
	if db == nil {
		return nil, nil, eris.New("Cannot init sql database")
	}
	database := NewSQLDatabaseAdapter(db)
	if database == nil {
		return nil, nil, eris.New("Cannot init sql database adapter")
	}
	err := db.Migrator().AutoMigrate(&Test{})
	if err != nil {
		return nil, nil, eris.Wrap(err, "Cannot migrate sql database")
	}
	return db, database, nil
}

func tearingSQLDown() {
	db := InitSQLDB()
	if db == nil {
		logging.Logger.Panic("Cannot init sql database")
	}
	err := db.Migrator().DropTable(&Test{})
	if err != nil {
		logging.Logger.Info("Cannot drop table")
		return
	}
}

func TestNewSQLDatabaseAdapter(t *testing.T) {
	defer tearingSQLDown()
	_, database, err := initSQLTest()
	if err != nil {
		t.Error(eris.Wrap(err, "Cannot init sql database"))
	}
	var db2 *gorm.DB
	database.Unwrap(&db2)
	if db2 == nil {
		t.Error("Cannot unwrap sql database")
	}
	if err != nil {
		t.Error("Cannot migrate sql database")
		return
	}
}

func TestCreate(t *testing.T) {
	defer tearingSQLDown()
	_, database, err := initSQLTest()
	if err != nil {
		t.Error(eris.Wrap(err, "Cannot init sql database"))
	}
	test := &Test{Name: "test"}
	err = database.Create(test)
	if err != nil {
		t.Error(eris.Wrap(err, "Cannot create test"))
		return
	}
	if test.ID == 0 {
		t.Error("Cannot create test")
	}
}

func TestUpdate(t *testing.T) {
	defer tearingSQLDown()
	_, database, err := initSQLTest()
	if err != nil {
		t.Error(eris.Wrap(err, "Cannot init sql database"))
	}
	test := &Test{Name: "test"}
	err = database.Create(test)
	if err != nil {
		t.Error(eris.Wrap(err, "Cannot create test"))
		return
	}
	test.Name = "test2"
	err = database.Update(test.TableName(), test, Eq("id", test.ID))
	if err != nil {
		t.Error(eris.Wrap(err, "Cannot update test"))
		return
	}
	var test2 Test
	err = database.FindOne(&test2, Eq("id", test.ID))
	if err != nil {
		t.Error(eris.Wrap(err, "Cannot find test"))
		return
	}
	if test2.Name != "test2" {
		t.Error("Cannot update test")
	}
}

func TestFindMany(t *testing.T) {
	defer tearingSQLDown()
	_, database, err := initSQLTest()
	if err != nil {
		t.Error(eris.Wrap(err, "Cannot init sql database"))
	}
	tests := []*Test{
		{Name: "test1"},
		{Name: "test2"},
	}
	err = database.Create(tests)
	if err != nil {
		t.Error(eris.Wrap(err, "Cannot create test"))
		return
	}
	var tests1 []Test
	err = database.FindMany(&tests1, Like("name", "%test%"))
	if err != nil {
		t.Error(eris.Wrap(err, "Cannot find test"))
		return
	}
	if len(tests1) != 2 {
		t.Error("Cannot find test")
	}
}

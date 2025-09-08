package database

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	config2 "nso/config"
	. "nso/logging"
)

func InitSQLDB() *gorm.DB {
	var db *gorm.DB
	var err error
	config := config2.GetAppConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.DbName)
	Logger.Info("SQL", zap.String("SQL", "Connecting to database:"+dsn))
	var _logger logger.Interface
	if config.LogSql {
		_logger = logger.Default.LogMode(logger.Info)
	}
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   _logger, // -> Log các câu lệnh truy vấn database trong terminal
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		Logger.Info("SQL", zap.Error(err))
	}
	schema.RegisterSerializer("my_json", &JSONSerializer{})

	if config.DropDb {
		exec := db.Exec("Drop database if exists " + config.DbName + ";")
		if exec.Error != nil {
			Logger.Info("SQL", zap.Error(exec.Error))
			return nil
		}
		tx := db.Exec("Create database " + config.DbName + ";")
		if tx.Error != nil {
			Logger.Info("SQL", zap.Error(tx.Error))
			return nil
		}
		db.Exec("USE " + config.DbName + ";")
	}
	// Register serializer
	schema.RegisterSerializer("my_json", &JSONSerializer{})
	if config.Migrate {
		Migrate(db)
	}

	if err != nil {
		Logger.Info("SQL", zap.Error(err))
	}

	scan := db.Exec("SET SQL_MODE=NO_AUTO_VALUE_ON_ZERO;")

	if scan.Error != nil {
		Logger.Info("SQL", zap.Error(scan.Error))
	}

	if config.InitDataTest {
		_ = db.Exec("SET foreign_key_checks = 0")
		Logger.Info("InitDataTest success")
		_ = db.Exec("SET foreign_key_checks = 1")
	}
	return db
}

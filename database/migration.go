package database

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nso/entity"
	"nso/logging"
)

func Migrate(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(
		&entity.MapEntity{},
		&entity.NPCEntity{},
		&entity.MobEntity{},
		&entity.OptionItemEntity{},
		&entity.ItemEntity{})
	if err != nil {
		logging.Logger.Info("Error migrating maps", zap.Error(err))
	}
}

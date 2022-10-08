package database

import (
	"aliyunoss/app/model"
	"aliyunoss/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dbName string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		logger.Error("database", zap.String("init", err.Error()),
			zap.String("dbName", dbName))
	}
	DB.AutoMigrate(&model.User{})
}

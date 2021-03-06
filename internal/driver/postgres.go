package driver

import (
	"fmt"
	"github.com/A-ndrey/tododo/internal/config"
	"github.com/A-ndrey/tododo/internal/domains"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const schemaName = "tododo"

var loggerMods = map[config.Environment]logger.LogLevel{
	config.ENV_DEV:  logger.Info,
	config.ENV_PROD: logger.Silent,
}

var entities = []interface{}{
	&domains.Task{},
	&domains.User{},
}

func NewPostgresGorm() (*gorm.DB, error) {
	conf := config.GetPostgres()

	gcfg := gorm.Config{
		Logger:         logger.Default.LogMode(loggerMods[config.GetEnvironment()]),
		NamingStrategy: schema.NamingStrategy{TablePrefix: schemaName + "."},
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s", conf.User, conf.Password, conf.DBName, conf.Host)
	db, err := gorm.Open(postgres.Open(dsn), &gcfg)
	if err != nil {
		return nil, err
	}

	result := db.Exec("create schema if not exists " + schemaName)
	if result.Error != nil {
		return nil, result.Error
	}

	if err := db.AutoMigrate(entities...); err != nil {
		return nil, err
	}

	return db, nil
}

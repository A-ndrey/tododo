package driver

import (
	"fmt"
	"github.com/A-ndrey/tododo/internal/config"
	"github.com/A-ndrey/tododo/internal/task"
	"github.com/A-ndrey/tododo/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const schemaName = "tododo"

var entities = []interface{}{
	&task.Task{},
	&user.User{},
}

func NewPostgresGorm() (*gorm.DB, error) {
	conf := config.GetPostgres()

	gcfg := gorm.Config{
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

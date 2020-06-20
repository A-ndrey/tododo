package postgres

import (
	"fmt"
	"github.com/A-ndrey/tododo/internal/config"
	"github.com/A-ndrey/tododo/internal/task"
	"github.com/A-ndrey/tododo/internal/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var entities = []interface{}{
	&task.Task{},
	&user.User{},
}

func Connect() (*gorm.DB, error) {
	conf := config.GetPostgres()
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", conf.User, conf.Password, conf.DBName, conf.SSLMode)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if config.GetEnvironment() == config.ENV_DEV {
		db.DropTableIfExists(entities...)
	}

	db.AutoMigrate(entities...)

	return db, nil
}

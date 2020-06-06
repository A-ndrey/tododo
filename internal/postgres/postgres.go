package postgres

import (
	"database/sql"
	"fmt"
	"github.com/A-ndrey/tododo/internal/config"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	conf := config.GetPostgres()
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", conf.User, conf.Password, conf.DBName, conf.SSLMode)
	return sql.Open("postgres", connStr)
}

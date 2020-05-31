package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "user=tduser password=topadossdo dbname=tododo sslmode=disable"
	return sql.Open("postgres", connStr)
}

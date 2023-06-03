package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	User     string
	Pass     string
	Host     string
	Port     string
	Database string
}

func NewPostgres(c DBConfig) (*sql.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Pass, c.Host, c.Port, c.Database)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}

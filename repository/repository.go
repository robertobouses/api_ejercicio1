package repository

import (
	"database/sql"
	_ "embed"
)

//go:embed sql/area.sql
var insertAreaQuery string

type repository struct {
	db         *sql.DB
	insertArea *sql.Stmt
}

func NewRepo(db *sql.DB) (*repository, error) {
	insertStmt, err := db.Prepare(insertAreaQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:         db,
		insertArea: insertStmt,
	}, nil

}
func (r *repository) InsertArea(area int) error {
	_, err := r.insertArea.Exec(area)
	return err
}

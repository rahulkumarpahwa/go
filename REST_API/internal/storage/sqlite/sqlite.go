package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Connection(storagePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

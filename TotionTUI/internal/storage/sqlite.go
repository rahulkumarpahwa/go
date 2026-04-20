package storage

import (
	"database/sql"
	"errors"
	"time"
)

type NotesStorage struct {
}

func Open(StoragePath string) (*sql.DB, error) {
	if StoragePath == "" {
		return nil, errors.New("No StoragePath Provided!")
	}
	
	DB, err := sql.Open("", StoragePath)
	if err != nil {
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	DB.SetConnMaxIdleTime(time.Second * 20)

	return DB, nil
}

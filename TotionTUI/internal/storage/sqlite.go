package storage

import (
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"github.com/rahulkumarpahwa/go/TotionTUI/config"
	_ "modernc.org/sqlite"
)

type NotesStorage struct {
	DB *sql.DB
}

func Open(cfg *config.Config) (*sql.DB, error) {
	if cfg.StoragePath == "" {
		return nil, errors.New("No StoragePath Provided!")
	}

	DB, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	DB.SetConnMaxIdleTime(time.Second * 20)

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS notes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    description TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`)

	slog.Info("Created Table Notes Successfully!")

	if err != nil {
		return nil, err
	}
	return DB, nil
}

func (ns *NotesStorage) CreateNote(title string, description string) (*int64, error) {
	statement, err := ns.DB.Prepare("INSERT INTO notes (title, description) VALUES ($1, $2)")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	result, err := statement.Exec(title, description)
	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &lastId, nil
}

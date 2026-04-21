package storage

import (
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"github.com/rahulkumarpahwa/go/TotionTUI/config"
	"github.com/rahulkumarpahwa/go/TotionTUI/internal/types"
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

func (ns *NotesStorage) CreateNote(title, description string) (*int64, error) {
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

func (ns *NotesStorage) GetNotes(pageNo int, pageSize int) ([]types.Notes, error) {
	if pageSize == 0 {
		pageSize = 20 // default limit
	}
	if pageNo < 1 {
		pageNo = 1
	}
	offset := (pageNo - 1) * pageSize

	statement, err := ns.DB.Prepare("Select id, title, description, created_at, updated_at FROM notes LIMIT $1 OFFSET $2")

	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(pageSize, offset)
	if err != nil {
		return nil, err
	}

	var notes []types.Notes
	for rows.Next() {
		var note types.Notes
		err := rows.Scan(&note.Id, &note.Title, &note.Description, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

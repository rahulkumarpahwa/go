package storage

import (
	"database/sql"
	"log"

	"github.com/rahulkumarpahwa/go/REST_API/internal/storage/sqlite"
)

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
}

type StudentStorage struct {
	db     *sql.DB
	logger *log.Logger
}

func (ss *StudentStorage) CreateStudent(name string, email string, age int) (int64, error) {
	query := `INSERT INTO STUDENT (name, email, age) VALUES ($1, $2, $4) RETURNING id`;

	sqlite.Sqlite.DB
	

	return 0, nil
}

package sqlite

import (
	"database/sql"

	"github.com/rahulkumarpahwa/go/REST_API/internal/config"
	"github.com/rahulkumarpahwa/go/REST_API/internal/types"
	_ "modernc.org/sqlite"
)

type Sqlite struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	); `)

	if err != nil {
		return nil, err
	}

	return &Sqlite{DB: db}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	statement, err := s.DB.Prepare("INSERT INTO STUDENTS (name, email, age) VALUES ($1, $2, $3)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (*types.Student, error) {

	var student types.Student
	statement, err := s.DB.Prepare("SELECT name, email, age FROM students WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	student.Id = id
	err = statement.QueryRow(id).Scan(&student.Name, &student.Email, &student.Age)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

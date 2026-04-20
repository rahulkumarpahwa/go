package storage

import "github.com/rahulkumarpahwa/go/REST_API/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int) (*types.Student, error)
	GetStudentsList() ([]types.Student, error)
	UpdateStudent(id int, name string, age int) (*types.Student, error)
	DeleteStudent(id int) (*int64, error)
}

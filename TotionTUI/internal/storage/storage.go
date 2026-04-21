package storage

type Notes interface {
	CreateNote(title, description string) (*int64, error)
}

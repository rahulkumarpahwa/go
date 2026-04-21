package storage

type Notes interface {
	CreateNote(string, string) (*int64, error)
}

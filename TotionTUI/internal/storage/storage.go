package storage

import "github.com/rahulkumarpahwa/go/TotionTUI/internal/types"

type Notes interface {
	CreateNote(title, description string) (*int64, error)
	GetNotes(pageNo int64, pageSize int64) ([]types.Notes, error)
	DeleteNoteById(id int) error
	GetNoteById(id int) (*types.Notes, error)
}

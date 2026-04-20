package storage

import "github.com/rahulkumarpahwa/go/TotionTUI/internal/types"

type NotesStorage interface {
	CreateNotes(note *types.Notes) (*types.Notes, error)
}
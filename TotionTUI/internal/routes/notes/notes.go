package notes

import (
	"log"

	"github.com/rahulkumarpahwa/go/TotionTUI/internal/storage"
)

type NotesHandler struct {
	Storage storage.NotesStorage
	Logger *log.Logger
}




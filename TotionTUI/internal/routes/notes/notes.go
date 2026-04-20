package notes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rahulkumarpahwa/go/TotionTUI/internal/storage"
	"github.com/rahulkumarpahwa/go/TotionTUI/internal/types"
)

type NotesHandler struct {
	Storage storage.Notes
	Logger  *log.Logger
}

func (h *NotesHandler) CreateNotes(w http.ResponseWriter, r *http.Request) {
	var note types.Notes

	json.NewDecoder(r.Body).Decode(&note)


}

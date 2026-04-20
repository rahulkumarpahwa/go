package notes

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"

	"github.com/rahulkumarpahwa/go/TotionTUI/internal/storage"
	"github.com/rahulkumarpahwa/go/TotionTUI/internal/types"
	"github.com/rahulkumarpahwa/go/TotionTUI/internal/utils"
)

type NotesHandler struct {
	Storage storage.Notes
	Logger  *log.Logger
}

func (h *NotesHandler) CreateNotes(w http.ResponseWriter, r *http.Request) {
	var note types.NotesRequest

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteJson(w, http.StatusBadRequest, utils.Utils{Message: "Can't Decode the Note"})
		return
	}

	

}

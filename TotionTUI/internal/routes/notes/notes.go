package notes

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/rahulkumarpahwa/go/TotionTUI/internal/storage"
	"github.com/rahulkumarpahwa/go/TotionTUI/internal/types"
	"github.com/rahulkumarpahwa/go/TotionTUI/internal/utils"
)

type NotesHandler struct {
	Storage storage.Notes
	Logger  *log.Logger
}

func (h *NotesHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note types.NotesRequest

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteJson(w, http.StatusBadRequest, utils.Utils{Message: "Can't Decode the Note"})
		return
	}

	lastId, err := h.Storage.CreateNote(note.Title, note.Description)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteJson(w, http.StatusInternalServerError, utils.Utils{Message: "Can't Storage the Note"})
		return
	}
	utils.WriteJson(w, http.StatusCreated, utils.Utils{Message: "created note successfully!", Data: map[string]any{"Created Note Id": lastId}})
}

func (h *NotesHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var pageSize int64 = 0
	var pageNo int64 = 0

	pageNoStr := query.Get("pageNo")
	pageSizeStr := query.Get("pageSize")

	// todo : Write the single method to parse query params under utils

	if pageSizeStr != "" {
		var err error
		pageSize, err = strconv.ParseInt(pageSizeStr, 10, 64)
		if err != nil {
			slog.Error(err.Error())
			utils.WriteJson(w, http.StatusBadRequest, utils.Utils{Message: "Can't Decode the PageSize"})
			return
		}
	}

	if pageNoStr != "" {
		var err error
		pageNo, err = strconv.ParseInt(pageNoStr, 10, 64)
		if err != nil {
			slog.Error(err.Error())
			utils.WriteJson(w, http.StatusBadRequest, utils.Utils{Message: "Can't Decode the PageNo"})
			return
		}
	}

	notes, err := h.Storage.GetNotes(pageNo, pageSize)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteJson(w, http.StatusBadRequest, utils.Utils{Message: "Can't Get the Notes!"})
		return
	}

	utils.WriteJson(w, http.StatusOK, utils.Utils{Message: "Got Notes Successfully!", Data: map[string]any{"Notes": notes}})
}

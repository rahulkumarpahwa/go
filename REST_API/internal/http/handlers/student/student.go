package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/rahulkumarpahwa/go/REST_API/internal/storage"
	"github.com/rahulkumarpahwa/go/REST_API/internal/types"
	"github.com/rahulkumarpahwa/go/REST_API/internal/utils/response"
)

type StudentHandler struct {
	Storage storage.Storage
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student types.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if errors.Is(err, io.EOF) {
		slog.Error("Can't decode ")
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("Empty Body!")))
		return
	}

	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	// request Validation:
	err = validator.New().Struct(student)
	if err != nil {
		validateErrors := err.(validator.ValidationErrors) // typecasting the error into the slice
		response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrors))
		return
	}

	lastId, err := h.Storage.CreateStudent(student.Name, student.Email, student.Age)
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}
	slog.Info("User Created Successfully!", slog.String("userID", fmt.Sprint(lastId)))

	err = response.WriteJson(w, http.StatusCreated, map[string]any{"id": lastId})
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

}

func (h *StudentHandler) GetStudentById(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")

	if idstr == "" {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("Id not Found!")))
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	student, err := h.Storage.GetStudentById(id)
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}
	slog.Info("Student Founded Successfully!", slog.String("student", student.Name))

	err = response.WriteJson(w, http.StatusOK, map[string]any{"Student": student})
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

}

func (h *StudentHandler) GetStudentsList(w http.ResponseWriter, r *http.Request) {

	students, err := h.Storage.GetStudentsList()
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}
	slog.Info("Students List Founded Successfully!")

	err = response.WriteJson(w, http.StatusOK, map[string]any{"Students List": students})
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

}

func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")

	if idstr == "" {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("Id not Found!")))
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	student, err := h.Storage.UpdateStudent(id, name, age)
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}
	slog.Info("Student Founded Successfully!", slog.String("student", student.Name))

	err = response.WriteJson(w, http.StatusOK, map[string]any{"Student": student})
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

}

func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")

	if idstr == "" {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("Id not Found!")))
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	student, err := h.Storage.GetStudentById(id)
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}
	slog.Info("Student Founded Successfully!", slog.String("student", student.Name))

	err = response.WriteJson(w, http.StatusOK, map[string]any{"Student": student})
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

}

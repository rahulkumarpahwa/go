package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/rahulkumarpahwa/go/REST_API/internal/types"
	"github.com/rahulkumarpahwa/go/REST_API/internal/utils/response"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
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

	err = response.WriteJson(w, http.StatusCreated, response.Response{Status: response.StatusOK})
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

}

package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

type Envelope map[string]any

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	Envelope
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errors validator.ValidationErrors) Response {
	var errMaps[]string

	for _, err := range errors {
		switch err.ActualTag() {
		case "required" : 
			errMaps = append(errMaps, fmt.Sprintf("field %s is required field", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

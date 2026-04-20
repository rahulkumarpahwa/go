package utils

import (
	"encoding/json"
	"net/http"
)

type Utils struct {
	Message string `json:"message"`
	Data    map[string]any `json:"data"`
}

func WriteJson(w http.ResponseWriter, status int, data Utils) error {
	js, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

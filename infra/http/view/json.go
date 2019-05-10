package view

import (
	"encoding/json"
	"net/http"
)

func NewJSON() *JSON {
	return new(JSON)
}

type JSON struct{}

func (j *JSON) Show(w http.ResponseWriter, name string, data interface{}) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

	return nil
}

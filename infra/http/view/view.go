package view

import "net/http"

type View interface {
	Show(w http.ResponseWriter, name string, data interface{}) error
}

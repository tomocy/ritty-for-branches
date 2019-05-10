package handler

import (
	"log"
	"net/http"

	"github.com/tomocy/ritty-for-branches/infra/http/view"
)

func New(view view.View) *Handler {
	return &Handler{
		branchHandler: newBranchHandler(view),
	}
}

type Handler struct {
	*branchHandler
}

func logInternalServerError(w http.ResponseWriter, did string, msg interface{}) {
	log.Printf("failed to %s: %v\n", did, msg)
	w.WriteHeader(http.StatusInternalServerError)
}

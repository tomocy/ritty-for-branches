package handler

import (
	"log"
	"net/http"

	"github.com/tomocy/ritty-for-branches/domain/repository"
	"github.com/tomocy/ritty-for-branches/infra/http/view"
)

func New(
	view view.View,
	branchRepo repository.BranchRepository,
) *Handler {
	return &Handler{
		branchHandler: newBranchHandler(view, branchRepo),
	}
}

type Handler struct {
	*branchHandler
}

func logInternalServerError(w http.ResponseWriter, did string, msg interface{}) {
	log.Printf("failed to %s: %v\n", did, msg)
	w.WriteHeader(http.StatusInternalServerError)
}

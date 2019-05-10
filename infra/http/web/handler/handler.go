package handler

import (
	"log"
	"net/http"

	"github.com/tomocy/ritty-for-branches/domain/repository"
	"github.com/tomocy/ritty-for-branches/domain/service"
	"github.com/tomocy/ritty-for-branches/infra/http/view"
)

func New(
	view view.View,
	branchRepo repository.BranchRepository,
	authorizationServ service.AuthorizationService,
) *Handler {
	return &Handler{
		authorizationHandler: newAuthorizationHandler(authorizationServ),
		branchHandler:        newBranchHandler(view, branchRepo),
	}
}

type Handler struct {
	*authorizationHandler
	*branchHandler
}

func redirect(w http.ResponseWriter, r *http.Request, dest string) {
	http.Redirect(w, r, dest, http.StatusSeeOther)
}

func logInternalServerError(w http.ResponseWriter, did string, msg interface{}) {
	log.Printf("failed to %s: %v\n", did, msg)
	w.WriteHeader(http.StatusInternalServerError)
}

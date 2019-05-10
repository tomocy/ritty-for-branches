package handler

import (
	"net/http"

	"github.com/tomocy/ritty-for-branches/domain/service"
)

func newAuthorizationHandler(
	serv service.AuthorizationService,
) *authorizationHandler {
	return &authorizationHandler{
		serv: serv,
	}
}

type authorizationHandler struct {
	serv service.AuthorizationService
}

func (h *authorizationHandler) FetchAuthorizationCode(w http.ResponseWriter, r *http.Request) {
	dest := h.serv.BuildAuthorizationCodeURL()
	redirect(w, r, dest.String())
}

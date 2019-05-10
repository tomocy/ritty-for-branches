package handler

import (
	"net/http"

	"github.com/tomocy/ritty-for-branches/domain/model"
	"github.com/tomocy/ritty-for-branches/domain/repository"
	"github.com/tomocy/ritty-for-branches/domain/service"
	"github.com/tomocy/ritty-for-branches/infra/http/session"
)

func newAuthorizationHandler(
	repo repository.BranchRepository,
	serv service.AuthorizationService,
) *authorizationHandler {
	return &authorizationHandler{
		repo: repo,
		serv: serv,
	}
}

type authorizationHandler struct {
	repo repository.BranchRepository
	serv service.AuthorizationService
}

func (h *authorizationHandler) FetchAuthorizationCode(w http.ResponseWriter, r *http.Request) {
	dest := h.serv.BuildAuthorizationCodeURL()
	redirect(w, r, dest.String())
}

func (h *authorizationHandler) FetchAuthorization(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	auth, err := h.serv.FetchAuthorization(code, r.Cookies()...)
	if err != nil {
		logInternalServerError(w, "fetch authorization", err)
		return
	}

	var branch *model.Branch
	if branch, err = h.repo.FindBranch(auth.BranchID); err != nil {
		branch = &model.Branch{
			ID: auth.BranchID,
		}
	}

	branch.Authorization = auth
	if err := h.repo.SaveBranch(branch); err != nil {
		logInternalServerError(w, "fetch authorization", err)
		return
	}

	session.KeepBranchAuthenticated(w, r, branch.ID)

	dest := "/"
	if intended, err := session.FindIntendedURL(r); err == nil {
		dest = intended
		session.RemoveIntendedURL(w, r)
	}

	redirect(w, r, dest)
}

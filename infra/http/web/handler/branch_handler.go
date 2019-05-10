package handler

import (
	"net/http"

	"github.com/tomocy/ritty-for-branches/domain/repository"
	"github.com/tomocy/ritty-for-branches/infra/http/session"
	"github.com/tomocy/ritty-for-branches/infra/http/view"
)

func newBranchHandler(
	view view.View,
	repo repository.BranchRepository,
) *branchHandler {
	return &branchHandler{
		view: view,
		repo: repo,
	}
}

type branchHandler struct {
	view view.View
	repo repository.BranchRepository
}

func (h *branchHandler) ShowMenus(w http.ResponseWriter, r *http.Request) {
	id, err := session.FindAuthenticBranch(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	branch, err := h.repo.FindBranch(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := h.view.Show(w, "menu.index", map[string]interface{}{
		"Branch": branch,
	}); err != nil {
		logInternalServerError(w, "show restaurant menus", err)
	}
}

func (h *branchHandler) ShowMenuRegistrationForm(w http.ResponseWriter, r *http.Request) {
	if err := h.view.Show(w, "menu.new", nil); err != nil {
		logInternalServerError(w, "show branch menu registration form", err)
	}
}

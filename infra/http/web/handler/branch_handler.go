package handler

import (
	"net/http"

	"github.com/tomocy/ritty-for-branches/infra/http/view"
)

func newBranchHandler(view view.View) *branchHandler {
	return &branchHandler{
		view: view,
	}
}

type branchHandler struct {
	view view.View
}

func (h *branchHandler) ShowMenuRegistrationForm(w http.ResponseWriter, r *http.Request) {
	if err := h.view.Show(w, "menu.new", nil); err != nil {
		logInternalServerError(w, "show branch menu registration form", err)
	}
}

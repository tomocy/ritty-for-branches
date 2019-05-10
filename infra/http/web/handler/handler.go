package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/thedevsaddam/govalidator"
	derr "github.com/tomocy/ritty-for-branches/domain/error"
	"github.com/tomocy/ritty-for-branches/domain/repository"
	"github.com/tomocy/ritty-for-branches/domain/service"
	"github.com/tomocy/ritty-for-branches/infra/http/view"
)

func New(
	view view.View,
	branchRepo repository.BranchRepository,
	authorizationServ service.AuthorizationService,
	storageServ service.StorageService,
) *Handler {
	return &Handler{
		authorizationHandler: newAuthorizationHandler(branchRepo, authorizationServ),
		branchHandler:        newBranchHandler(view, branchRepo, storageServ),
	}
}

type Handler struct {
	*authorizationHandler
	*branchHandler
}

func validate(names []string, options govalidator.Options) error {
	valid := govalidator.New(options)
	msgsMap := valid.Validate()
	if 1 <= len(msgsMap) {
		for _, name := range names {
			if msgs := msgsMap[name]; 1 <= len(msgs) {
				return errors.New(msgs[0])
			}
		}
	}

	return nil
}

func validationErrorf(did string, msg interface{}) *derr.ValidationError {
	return derr.ValidationErrorf("failed to %s: %v", did, msg)
}

func devErrorf(did string, msg interface{}) *derr.DevError {
	return derr.DevErrorf("failed to %s: %v", did, msg)
}

func redirect(w http.ResponseWriter, r *http.Request, dest string) {
	http.Redirect(w, r, dest, http.StatusSeeOther)
}

func logInternalServerError(w http.ResponseWriter, did string, msg interface{}) {
	log.Printf("failed to %s: %v\n", did, msg)
	w.WriteHeader(http.StatusInternalServerError)
}

package session

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/tomocy/sensei"
)

func KeepBranchAuthenticated(w http.ResponseWriter, r *http.Request, id string) {
	if err := manager.Set(w, r, authenticBranchID, id); err != nil {
		logError("keep branch authenticated", err)
	}
}

func IsBranchAuthenticated(r *http.Request) bool {
	_, err := FindAuthenticBranch(r)
	return err == nil
}

func FindAuthenticBranch(r *http.Request) (string, error) {
	i, err := manager.Get(r, authenticBranchID)
	if err != nil {
		return "", err
	}
	if s, ok := i.(string); ok && s != "" {
		return s, nil
	}

	return "", errors.New("no authentic branch")
}

func KeepIntendedURL(w http.ResponseWriter, r *http.Request, url string) {
	if err := manager.Set(w, r, intendedURL, url); err != nil {
		logError("keep intended url", err)
	}
}

func FindIntendedURL(r *http.Request) (string, error) {
	i, err := manager.Get(r, intendedURL)
	if err != nil {
		return "", err
	}
	if s, ok := i.(string); ok && s != "" {
		return s, nil
	}

	return "", errors.New("no intended url")
}

func RemoveIntendedURL(w http.ResponseWriter, r *http.Request) {
	if err := manager.Remove(w, r, intendedURL); err != nil {
		logError("remove intended url", err)
	}
}

func IsSessionValid(r *http.Request) bool {
	_, err := manager.Session(r)
	return err == nil
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	sess, _ := manager.Session(r)
	sess.Options.MaxAge = -1
	if err := sess.Save(r, w); err != nil {
		logError("delete session", err)
	}
}

var manager = sensei.New(store, key)

var store = sessions.NewCookieStore(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

const (
	key = "RittyForBranches"

	authenticBranchID = "authentic_branch_id"

	intendedURL = "intended_url"
)

func logError(did string, msg interface{}) {
	log.Printf("failed to %s: %v\n", did, msg)
}

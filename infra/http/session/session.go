package session

import (
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/tomocy/sensei"
)

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
)

func logError(did string, msg interface{}) {
	log.Printf("failed to %s: %v\n", did, msg)
}

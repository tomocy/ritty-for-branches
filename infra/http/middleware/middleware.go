package middleware

import (
	"net/http"

	"github.com/tomocy/ritty-for-branches/infra/http/session"
)

func RenewInvalidSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !session.IsSessionValid(r) {
			session.DeleteSession(w, r)
			http.Redirect(w, r, r.URL.RequestURI(), http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}

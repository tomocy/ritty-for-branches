package middleware

import (
	"net/http"

	"github.com/tomocy/ritty-for-branches/infra/http/route"
	"github.com/tomocy/ritty-for-branches/infra/http/session"
)

func AcceptAuthenticBranch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !session.IsBranchAuthenticated(r) {
			session.KeepIntendedURL(w, r, r.URL.RequestURI())
			http.Redirect(w, r, route.Web.Route("authorization_code.new").String(), http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func DenyAuthenticBranch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session.IsBranchAuthenticated(r) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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

package registerer

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tomocy/ritty-for-branches/infra/http/middleware"
	"github.com/tomocy/ritty-for-branches/infra/http/route"
	"github.com/tomocy/ritty-for-branches/infra/http/web/handler"
)

func NewWebRegisterer(handler *handler.Handler) *WebRegisterer {
	return &WebRegisterer{
		handler: handler,
	}
}

type WebRegisterer struct {
	handler *handler.Handler
}

func (r *WebRegisterer) RegisterRoutes(router chi.Router) {
	router.Use(middleware.RenewInvalidSession)
	router.Get("/*", http.FileServer(http.Dir("resource/public")).ServeHTTP)
	router.Get(route.Web.Route("menu.new").Path, r.handler.ShowMenuRegistrationForm)
}

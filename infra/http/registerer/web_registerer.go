package registerer

import (
	"net/http"
	"path"

	"github.com/tomocy/chme"

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
	router.Use(
		middleware.RenewInvalidSession,
		chme.ChangePostToHiddenMethod,
	)
	router.Get("/*", http.FileServer(http.Dir("resource/public")).ServeHTTP)

	router.Group(func(router chi.Router) {
		router.Use(middleware.AcceptAuthenticBranch)
		router.Get(route.Web.Route("menu.index").Path, r.handler.ShowMenus)
		router.Get(route.Web.Route("menu.new").Path, r.handler.ShowMenuRegistrationForm)
		router.Post(route.Web.Route("menu.create").Path, r.handler.RegisterMenu)
		router.Get(path.Join(route.Web.Route("menu.show").Path, "{category_name}", "{name}"), r.handler.ShowMenu)
		router.Put(path.Join(route.Web.Route("menu.update").Path, "{category_name}", "{name}"), r.handler.UpdateMenu)
		router.Delete(path.Join(route.Web.Route("menu.delete").Path, "{category_name}", "{name}"), r.handler.DeleteMenu)
	})

	router.Group(func(router chi.Router) {
		router.Use(middleware.DenyAuthenticBranch)
		router.Get(route.Web.Route("authorization_code.new").Path, r.handler.FetchAuthorizationCode)
		router.Get(route.Web.Route("authorization.new").Path, r.handler.FetchAuthorization)
	})
}

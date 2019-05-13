package registerer

import (
	"path"

	"github.com/go-chi/chi"
	"github.com/tomocy/ritty-for-branches/infra/http/api/handler"
	"github.com/tomocy/ritty-for-branches/infra/http/route"
)

func NewAPIRegisterer(handler *handler.Handler) *APIRegisterer {
	return &APIRegisterer{
		handler: handler,
	}
}

type APIRegisterer struct {
	handler *handler.Handler
}

func (r *APIRegisterer) RegisterRoutes(router chi.Router) {
	router.Group(func(router chi.Router) {
		router.Get(route.API.Route("branch.index").Path, r.handler.DisposeBranches)
		router.Get(path.Join(route.API.Route("branch.show").Path, "{id}"), r.handler.DisposeBranch)
	})
}

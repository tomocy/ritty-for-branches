package registerer

import (
	"github.com/go-chi/chi"
	"github.com/tomocy/ritty-for-branches/infra/http/api/handler"
)

func NewAPIRegisterer(handler *handler.Handler) *APIRegisterer {
	return &APIRegisterer{
		handler: handler,
	}
}

type APIRegisterer struct {
	handler *handler.Handler
}

func (r *APIReigsterer) RegisterRoutes(router chi.Router) {}

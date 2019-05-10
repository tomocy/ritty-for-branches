package registerer

import (
	"github.com/go-chi/chi"
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

func (r *WebRegisterer) RegisterRoutes(router chi.Router) {}

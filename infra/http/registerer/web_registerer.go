package registerer

import "github.com/go-chi/chi"

func NewWebRegisterer() *WebRegisterer {
	return new(WebRegisterer)
}

type WebRegisterer struct{}

func (r *WebRegisterer) RegisterRoutes(router chi.Router) {}

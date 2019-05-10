package registerer

import "github.com/go-chi/chi"

type Registerer interface {
	RegisterRoutes(router chi.Router)
}

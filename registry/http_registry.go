package registry

import (
	"github.com/tomocy/ritty-for-branches/infra/db"
	"github.com/tomocy/ritty-for-branches/infra/http/view"
	"github.com/tomocy/ritty-for-branches/infra/http/web/handler"
	"github.com/tomocy/ritty-for-branches/infra/ritty"
)

func NewHTTPRegistry() *HTTPRegistry {
	return &HTTPRegistry{
		web: newHTTPWebRegistry(),
	}
}

type HTTPRegistry struct {
	web *httpWebRegistry
}

func (r *HTTPRegistry) NewWebHandler() *handler.Handler {
	return r.web.newHandler()
}

func newHTTPWebRegistry() *httpWebRegistry {
	return &httpWebRegistry{
		db:    db.NewMemory(),
		ritty: ritty.New(),
		view:  view.NewHTML(),
	}
}

type httpWebRegistry struct {
	db    *db.Memory
	ritty *ritty.Ritty
	view  *view.HTML
}

func (r *httpWebRegistry) newHandler() *handler.Handler {
	return handler.New(r.view, r.db, r.ritty)
}

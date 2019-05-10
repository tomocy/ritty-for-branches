package registry

import (
	"github.com/tomocy/ritty-for-branches/infra/db"
	"github.com/tomocy/ritty-for-branches/infra/fs"
	api "github.com/tomocy/ritty-for-branches/infra/http/api/handler"
	"github.com/tomocy/ritty-for-branches/infra/http/view"
	web "github.com/tomocy/ritty-for-branches/infra/http/web/handler"
	"github.com/tomocy/ritty-for-branches/infra/ritty"
)

func NewHTTPRegistry() *HTTPRegistry {
	db := db.NewMemory()
	return &HTTPRegistry{
		api: newHTTPAPIRegistry(db),
		web: newHTTPWebRegistry(db),
	}
}

type HTTPRegistry struct {
	api *httpAPIRegistry
	web *httpWebRegistry
}

func (r *HTTPRegistry) NewAPIHandler() *api.Handler {
	return r.api.newHandler()
}

func (r *HTTPRegistry) NewWebHandler() *web.Handler {
	return r.web.newHandler()
}

func (r *httpWebRegistry) newHandler() *web.Handler {
	return web.New(r.view, r.db, r.ritty, r.fs)
}

func newHTTPAPIRegistry(db *db.Memory) *httpAPIRegistry {
	return &httpAPIRegistry{
		db:   db,
		view: view.NewJSON(),
	}
}

type httpAPIRegistry struct {
	db   *db.Memory
	view *view.JSON
}

func (r *httpAPIRegistry) newHandler() *api.Handler {
	return api.New(r.view, r.db)
}

func newHTTPWebRegistry(db *db.Memory) *httpWebRegistry {
	return &httpWebRegistry{
		db:    db,
		fs:    fs.NewLocal(),
		ritty: ritty.New(),
		view:  view.NewHTML(),
	}
}

type httpWebRegistry struct {
	db    *db.Memory
	fs    *fs.Local
	ritty *ritty.Ritty
	view  *view.HTML
}

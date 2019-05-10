package registry

import (
	"github.com/tomocy/ritty-for-branches/infra/http/view"
	"github.com/tomocy/ritty-for-branches/infra/http/web/handler"
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
		view: view.NewHTML(),
	}
}

type httpWebRegistry struct {
	view *view.HTML
}

func (r *httpWebRegistry) newHandler() *handler.Handler {
	return handler.New(r.view)
}

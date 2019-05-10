package registry

func NewHTTPRegistry() *HTTPRegistry {
	return new(HTTPRegistry)
}

type HTTPRegistry struct{}

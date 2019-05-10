package registerer

func NewWebRegisterer() *WebRegisterer {
	return new(WebRegisterer)
}

type WebRegisterer struct{}

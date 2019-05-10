package registerer

func NewAPIRegisterer() *APIRegisterer {
	return new(APIRegisterer)
}

type APIRegisterer struct{}

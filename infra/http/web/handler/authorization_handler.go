package handler

func newAuthorizationHandler() *authorizationHandler {
	return new(authorizationHandler)
}

type authorizationHandler struct{}

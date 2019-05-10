package service

import (
	"net/http"
	"net/url"

	"github.com/tomocy/ritty-for-branches/domain/model"
)

type AuthorizationService interface {
	BuildAuthorizationCodeURL() *url.URL
	FetchAuthorization(code string, cookies ...*http.Cookie) (*model.Authorization, error)
}

package service

import "net/url"

type AuthorizationService interface {
	BuildAuthorizationCodeURL() *url.URL
}

package ritty

import (
	"net/url"

	"github.com/tomocy/ritty-for-branches/config"
	"github.com/tomocy/ritty-for-branches/infra/http/route"
)

type branchAuth struct{}

func (b *branchAuth) BuildAuthorizationCodeURL() *url.URL {
	values := make(url.Values)
	values.Set("client_id", config.Current.BranchAuth.ClientID)
	values.Set("redirect_uri", config.Current.BranchAuth.RedirectURI)

	dest := *route.BranchAuthAPI.Route("authorization.prepare")
	dest.RawQuery = values.Encode()

	return &dest
}

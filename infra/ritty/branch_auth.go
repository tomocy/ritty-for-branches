package ritty

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/tomocy/ritty-for-branches/config"
	"github.com/tomocy/ritty-for-branches/domain/model"
	"github.com/tomocy/ritty-for-branches/infra/http/route"
)

type branchAuth struct{}

func (b *branchAuth) BuildAuthorizationCodeURL() *url.URL {
	values := make(url.Values)
	values.Set("client_id", config.Current.RittyBranchAuth.ClientID)
	values.Set("redirect_uri", config.Current.RittyBranchAuth.RedirectURI)

	dest := *route.RittyBranchAuthAPI.Route("authorization.prepare")
	dest.RawQuery = values.Encode()

	return &dest
}

func (b *branchAuth) FetchAuthorization(code string, cookies ...*http.Cookie) (*model.Authorization, error) {
	req, err := buildRequestToFetchAuthorization(code, cookies...)
	if err != nil {
		return nil, devErrorf("fatch authorization", err)
	}

	fetched, err := requestToFetchAuthorization(req)
	if err != nil {
		return nil, devErrorf("fetch authorization", err)
	}

	return fetched, nil
}

func buildRequestToFetchAuthorization(code string, cookies ...*http.Cookie) (*http.Request, error) {
	values := make(url.Values)
	values.Set("client_id", config.Current.RittyBranchAuth.ClientID)
	values.Set("client_secret", config.Current.RittyBranchAuth.ClientSecret)
	values.Set("code", code)
	values.Set("redirect_uri", config.Current.RittyBranchAuth.RedirectURI)

	req, err := http.NewRequest(
		http.MethodPost,
		route.RittyBranchAuthAPI.Route("authorization.create").String(),
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	return req, nil
}

func requestToFetchAuthorization(req *http.Request) (*model.Authorization, error) {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("content type is not supported")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dest *fetchedAuthorization
	if err := json.Unmarshal(body, &dest); err != nil {
		return nil, err
	}

	return dest.authorization(), nil
}

type fetchedAuthorization struct {
	Token struct {
		BranchID string
		Token    string
	}
}

func (a *fetchedAuthorization) authorization() *model.Authorization {
	return &model.Authorization{
		BranchID: a.Token.BranchID,
		Token:    a.Token.Token,
	}
}

package spotify

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/csothen/yt2spotify/utils"
)

const (
	// Scopes
	UserReadPrivate = "user-read-private"
	UserReadEmail   = "user-read-email"

	// Fields
	fieldScope        = "scope"
	fieldCode         = "code"
	fieldResponseType = "response_type"
	fieldClientID     = "client_id"
	fieldRedirectURI  = "redirect_uri"
	fieldState        = "state"
	fieldGrantType    = "grant_type"
	fieldRefreshToken = "refresh_token"

	// Form Values
	responseTypeCode           = "code"
	grantTypeClientCredentials = "client_credentials"
	grantTypeAuthorizationCode = "authorization_code"
	grantTypeRefreshToken      = "refresh_token"

	// URLs
	authorizationUrl = "https://accounts.spotify.com/authorize"
	apiTokenUrl      = "https://accounts.spotify.com/api/token"
)

func (s *service) AuthorizeClient() (*ClientAuth, error) {
	form := url.Values{}
	form.Set(fieldGrantType, grantTypeClientCredentials)

	resp, err := s.doFormRequest(http.MethodPost, apiTokenUrl, form)
	if err != nil {
		return nil, fmt.Errorf("failed to authorize client: %w", err)
	}

	auth := &ClientAuth{}
	err = utils.FromJSON(auth, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not parse response body: %w", err)
	}
	return auth, nil
}

func (s *service) StartUserAuthorization(scopes []string) (authUrl string, state string, err error) {
	scope := strings.Join(scopes, " ")
	state = utils.GenerateRandomString(16)
	parsedUrl, err := url.Parse(authorizationUrl)
	if err != nil {
		return "", "", err
	}

	query := parsedUrl.Query()
	query.Set(fieldScope, scope)
	query.Set(fieldResponseType, responseTypeCode)
	query.Set(fieldClientID, s.config.ClientID)
	query.Set(fieldRedirectURI, s.config.RedirectURL)
	query.Set(fieldState, state)

	parsedUrl.RawQuery = query.Encode()
	return parsedUrl.String(), state, nil
}

func (s *service) AuthorizeUser(code, state, storedState string) (*UserAuth, error) {
	if state == "" || state != storedState {
		return nil, fmt.Errorf("failed to authorize user: state mismatch")
	}

	form := url.Values{}
	form.Set(fieldGrantType, grantTypeAuthorizationCode)
	form.Set(fieldCode, code)
	form.Set(fieldRedirectURI, s.config.RedirectURL)

	resp, err := s.doFormRequest(http.MethodPost, apiTokenUrl, form)
	if err != nil {
		return nil, fmt.Errorf("could not authorize user: %w", err)
	}

	auth := &UserAuth{}
	err = utils.FromJSON(auth, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not parse response body: %w", err)
	}
	return auth, nil
}

func (s *service) RefreshUserAuthorization(refreshToken string) (*UserAuth, error) {
	form := url.Values{}
	form.Set(fieldGrantType, grantTypeRefreshToken)
	form.Set(fieldRefreshToken, refreshToken)

	resp, err := s.doFormRequest(http.MethodPost, apiTokenUrl, form)
	if err != nil {
		return nil, fmt.Errorf("could not refresh user authorization: %w", err)
	}

	auth := &UserAuth{}
	err = utils.FromJSON(auth, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not parse response body: %w", err)
	}
	return auth, nil
}

func (s *service) doFormRequest(method string, url string, form url.Values) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(s.config.ClientID, s.config.ClientSecret)
	req.Form = form

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

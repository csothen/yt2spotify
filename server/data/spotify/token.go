package spotify

import "time"

type UserAccessTokenResponse struct {
	AccessToken  string        `json:"access_token"`
	TokenType    string        `json:"token_type"`
	RefreshToken string        `json:"refresh_token"`
	Scope        string        `json:"scope"`
	ExpiresIn    time.Duration `json:"expires_in"`
}

type ClientAccessTokenResponse struct {
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
	ExpiresIn   time.Duration `json:"expires_in"`
}

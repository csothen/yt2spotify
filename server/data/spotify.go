package data

import "time"

type SpotifyUser struct {
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}

type SpotifyAPITokenBody struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	RedirectURI  string `json:"redirect_uri"`
}

type SpotifyAPITokenResponse struct {
	AccessToken  string        `json:"access_token"`
	TokenType    string        `json:"token_type"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresIn    time.Duration `json:"expires_in"`
	Error        string        `json:"error"`
}

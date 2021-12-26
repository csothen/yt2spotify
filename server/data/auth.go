package data

import "time"

type AuthURL struct {
	URL string `json:"url"`
}

type User struct {
	ID string `json:"id"`
	WithSpotifyAuth
}

type WithSpotifyAuth struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    time.Time `json:"expires_in"`
}

type IsAuthenticated struct {
	Status bool `json:"status"`
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

package data

import "time"

type AuthURL struct {
	URL string `json:"url"`
}

type User struct {
	SessionID string `json:"session_id"`
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

type CallbackData struct {
	SessionData string
	Location    string
}

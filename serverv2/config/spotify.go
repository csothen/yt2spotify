package config

import "github.com/csothen/env"

type spotifyConfiguration struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

func newSpotifyConfiguration() spotifyConfiguration {
	return spotifyConfiguration{
		ClientID:     env.String("SPOTIFY_CLIENT_ID", ""),
		ClientSecret: env.String("SPOTIFY_CLIENT_SECRET", ""),
		RedirectURI:  env.String("SPOTIFY_REDIRECT_URI", "http://localhost:8080/auth/spotify/redirect"),
	}
}

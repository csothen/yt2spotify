package config

import "github.com/csothen/env"

type spotifyConfiguration struct {
	SpotifyClientID     string
	SpotifyClientSecret string
	SpotifyRedirectURI  string
}

func newSpotifyConfiguration() spotifyConfiguration {
	return spotifyConfiguration{
		SpotifyClientID:     env.String("SPOTIFY_CLIENT_ID", ""),
		SpotifyClientSecret: env.String("SPOTIFY_CLIENT_SECRET", ""),
		SpotifyRedirectURI:  env.String("SPOTIFY_REDIRECT_URI", "http://localhost:8080/auth/spotify/redirect"),
	}
}

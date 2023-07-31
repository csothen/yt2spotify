package integrations

import (
	"github.com/csothen/yt2spotify/config"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

const (
	spotifyReadPlaylistsScope   = "playlist-read-private"
	spotifyModifyPlaylistsScope = "playlist-modify-private"
)

type spotifyIntegration struct {
	l      hclog.Logger
	config oauth2.Config
}

func NewSpotifyIntegration(l hclog.Logger, config config.Configuration) *spotifyIntegration {
	spotifyConfig := oauth2.Config{
		ClientID:     config.SpotifyClientID,
		ClientSecret: config.SpotifyClientSecret,
		RedirectURL:  config.SpotifyRedirectURI,
		Scopes: []string{
			spotifyReadPlaylistsScope,
			spotifyModifyPlaylistsScope,
		},
		Endpoint: spotify.Endpoint,
	}
	return &spotifyIntegration{l, spotifyConfig}
}

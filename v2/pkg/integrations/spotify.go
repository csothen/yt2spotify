package integrations

import (
	"context"
	"fmt"

	"github.com/csothen/yt2spotify/pkg/configuration"
	"github.com/csothen/yt2spotify/pkg/core"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

const (
	spotifyReadPlaylistsScope   = "playlist-read-private"
	spotifyModifyPlaylistsScope = "playlist-modify-private"
)

var Spotify *SpotifyIntegration

type SpotifyIntegration struct {
	config *oauth2.Config
}

func configureSpotify(config *configuration.Configuration) error {
	spotifyConfig, ok := config.Integrations["spotify"]
	if !ok {
		return fmt.Errorf("could not find integration configuration for key 'spotify'")
	}

	oauthConfig := &oauth2.Config{
		ClientID:     spotifyConfig.ClientID,
		ClientSecret: spotifyConfig.ClientSecret,
		Scopes:       []string{spotifyModifyPlaylistsScope, spotifyReadPlaylistsScope},
		Endpoint:     spotify.Endpoint,
	}

	Spotify = &SpotifyIntegration{oauthConfig}

	return nil
}

func (s *SpotifyIntegration) authenticate(code string) error {
	token, err := s.config.Exchange(context.Background(), code)
	if err != nil {
		return err
	}

	saveToken(core.SpotifySource, token)
	return nil
}

func (s *SpotifyIntegration) getAuthenticationURL() string {
	return s.config.AuthCodeURL("state", oauth2.AccessTypeOnline)
}

func (s *SpotifyIntegration) LoadPlaylist(url string) (*core.Playlist, error) {
	playlist := &core.Playlist{
		Source: core.SpotifySource,
		Name:   fmt.Sprintf("Playlist %s", url),
	}
	return playlist, nil
}

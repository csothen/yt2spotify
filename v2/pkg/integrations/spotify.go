package integrations

import (
	"context"
	"fmt"

	"github.com/csothen/env"
	"github.com/csothen/yt2spotify/pkg/core"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

const (
	spotifyReadPlaylistsScope   = "playlist-read-private"
	spotifyModifyPlaylistsScope = "playlist-modify-private"
)

var Spotify = newSpotify()

type SpotifyIntegration struct {
	config *oauth2.Config
}

func newSpotify() *SpotifyIntegration {
	config := &oauth2.Config{
		ClientID:     env.String("SPOTIFY_CLIENT_ID", ""),
		ClientSecret: env.String("SPOTIFY_CLIENT_SECRET", ""),
		Scopes:       []string{spotifyModifyPlaylistsScope, spotifyReadPlaylistsScope},
		Endpoint:     spotify.Endpoint,
	}

	return &SpotifyIntegration{config}
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

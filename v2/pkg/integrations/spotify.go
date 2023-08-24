package integrations

import (
	"fmt"

	"github.com/csothen/yt2spotify/pkg/core"
)

type Spotify struct{}

func NewSpotify() *Spotify {
	return &Spotify{}
}

func (s *Spotify) LoadPlaylist(url string) (*core.Playlist, error) {
	playlist := &core.Playlist{
		Source: core.SpotifySource,
		Name:   fmt.Sprintf("Playlist %s", url),
	}
	return playlist, nil
}

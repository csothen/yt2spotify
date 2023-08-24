package integrations

import (
	"fmt"

	"github.com/csothen/yt2spotify/pkg/core"
)

type Youtube struct{}

func NewYoutube() *Youtube {
	return &Youtube{}
}

func (y *Youtube) LoadPlaylist(url string) (*core.Playlist, error) {
	playlist := &core.Playlist{
		Source: core.YoutubeSource,
		Name:   fmt.Sprintf("Playlist %s", url),
	}
	return playlist, nil
}

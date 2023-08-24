package views

import (
	"fmt"

	"github.com/csothen/yt2spotify/pkg/core"
	"github.com/csothen/yt2spotify/pkg/integrations"
)

type ConvertPlaylistFormView struct {
	Sources []core.Source
}

type PlaylistView struct {
	Source core.Source
	Name   string
}

func GetConvertPlaylistFormView() *ConvertPlaylistFormView {
	return &ConvertPlaylistFormView{
		Sources: []core.Source{core.SpotifySource, core.YoutubeSource},
	}
}

func GetPlaylistView(source string, url string) (*PlaylistView, error) {
	var playlist *core.Playlist
	var err error

	switch source {
	case core.SpotifySource.Value:
		playlist, err = integrations.NewSpotify().LoadPlaylist(url)
		break
	case core.YoutubeSource.Value:
		playlist, err = integrations.NewYoutube().LoadPlaylist(url)
		break
	default:
		err = fmt.Errorf("invalid source '%s'", source)
		break
	}

	if err != nil {
		return nil, err
	}

	view := &PlaylistView{
		Source: playlist.Source,
		Name:   playlist.Name,
	}

	return view, nil
}

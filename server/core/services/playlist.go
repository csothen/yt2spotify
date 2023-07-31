package services

import (
	"fmt"

	"github.com/csothen/yt2spotify/core/models"
	"github.com/csothen/yt2spotify/integrations"
	"github.com/hashicorp/go-hclog"
)

type playlistService struct {
	l     hclog.Logger
	intgs map[models.Provider]integrations.Integration
}

func NewPlaylistService(l hclog.Logger, intgs map[models.Provider]integrations.Integration) PlaylistService {
	return &playlistService{l, intgs}
}

func (ps *playlistService) GetByUrl(provider models.Provider, url string) (models.Playlist, error) {
	switch provider {
	case models.Spotify:
		return nil, nil
	case models.Youtube:
		return nil, nil
	default:
		return nil, nil
	}
}

func (ps *playlistService) Convert(playlist models.Playlist, to models.Provider) (models.Playlist, error) {
	from := playlist.Provider()
	if from == to {
		err := fmt.Sprintf("tried to convert playlist to the same provider '%s'", to)
		ps.l.Error(err)
		return nil, fmt.Errorf(err)
	}

	switch from {
	case models.Spotify:
		return ps.convertSpotify(playlist.(*models.SpotifyPlaylist), to)
	case models.Youtube:
		return ps.convertYoutube(playlist.(*models.YoutubePlaylist), to)
	default:
		err := fmt.Sprintf("playlist provider '%s' not supported", from)
		ps.l.Error(err)
		return nil, fmt.Errorf(err)
	}
}

func (ps *playlistService) convertSpotify(playlist *models.SpotifyPlaylist, to models.Provider) (models.Playlist, error) {
	ps.l.Debug("converting Spotify playlist")
	switch to {
	case models.Youtube:
		return nil, nil
	default:
		err := fmt.Sprintf("convertion from '%s' to '%s' not supported", models.Spotify, to)
		ps.l.Error(err)
		return nil, fmt.Errorf(err)
	}
}

func (ps *playlistService) convertYoutube(playlist *models.YoutubePlaylist, to models.Provider) (models.Playlist, error) {
	ps.l.Debug("converting Youtube playlist")
	switch to {
	case models.Spotify:
		return nil, nil
	default:
		err := fmt.Sprintf("convertion from '%s' to '%s' not supported", models.Youtube, to)
		ps.l.Error(err)
		return nil, fmt.Errorf(err)
	}
}

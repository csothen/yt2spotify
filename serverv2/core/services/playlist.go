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

func (ps *playlistService) Convert(playlist models.Playlist) (models.Playlist, error) {
	provider := playlist.Provider()
	switch provider {
	case models.Spotify:
		ps.l.Debug("converting Spotify playlist")
	case models.Youtube:
		ps.l.Debug("converting Youtube playlist")
	default:
		msg := fmt.Sprintf("playlist provider '%s' not supported", provider)
		ps.l.Error(msg)
		return nil, fmt.Errorf(msg)
	}
	return playlist, nil
}

package integrations

import "github.com/csothen/yt2spotify/core/models"

type Integration interface {
	// TODO: Refine the Integration interface
	GetPlaylistByUrl(url string) (models.Playlist, error)
	GetTrackByName(name string) (models.Track, error)
}

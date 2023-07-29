package core

import "github.com/csothen/yt2spotify/data/core/playlist"

type Service interface {
	GeneratePlaylists(from string, to []string, playlist *playlist.List) ([]*playlist.List, error)
	DiscardMediaSource(pid string, idx int) error
	NextMediaSource(pid string, idx int) (*playlist.Source, error)
	ReplaceMediaSource(pid string, idx int, query string) (*playlist.Source, error)
	SavePlaylists(playlists []*playlist.List) ([]string, error)
}

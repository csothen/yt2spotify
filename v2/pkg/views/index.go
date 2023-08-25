package views

import "github.com/csothen/yt2spotify/pkg/core"

type HeaderView struct {
	Title string
}

type IndexView struct {
	Header    HeaderView
	Playlists []core.Playlist
	Sources   []core.Source
}

func GetIndexView() (IndexView, error) {
	view := IndexView{
		Header:    HeaderView{Title: "Home"},
		Playlists: []core.Playlist{},
		Sources:   []core.Source{core.YoutubeSource, core.SpotifySource},
	}
	return view, nil
}

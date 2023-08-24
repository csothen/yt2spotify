package core

type Source struct {
	Value string
	Name  string
}

var (
	YoutubeSource = Source{Value: "youtube", Name: "Youtube"}
	SpotifySource = Source{Value: "spotify", Name: "Spotify"}
)

type Playlist struct {
	Source Source
	Name   string
}

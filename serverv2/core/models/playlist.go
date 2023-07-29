package models

import (
	"github.com/google/uuid"
)

type Provider string

const (
	Spotify Provider = "spotify"
	Youtube Provider = "youtube"
)

type Playlist interface {
	Provider() Provider
}

type SpotifyPlaylist struct {
	ID          uuid.UUID
	ExternalID  string
	Name        string
	Description string
	URI         string
	Count       uint
	Owner       string
	Image       Image
	Tracks      []*SpotifyTrack
}

type YoutubePlaylist struct {
	ID          uuid.UUID
	ExternalID  string
	Name        string
	Description string
	URI         string
	Count       uint
	Owner       string
	Image       Image
	Tracks      []*YoutubeTrack
}

func (sp *SpotifyPlaylist) Provider() Provider {
	return Spotify
}

func (yp *YoutubePlaylist) Provider() Provider {
	return Youtube
}

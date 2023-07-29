package models

import "github.com/google/uuid"

type SpotifyTrack struct {
	ID         uuid.UUID
	ExternalID string
	Title      string
	URI        string
	Position   uint
	Duration   uint32
	Owner      string
}

type YoutubeTrack struct {
	ID         uuid.UUID
	ExternalID string
	Title      string
	URI        string
	Position   uint
	Duration   uint32
	Owner      string
}

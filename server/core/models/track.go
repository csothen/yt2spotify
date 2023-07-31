package models

import "github.com/google/uuid"

type Track struct {
	ID         uuid.UUID
	ExternalID string
	Title      string
	URI        string
	Position   uint
	Duration   uint32
	Owner      string
}

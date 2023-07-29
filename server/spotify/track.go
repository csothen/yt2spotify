package spotify

import (
	"fmt"
	"log"
)

func (s *service) SearchTracks(opts SearchOpts) []*Track {
	token, err := s.checkAuth()
	if err != nil {
		log.Printf("failed to search tracks: %w\n", err)
		return nil
	}
	return nil
}

func (s *service) FindTrack(query string) (*Track, error) {
	token, err := s.checkAuth()
	if err != nil {
		return nil, fmt.Errorf("failed to find track: %w", err)
	}
	return nil, nil
}

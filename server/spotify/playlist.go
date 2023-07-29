package spotify

import "fmt"

func (s *service) GetPlaylist(url string) (*Playlist, error) {
	token, err := s.checkAuth()
	if err != nil {
		return nil, fmt.Errorf("failed to get playlist: %w", err)
	}
	return nil, nil
}

func (s *service) CreatePlaylist(playlist *Playlist) (*Playlist, error) {
	token, err := s.checkAuth()
	if err != nil {
		return nil, fmt.Errorf("failed to create playlist: %w", err)
	}
	return nil, nil
}

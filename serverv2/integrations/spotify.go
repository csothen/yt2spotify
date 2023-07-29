package integrations

import "github.com/hashicorp/go-hclog"

type spotifyIntegration struct {
	l hclog.Logger
}

func NewSpotifyIntegration(l hclog.Logger) *spotifyIntegration {
	return &spotifyIntegration{l}
}

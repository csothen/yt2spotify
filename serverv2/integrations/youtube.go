package integrations

import "github.com/hashicorp/go-hclog"

type youtubeIntegration struct {
	l hclog.Logger
}

func NewYoutubeIntegration(l hclog.Logger) *youtubeIntegration {
	return &youtubeIntegration{l}
}

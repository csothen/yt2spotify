package integrations

import (
	"context"
	"fmt"

	"github.com/csothen/env"
	"github.com/csothen/yt2spotify/pkg/core"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

var Youtube = newYoutube()

type YoutubeIntegration struct {
	config *oauth2.Config
}

func newYoutube() *YoutubeIntegration {
	config := &oauth2.Config{
		ClientID:     env.String("YOUTUBE_CLIENT_ID", ""),
		ClientSecret: env.String("YOUTUBE_CLIENT_SECRET", ""),
		Scopes:       []string{youtube.YoutubeScope},
		Endpoint:     google.Endpoint,
	}

	return &YoutubeIntegration{config}
}

func (y *YoutubeIntegration) authenticate(code string) error {
	token, err := y.config.Exchange(context.Background(), code)
	if err != nil {
		return err
	}

	saveToken(core.YoutubeSource, token)
	return nil
}

func (y *YoutubeIntegration) getAuthenticationURL() string {
	return y.config.AuthCodeURL("state", oauth2.AccessTypeOnline)
}

func (y *YoutubeIntegration) LoadPlaylist(url string) (*core.Playlist, error) {
	playlist := &core.Playlist{
		Source: core.YoutubeSource,
		Name:   fmt.Sprintf("Playlist %s", url),
	}
	return playlist, nil
}

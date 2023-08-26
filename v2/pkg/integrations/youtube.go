package integrations

import (
	"context"
	"fmt"

	"github.com/csothen/yt2spotify/pkg/configuration"
	"github.com/csothen/yt2spotify/pkg/core"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

var Youtube *YoutubeIntegration

type YoutubeIntegration struct {
	config *oauth2.Config
}

func configureYoutube(config *configuration.Configuration) error {
	youtubeConfig, ok := config.Integrations["youtube"]
	if !ok {
		return fmt.Errorf("could not find integration configuration for key 'youtube'")
	}

	oauthConfig := &oauth2.Config{
		ClientID:     youtubeConfig.ClientID,
		ClientSecret: youtubeConfig.ClientSecret,
		Scopes:       []string{youtube.YoutubeScope},
		Endpoint:     google.Endpoint,
	}

	Youtube = &YoutubeIntegration{oauthConfig}

	return nil
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

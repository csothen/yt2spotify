package integrations

import (
	"github.com/csothen/yt2spotify/config"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

type youtubeIntegration struct {
	l         hclog.Logger
	config    oauth2.Config
	cacheFile string
}

func NewYoutubeIntegration(l hclog.Logger, config config.Configuration) *youtubeIntegration {
	youtubeConfig := oauth2.Config{
		ClientID:     config.SpotifyClientID,
		ClientSecret: config.SpotifyClientSecret,
		RedirectURL:  config.SpotifyRedirectURI,
		Scopes:       []string{youtube.YoutubeScope},
		Endpoint:     google.Endpoint,
	}

	cacheFile, err := loadCacheFile("youtube.json")
	if err != nil {
		l.Error("could not load cache file for youtube integration", "error", err)
		panic(err)
	}

	return &youtubeIntegration{l, youtubeConfig, cacheFile}
}

func (yi *youtubeIntegration) GetToken() *oauth2.Token {

}

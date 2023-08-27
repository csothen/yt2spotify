package integrations

import (
	"fmt"

	"github.com/csothen/yt2spotify/pkg/core"
	"golang.org/x/oauth2"
)

func Authenticate(source string, code string) (*oauth2.Token, error) {
	switch source {
	case core.SpotifySource.Value:
		return Spotify.authenticate(code)
	case core.YoutubeSource.Value:
		return Youtube.authenticate(code)
	default:
		return nil, fmt.Errorf("invalid source '%s'", source)
	}
}

func GetAuthenticationURL(source string) (string, error) {
	switch source {
	case core.SpotifySource.Value:
		return Spotify.getAuthenticationURL(), nil
	case core.YoutubeSource.Value:
		return Youtube.getAuthenticationURL(), nil
	default:
		return "", fmt.Errorf("invalid source '%s'", source)
	}
}

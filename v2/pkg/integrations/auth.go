package integrations

import (
	"fmt"

	"github.com/csothen/yt2spotify/pkg/core"
	"golang.org/x/oauth2"
)

func Authenticate(source string, code string) error {
	switch source {
	case core.SpotifySource.Value:
		return Spotify.authenticate(code)
	case core.YoutubeSource.Value:
		return Youtube.authenticate(code)
	default:
		return fmt.Errorf("invalid source '%s'", source)
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

func getToken(source core.Source) *oauth2.Token {
	// TODO: Actually retrieve the token
	return nil
}

func saveToken(source core.Source, token *oauth2.Token) {
	// TODO: Actually save the token
	fmt.Printf("token from source '%s' -> '%s'", source.Name, token.AccessToken)
}

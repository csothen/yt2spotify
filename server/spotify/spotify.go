package spotify

import (
	"log"
	"net/http"
	"os"
)

type Service interface {
	// Client Authorization
	AuthorizeClient() (*ClientAuth, error)
	// OAuth 2.0 Authorization
	StartUserAuthorization(scopes []string) (string, string, error)
	AuthorizeUser(code, state, storedState string) (*UserAuth, error)
	RefreshUserAuthorization(refreshToken string) (*UserAuth, error)
	// Playlists
	GetPlaylist(url string) (*Playlist, error)
	CreatePlaylist(playlist *Playlist) (*Playlist, error)
	// Tracks
	SearchTracks(opts SearchOpts) []*Track
	FindTrack(query string) (*Track, error)
}

type service struct {
	log    *log.Logger
	config *Configuration
	auth   *auth
	client *http.Client
}

func New(c *Configuration) Service {
	l := log.New(os.Stdout, "[ yt2spotify:spotify ] - ", log.LstdFlags)
	s := &service{
		log:    l,
		config: c,
		auth:   &auth{},
		client: &http.Client{},
	}

	// try to initialize the client authorization
	clientAuth, err := s.AuthorizeClient()
	if err != nil {
		log.Println(err)
		return s
	}
	s.auth.client = clientAuth
	return s
}

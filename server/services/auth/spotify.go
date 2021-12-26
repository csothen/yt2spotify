package auth

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/csothen/yt2spotify/config"
	"github.com/csothen/yt2spotify/data"
	"github.com/csothen/yt2spotify/mysql/auth"
	"github.com/google/uuid"
)

const (
	spotifyScopes = "playlist-modify-private"
)

type SpotifyAuthService struct {
	config *config.Config
	repo   *auth.MySQLRepository
}

func NewSpotifyAuthService(c *config.Config, db *sql.DB) *SpotifyAuthService {
	return &SpotifyAuthService{
		config: c,
		repo:   auth.NewMySQLRepository(db),
	}
}

func (s *SpotifyAuthService) BuildAuthURL() (*data.AuthURL, error) {
	authURL, err := url.Parse("https://accounts.spotify.com/authorize")
	if err != nil {
		return nil, err
	}

	query := authURL.Query()
	query.Set("scope", spotifyScopes)
	query.Set("response_type", "code")
	query.Set("client_id", s.config.ClientID)
	query.Set("client_secret", s.config.ClientSecret)
	query.Set("redirect_uri", s.config.RedirectURI)

	authURL.RawQuery = query.Encode()

	result := &data.AuthURL{
		URL: authURL.String(),
	}

	return result, nil
}

func (s *SpotifyAuthService) HandleCallback(code, qErr string) (string, error) {
	if qErr != "" {
		return "", fmt.Errorf("got error from spotify callback: %s", qErr)
	}

	payload, err := json.Marshal(&data.SpotifyAPITokenBody{
		GrantType:    "authorization_code",
		Code:         code,
		ClientID:     s.config.ClientID,
		ClientSecret: s.config.ClientSecret,
		RedirectURI:  s.config.RedirectURI,
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post("https://accounts.spotify.com/api/token", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	d := &data.SpotifyAPITokenResponse{}
	err = data.FromJSON(d, resp.Body)
	if err != nil {
		return "", err
	}

	if d.Error != "" {
		return "", fmt.Errorf("error retrieving spotify token: %s", d.Error)
	}

	u := &data.User{
		ID: uuid.NewString(),
		WithSpotifyAuth: data.WithSpotifyAuth{
			AccessToken:  d.AccessToken,
			RefreshToken: d.RefreshToken,
			TokenType:    d.TokenType,
			ExpiresIn:    time.Now().Add(d.ExpiresIn * time.Second),
		},
	}

	_, err = s.repo.Upsert(u)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/start", s.config.ClientLocation), nil
}

func (s *SpotifyAuthService) IsAuthenticated(id string) *data.IsAuthenticated {
	result := &data.IsAuthenticated{}

	u, err := s.repo.GetByID(id)
	if err != nil {
		result.Status = false
		return result
	}

	if u.ExpiresIn.Before(time.Now()) {
		err := s.refreshToken(u)
		if err != nil {
			result.Status = false
			return result
		}
	}

	result.Status = true
	return result
}

func (s *SpotifyAuthService) RefreshToken(id string) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.refreshToken(user)
}

func (s *SpotifyAuthService) refreshToken(user *data.User) error {
	payload, err := json.Marshal(&data.SpotifyAPITokenBody{
		GrantType:    "refresh_token",
		RefreshToken: user.RefreshToken,
		ClientID:     s.config.ClientID,
		ClientSecret: s.config.ClientSecret,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post("https://accounts.spotify.com/api/token", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	d := &data.SpotifyAPITokenResponse{}
	err = data.FromJSON(d, resp.Body)
	if err != nil {
		return err
	}

	if d.Error != "" {
		return fmt.Errorf("error refreshing spotify token: %s", d.Error)
	}

	u := &data.User{
		ID: user.ID,
		WithSpotifyAuth: data.WithSpotifyAuth{
			AccessToken:  d.AccessToken,
			RefreshToken: d.RefreshToken,
			TokenType:    d.TokenType,
			ExpiresIn:    time.Now().Add(d.ExpiresIn * time.Second),
		},
	}

	_, err = s.repo.Upsert(u)
	if err != nil {
		return err
	}

	return nil
}

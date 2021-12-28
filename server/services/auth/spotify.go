package auth

import (
	"bytes"
	"database/sql"
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/csothen/yt2spotify/config"
	"github.com/csothen/yt2spotify/data"
	"github.com/csothen/yt2spotify/mysql/auth"
	"github.com/gorilla/securecookie"
)

const (
	jsonContentType   = "application/json"
	spotifyAPIBaseURL = "https://api.spotify.com/v1"
	spotifyScopes     = "playlist-modify-private user-read-email"
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
	query.Set("client_id", s.config.SpotifyClientID)
	query.Set("client_secret", s.config.SpotifyClientSecret)
	query.Set("redirect_uri", s.config.SpotifyRedirectURI)

	authURL.RawQuery = query.Encode()

	result := &data.AuthURL{
		URL: authURL.String(),
	}

	return result, nil
}

func (s *SpotifyAuthService) HandleCallback(id, code, qErr string) (*data.CallbackData, error) {
	if qErr != "" {
		return nil, fmt.Errorf("got error from spotify callback: %s", qErr)
	}

	payload := url.Values{}
	payload.Set("grant_type", "authorization_code")
	payload.Set("code", code)
	payload.Set("redirect_uri", s.config.SpotifyRedirectURI)

	req, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}

	b64Creds := base64.RawStdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", s.config.SpotifyClientID, s.config.SpotifyClientSecret)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload.Encode())))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", b64Creds))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error requesting token from spotify")
	}

	d := &data.SpotifyAPITokenResponse{}
	err = data.FromJSON(d, resp.Body)
	if err != nil {
		return nil, err
	}

	if d.Error != "" {
		return nil, fmt.Errorf("error retrieving spotify token: %s", d.Error)
	}

	if id == "" {
		id = strings.TrimRight(
			base32.StdEncoding.EncodeToString(
				securecookie.GenerateRandomKey(32)), "=")
	}

	u := &data.User{
		SessionID: id,
		WithSpotifyAuth: data.WithSpotifyAuth{
			AccessToken:  d.AccessToken,
			RefreshToken: d.RefreshToken,
			TokenType:    d.TokenType,
			ExpiresIn:    time.Now().Add(d.ExpiresIn * time.Second),
		},
	}

	_, err = s.repo.Upsert(u)
	if err != nil {
		return nil, err
	}

	cbData := &data.CallbackData{
		SessionData: id,
		Location:    fmt.Sprintf("%s/start", s.config.ClientLocation),
	}

	return cbData, nil
}

func (s *SpotifyAuthService) GetUserData(token string) (*data.SpotifyUser, error) {
	req, err := http.NewRequest(http.MethodGet, spotifyAPIBaseURL+"/me", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	d := &data.SpotifyUser{}
	err = data.FromJSON(d, resp.Body)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (s *SpotifyAuthService) IsAuthenticated(sessionId string) *data.IsAuthenticated {
	result := &data.IsAuthenticated{}

	u, err := s.repo.GetBySessionID(sessionId)
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

func (s *SpotifyAuthService) RefreshToken(sessionId string) error {
	user, err := s.repo.GetBySessionID(sessionId)
	if err != nil {
		return err
	}

	return s.refreshToken(user)
}

func (s *SpotifyAuthService) refreshToken(user *data.User) error {
	payload, err := json.Marshal(&data.SpotifyAPITokenBody{
		GrantType:    "refresh_token",
		RefreshToken: user.RefreshToken,
		ClientID:     s.config.SpotifyClientID,
		ClientSecret: s.config.SpotifyClientSecret,
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
		SessionID: user.SessionID,
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

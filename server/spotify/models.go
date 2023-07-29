package spotify

import "time"

type auth struct {
	user   *UserAuth
	client *ClientAuth
}

type Configuration struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

type ClientAuth struct {
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
	ExpiresIn   time.Duration `json:"expires_in"`
}

type UserAuth struct {
	AccessToken  string        `json:"access_token"`
	TokenType    string        `json:"token_type"`
	RefreshToken string        `json:"refresh_token"`
	Scope        string        `json:"scope"`
	ExpiresIn    time.Duration `json:"expires_in"`
}

type Playlist struct {
	ID            string        `json:"id"`
	SnapshotID    string        `json:"snapshot_id"`
	HRef          string        `json:"href"`
	Name          string        `json:"name"`
	Owner         *PublicUser   `json:"owner"`
	Images        []*Image      `json:"images"`
	Description   *string       `json:"description"`
	ExternalUrls  *ExternalUrls `json:"external_urls"`
	Followers     *Followers    `json:"followers"`
	Tracks        *Tracks       `json:"tracks"`
	Collaborative bool          `json:"collaborative"`
	Public        bool          `json:"public"`
	Type          string        `json:"type"`
	URI           string        `json:"uri"`
}

type Tracks struct {
	HRef     string   `json:"href"`
	Items    []*Track `json:"items"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Limit    int      `json:"limit"`
	Offset   int      `json:"offset"`
	Total    int      `json:"total"`
}

type Track struct {
	ID               string        `json:"id"`
	Name             string        `json:"name"`
	Type             string        `json:"type"`
	Images           []*Image      `json:"images"`
	Album            *Album        `json:"album"`
	Artists          []*Artist     `json:"artists"`
	AvailableMarkets []string      `json:"available_markets"`
	DiscNumber       int           `json:"disc_number"`
	DurationMs       int           `json:"duration_ms"`
	TrackNumber      int           `json:"track_number"`
	Popularity       int           `json:"popularity"`
	HRef             string        `json:"href"`
	PreviewURL       string        `json:"preview_url"`
	URI              string        `json:"uri"`
	ExternalUrls     *ExternalUrls `json:"external_urls"`
	Explicit         bool          `json:"explicit"`
	IsPlayable       bool          `json:"is_playable"`
	IsLocal          bool          `json:"is_local"`
}

type Album struct {
	ID                   string          `json:"id"`
	Name                 string          `json:"name"`
	Type                 string          `json:"type"`
	AlbumType            string          `json:"album_type"`
	AlbumGroup           string          `json:"album_group"`
	URI                  string          `json:"uri"`
	HRef                 string          `json:"href"`
	Artists              []*PublicArtist `json:"artists"`
	Images               []*Image        `json:"images"`
	ExternalUrls         *ExternalUrls   `json:"external_urls"`
	AvailableMarkets     []string        `json:"available_markets"`
	TotalTracks          int             `json:"total_tracks"`
	ReleaseDate          string          `json:"release_date"`
	ReleaseDatePrecision string          `json:"release_date_precision"`
}

type PublicArtist struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	HRef         string        `json:"href"`
	URI          string        `json:"uri"`
	Type         string        `json:"type"`
	ExternalUrls *ExternalUrls `json:"external_urls"`
}

type Artist struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Genres       []string      `json:"genres"`
	Type         string        `json:"type"`
	URI          string        `json:"uri"`
	HRef         string        `json:"href"`
	Followers    *Followers    `json:"followers"`
	ExternalUrls *ExternalUrls `json:"external_urls"`
	Images       []*Image      `json:"images"`
	Popularity   int           `json:"popularity"`
}

type PublicUser struct {
	ID           string        `json:"id"`
	DisplayName  string        `json:"display_name"`
	Type         string        `json:"type"`
	HRef         string        `json:"href"`
	URI          string        `json:"uri"`
	ExternalUrls *ExternalUrls `json:"external_urls"`
	Followers    *Followers    `json:"followers"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}

type Followers struct {
	HRef  string `json:"href"`
	Total int    `json:"total"`
}

type SearchOpts struct {
	Query      string `json:"query"`
	MaxResults int    `json:"max_results"`
}

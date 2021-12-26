package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/csothen/yt2spotify/services/auth"
)

type Auth struct {
	l       *log.Logger
	service *auth.SpotifyAuthService
}

func NewAuth(l *log.Logger, s *auth.SpotifyAuthService) *Auth {
	return &Auth{
		l:       l,
		service: s,
	}
}

func (a *Auth) GetAuthURL(rw http.ResponseWriter, r *http.Request) {
	url, err := a.service.BuildAuthURL()
	if err != nil {
		a.l.Println(err)
		http.Error(rw, "Error building the URL", http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(rw)
	e.SetEscapeHTML(false)

	err = e.Encode(url)
	if err != nil {
		a.l.Println(err)
		http.Error(rw, "Error building the URL", http.StatusInternalServerError)
	}
}

func (a *Auth) HandleSpotifyCallback(rw http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	location, err := a.service.HandleCallback(q.Get("code"), q.Get("error"))
	if err != nil {
		a.l.Println(err)
		http.Error(rw, "Error authenticating", http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, location, http.StatusFound)
}

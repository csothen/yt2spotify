package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/csothen/yt2spotify/services/auth"
	"github.com/csothen/yt2spotify/services/sessions"
	"github.com/csothen/yt2spotify/utils"
	gsessions "github.com/gorilla/sessions"
)

type Auth struct {
	l       *log.Logger
	store   *sessions.CookieStore
	service *auth.SpotifyAuthService
}

func NewAuth(l *log.Logger, s *auth.SpotifyAuthService, store *sessions.CookieStore) *Auth {
	return &Auth{
		l:       l,
		store:   store,
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
	session, _ := a.store.Get(r)
	q := r.URL.Query()

	cbData, err := a.service.HandleCallback(session.ID, q.Get("code"), q.Get("error"))
	if err != nil {
		a.l.Println(err)
		http.Error(rw, "Error authenticating", http.StatusInternalServerError)
		return
	}

	http.SetCookie(rw, gsessions.NewCookie(session.Name(), cbData.SessionData, session.Options))
	http.Redirect(rw, r, cbData.Location, http.StatusFound)
}

func (a *Auth) CheckAuthorization(rw http.ResponseWriter, r *http.Request) {
	session, _ := a.store.Get(r)

	isAuth := a.service.IsAuthenticated(session.ID)
	err := utils.ToJSON(isAuth, rw)
	if err != nil {
		a.l.Println(err)
		http.Error(rw, "Error checking authorization", http.StatusInternalServerError)
	}
}

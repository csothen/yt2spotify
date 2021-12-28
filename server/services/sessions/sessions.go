package sessions

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type CookieStore struct {
	store *sessions.CookieStore
}

func NewStore() *CookieStore {
	return &CookieStore{
		store: sessions.NewCookieStore(securecookie.GenerateRandomKey(32)),
	}
}

func (s *CookieStore) Get(r *http.Request) (*sessions.Session, error) {
	return s.store.Get(r, "session")
}

func (s *CookieStore) GetNamed(r *http.Request, name string) (*sessions.Session, error) {
	return s.store.Get(r, name)
}

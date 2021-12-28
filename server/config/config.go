package config

import (
	"log"

	"github.com/csothen/env"
)

type Config struct {
	BindAddress    string
	ClientLocation string

	SpotifyClientID     string
	SpotifyClientSecret string
	SpotifyRedirectURI  string

	DBURL      string
	DBUser     string
	DBPassword string
	DBName     string
}

func New(l *log.Logger) *Config {
	return &Config{
		BindAddress:    env.String("BIND_ADDRESS", "8080"),
		ClientLocation: env.String("FE_URL", "http://localhost:3000"),

		SpotifyClientID:     env.String("SPOTIFY_CLIENT_ID", ""),
		SpotifyClientSecret: env.String("SPOTIFY_CLIENT_SECRET", ""),
		SpotifyRedirectURI:  env.String("SPOTIFY_REDIRECT_URI", "http://localhost:8080/auth/spotify/redirect"),

		DBURL:      env.String("DB_URL", "127.0.0.1:3306"),
		DBName:     env.String("DB_NAME", "test"),
		DBUser:     env.String("DB_USER", "username"),
		DBPassword: env.String("DB_PASSWORD", "password"),
	}
}

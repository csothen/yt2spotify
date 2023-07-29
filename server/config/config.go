package config

import (
	"log"

	"github.com/csothen/env"
)

type Config struct {
	ServerConfig
	SpotifyConfig
	DbConfig
}

type ServerConfig struct {
	BindAddress string
	ClientHost  string
}

type SpotifyConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

type DbConfig struct {
	Url      string
	User     string
	Password string
	Name     string
}

func New(l *log.Logger) *Config {
	return &Config{
		ServerConfig: ServerConfig{
			BindAddress: env.String("BIND_ADDRESS", "8080"),
			ClientHost:  env.String("CLIENT_HOST", "http://localhost:3000"),
		},
		SpotifyConfig: SpotifyConfig{
			ClientID:     env.String("SPOTIFY_CLIENT_ID", ""),
			ClientSecret: env.String("SPOTIFY_CLIENT_SECRET", ""),
			RedirectURI:  env.String("SPOTIFY_REDIRECT_URI", "http://localhost:8080/auth/spotify/redirect"),
		},
		DbConfig: DbConfig{
			Url:      env.String("DB_URL", "127.0.0.1:3306"),
			Name:     env.String("DB_NAME", "test"),
			User:     env.String("DB_USER", "username"),
			Password: env.String("DB_PASSWORD", "password"),
		},
	}
}

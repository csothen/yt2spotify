package config

import (
	"log"

	"github.com/csothen/env"
)

type Config struct {
	BindAddress    string
	ClientLocation string

	ClientID     string
	ClientSecret string
	RedirectURI  string

	DBUser     string
	DBPassword string
	DBName     string
}

func New(l *log.Logger) *Config {
	err := env.Load(".env")
	if err != nil {
		l.Println(err)
	}

	return &Config{
		BindAddress:    env.String("BIND_ADDRESS", ":8080"),
		ClientLocation: env.String("CLIENT_LOCATION", "http://localhost:3000"),

		ClientID:     env.String("CLIENT_ID", ""),
		ClientSecret: env.String("CLIENT_SECRET", ""),
		RedirectURI:  env.String("REDIRECT_URI", "http://localhost:3000/callback"),

		DBUser:     env.String("MYSQL_USER", "username"),
		DBPassword: env.String("MYSQL_PASSWORD", "password"),
		DBName:     env.String("MYSQL_NAME", "test"),
	}
}

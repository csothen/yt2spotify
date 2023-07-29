package config

import "github.com/csothen/env"

type serverConfiguration struct {
	ServerAddress string
	ClientHost    string
}

func newServerConfiguration() serverConfiguration {
	return serverConfiguration{
		ServerAddress: env.String("BIND_ADDRESS", "8080"),
		ClientHost:    env.String("CLIENT_HOST", "http://localhost:3000"),
	}
}

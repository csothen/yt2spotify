package config

import "github.com/csothen/env"

type serverConfiguration struct {
	ServerAddress string
}

func newServerConfiguration() serverConfiguration {
	return serverConfiguration{
		ServerAddress: env.String("BIND_ADDRESS", "8080"),
	}
}

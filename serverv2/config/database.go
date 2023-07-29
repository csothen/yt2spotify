package config

import "github.com/csothen/env"

type databaseConfiguration struct {
	DbUrl      string
	DbUser     string
	DbPassword string
	DbName     string
}

func newDatabaseConfiguration() databaseConfiguration {
	return databaseConfiguration{
		DbUrl:      env.String("DB_URL", "127.0.0.1:3306"),
		DbName:     env.String("DB_NAME", "test"),
		DbUser:     env.String("DB_USER", "username"),
		DbPassword: env.String("DB_PASSWORD", "password"),
	}
}

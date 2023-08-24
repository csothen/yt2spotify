package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/csothen/yt2spotify/config"
	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func Open(config config.Configuration) (*Database, error) {
	db, err := sql.Open("postgres", buildConnectionURL(config))
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func buildConnectionURL(config config.Configuration) string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", config.DbUser, config.DbPassword, config.DbUrl, config.DbName)
}

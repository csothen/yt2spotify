package mysql

import (
	"database/sql"
	"fmt"

	"github.com/csothen/yt2spotify/config"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(config *config.Config) (*sql.DB, error) {
	return sql.Open("mysql", buildConnectionURL(config))
}

func buildConnectionURL(config *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", config.DBUser, config.DBPassword, config.DBName)
}

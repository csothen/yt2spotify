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
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", config.DBUser, config.DBPassword, config.DBURL, config.DBName)
}

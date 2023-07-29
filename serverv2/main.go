package main

import (
	"github.com/csothen/yt2spotify/config"
	"github.com/csothen/yt2spotify/transport/http"
	"github.com/hashicorp/go-hclog"
)

func main() {
	cfg := config.Load()
	l := hclog.Default()
	s, err := http.NewServer(cfg, l)
	if err != nil {
		l.Error("failed to intialize the server", "error", err)
		return
	}
	s.Start()
}

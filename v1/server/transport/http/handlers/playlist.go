package handlers

import (
	"fmt"
	"net/http"

	"github.com/csothen/yt2spotify/core/services"
	"github.com/hashicorp/go-hclog"
)

type playlistHandler struct {
	l  hclog.Logger
	ps services.PlaylistService
}

func NewPlaylistHandler(l hclog.Logger, ps services.PlaylistService) *playlistHandler {
	return &playlistHandler{l, ps}
}

func (ph *playlistHandler) HandleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

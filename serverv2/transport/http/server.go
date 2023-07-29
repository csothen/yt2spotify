package http

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/csothen/yt2spotify/config"
	"github.com/csothen/yt2spotify/core/models"
	"github.com/csothen/yt2spotify/core/services"
	"github.com/csothen/yt2spotify/integrations"
	"github.com/csothen/yt2spotify/postgresql"
	"github.com/csothen/yt2spotify/transport/http/handlers"
	"github.com/csothen/yt2spotify/transport/http/middlewares"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

var (
	bindPort = flag.Int("port", 8080, "Bind port for the server")
)

type Server struct {
	l hclog.Logger
	s *http.Server
}

func NewServer(config config.Configuration, l hclog.Logger) (*Server, error) {
	// Create database connection
	db, err := postgresql.Open(config)
	if err != nil {
		l.Error("failed to open connection with database", "error", err)
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		l.Error("could not verify connection with database", "error", err)
		return nil, err
	}

	// Initialize integrations
	spotifyI := integrations.NewSpotifyIntegration(l)
	youtubeI := integrations.NewYoutubeIntegration(l)

	intgs := map[models.Provider]integrations.Integration{
		models.Spotify: spotifyI,
		models.Youtube: youtubeI,
	}

	// Create services
	ps := services.NewPlaylistService(l, intgs)

	// Create handlers
	ph := handlers.NewPlaylistHandler(l, ps)

	// Setup mux
	sm := mux.NewRouter()
	sm.Use(middlewares.Json)

	// Setup routes
	// TODO: Setup the routes for the endpoints
	sm.HandleFunc("/hello", ph.HandleHello).Methods(http.MethodGet)

	// Setup CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// Create server
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", *bindPort),                    // configure the bind address
		Handler:      ch(sm),                                           // set the default handler
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  5 * time.Second,                                  // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                 // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                // max time for connections using TCP Keep-Alive
	}

	return &Server{l, s}, nil
}

func (s *Server) Start() {
	go func() {
		s.l.Info("starting server", "port", *bindPort)

		err := s.s.ListenAndServe()
		if err != nil {
			s.l.Error("error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	log.Println("got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.s.Shutdown(ctx)
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/csothen/yt2spotify/config"
	"github.com/csothen/yt2spotify/handlers"
	"github.com/csothen/yt2spotify/middlewares"
	"github.com/csothen/yt2spotify/mysql"
	"github.com/csothen/yt2spotify/services/auth"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "[ yt2spotify ] - ", log.LstdFlags)
	cfg := config.New(l)

	// Establish connection with DB
	db, err := mysql.Connect(cfg)
	if err != nil {
		l.Fatal(err)
	}
	defer db.Close()

	// create services
	as := auth.NewSpotifyAuthService(cfg, db)

	// create handlers
	ah := handlers.NewAuth(l, as)

	// create a new serve mux
	sm := mux.NewRouter()
	sm.Use(middlewares.Json)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/auth/url", ah.GetAuthURL)
	getRouter.HandleFunc("/auth/redirect", ah.HandleSpotifyCallback)

	// create a new server
	s := http.Server{
		Addr:         cfg.BindAddress,   // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Printf("Starting server on port %s\n", cfg.BindAddress[1:])

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}

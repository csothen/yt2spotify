package main

import (
	"context"
	"fmt"
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
	"github.com/csothen/yt2spotify/services/sessions"
	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "[ yt2spotify ] - ", log.LstdFlags)
	cfg := config.New(l)

	l.Printf("%+v\n", cfg)

	// Establish connection with DB
	db, err := mysql.Connect(cfg)
	if err != nil {
		l.Fatal(err)
	}
	defer db.Close()

	l.Println("Pinging DB")
	if err := db.Ping(); err != nil {
		l.Fatal(err)
	}

	// create session store
	store := sessions.NewStore()

	// create services
	as := auth.NewSpotifyAuthService(cfg, db)

	// create handlers
	ah := handlers.NewAuth(l, as, store)

	// create a new serve mux
	sm := mux.NewRouter()
	sm.Use(middlewares.Json)

	credentials := gh.AllowCredentials()
	origins := gh.AllowedOrigins([]string{"*"})

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/auth/url", ah.GetAuthURL)
	getRouter.HandleFunc("/auth/spotify/redirect", ah.HandleSpotifyCallback)
	getRouter.HandleFunc("/auth/spotify/check", ah.CheckAuthorization)

	addr := fmt.Sprintf(":%s", cfg.BindAddress)
	// create a new server
	s := http.Server{
		Addr:         addr,                              // configure the bind address
		Handler:      gh.CORS(credentials, origins)(sm), // set the default handler
		ErrorLog:     l,                                 // set the logger for the server
		ReadTimeout:  5 * time.Second,                   // max time to read request from the client
		WriteTimeout: 10 * time.Second,                  // max time to write response to the client
		IdleTimeout:  120 * time.Second,                 // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Printf("Starting server on port %s\n", cfg.BindAddress)

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

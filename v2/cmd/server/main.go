package main

import (
	"log"
	"text/template"

	"github.com/csothen/yt2spotify/pkg/configuration"
	"github.com/csothen/yt2spotify/pkg/endpoints"
	"github.com/csothen/yt2spotify/pkg/integrations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config, err := configuration.Load("config.yaml")
	if err != nil {
		log.Fatalf("could not load the configuration: %e\n", err)
	}

	tmpl, err := template.ParseGlob("public/views/*.html")
	if err != nil {
		log.Fatalf("couldn't initialize templates: %e\n", err)
	}

	err = integrations.Configure(config)
	if err != nil {
		log.Fatalf("could not configure the integrations: %e\n", err)
	}

	e := echo.New()
	e.Renderer = endpoints.NewTemplateRenderer(tmpl)

	e.Use(middleware.Logger())
	e.Static("/css", "css")
	e.Static("/assets", "assets")

	handler := endpoints.NewRequestHandler()

	e.GET("/", handler.HandleIndex)
	e.GET("/convert/start", handler.HandleStartConvertion)
	e.POST("/playlist/load", handler.HandlePlaylistLoad)

	e.POST("/oauth/:source", handler.HandleOAuth)
	e.GET("/oauth/:source/callback", handler.HandleCallback)

	e.Logger.Fatal(e.Start(config.Server.Port))
}

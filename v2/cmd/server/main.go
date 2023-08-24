package main

import (
	"log"
	"text/template"

	"github.com/csothen/yt2spotify/pkg/endpoints"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	tmpl, err := template.ParseGlob("public/views/*.html")
	if err != nil {
		log.Fatalf("couldn't initialize templates: %e\n", err)
	}

	e := echo.New()
	e.Renderer = endpoints.NewTemplateRenderer(tmpl)

	e.Use(middleware.Logger())
	e.Static("/css", "css")
	e.Static("/assets", "assets")

	e.GET("/", endpoints.HandleIndex)
	e.GET("/convert/start", endpoints.HandleStartConvertion)
	e.POST("/playlist/load", endpoints.HandlePlaylistLoad)

	e.Logger.Fatal(e.Start(":8080"))
}

package endpoints

import (
	"net/http"

	"github.com/csothen/yt2spotify/pkg/views"
	"github.com/labstack/echo/v4"
)

func HandleIndex(c echo.Context) error {
	view, err := views.GetIndexView()
	if err != nil {
		c.Logger().Errorf("could not load index view: %+v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	return c.Render(200, "base", view)
}

func HandlePlaylistLoad(c echo.Context) error {
	source := c.FormValue("source")
	url := c.FormValue("url")

	view, err := views.GetPlaylistView(source, url)
	if err != nil {
		c.Logger().Errorf("could not load playlist view: %+v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	return c.Render(http.StatusOK, "playlist", view)
}

func HandleStartConvertion(c echo.Context) error {
	view := views.GetConvertPlaylistFormView()
	return c.Render(http.StatusOK, "convert-playlist-form", view)
}

package endpoints

import (
	"fmt"
	"net/http"

	"github.com/csothen/yt2spotify/pkg/integrations"
	"github.com/csothen/yt2spotify/pkg/views"
	"github.com/labstack/echo/v4"
)

type RequestHandler struct {
}

func NewRequestHandler() RequestHandler {
	return RequestHandler{}
}

func (h *RequestHandler) HandleIndex(c echo.Context) error {
	view, err := views.GetIndexView()
	if err != nil {
		c.Logger().Errorf("could not load index view: %+v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	return c.Render(http.StatusOK, "base", view)
}

func (h *RequestHandler) HandlePlaylistLoad(c echo.Context) error {
	source := c.FormValue("source")
	url := c.FormValue("url")

	view, err := views.GetPlaylistView(source, url)
	if err != nil {
		c.Logger().Errorf("could not load playlist view: %+v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	return c.Render(http.StatusOK, "playlist", view)
}

func (h *RequestHandler) HandleStartConvertion(c echo.Context) error {
	view := views.GetConvertPlaylistFormView()
	return c.Render(http.StatusOK, "convert-playlist-form", view)
}

func (h *RequestHandler) HandleOAuth(c echo.Context) error {
	source := c.Param("source")

	url, err := integrations.GetAuthenticationURL(source)
	if err != nil {
		c.Logger().Error("failed to retrieve authentication url: %+v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	c.Response().Header().Set("HX-Redirect", url)
	return c.String(http.StatusOK, url)
}

func (h *RequestHandler) HandleCallback(c echo.Context) error {
	source := c.Param("source")
	code := c.QueryParam("code")

	token, err := integrations.Authenticate(source, code)
	if err != nil {
		c.Logger().Error("failed to authenticate: %+v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	fmt.Printf("token from source '%s' -> '%s'\n", source, token.AccessToken)

	return c.Redirect(http.StatusPermanentRedirect, c.Request().Host)
}

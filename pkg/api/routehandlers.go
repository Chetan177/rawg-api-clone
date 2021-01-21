package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) getGames(c echo.Context) error{

	return c.JSON(http.StatusOK,"games api working")
}
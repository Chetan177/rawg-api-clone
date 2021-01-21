package api

import "github.com/labstack/echo/v4"

type route struct {
	Path    string
	Handler func(echo.Context) error
}

func (s *Server) loadGamesAPIRoutes(group *echo.Group) {
	routes := s.getGameAPIRoutes()
	for _, route := range routes {
		group.GET(route.Path, route.Handler)
	}
}

func (s *Server) getGameAPIRoutes() []route {
	return []route{
		{"", s.getGames},

	}
}

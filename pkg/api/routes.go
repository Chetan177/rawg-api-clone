package api

import "github.com/labstack/echo/v4"

// route structure for ap path and handler
type route struct {
	Path    string
	Handler func(echo.Context) error
}

// loadGamesAPIRoutes loads /game api routes
func (s *Server) loadGamesAPIRoutes(group *echo.Group) {
	routes := s.getGameAPIRoutes()
	for _, route := range routes {
		group.GET(route.Path, route.Handler)
	}
}

// getGameAPIRoutes method returns routes
func (s *Server) getGameAPIRoutes() []route {
	return []route{
		{"", s.getGames},

	}
}

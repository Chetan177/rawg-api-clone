// Package api provides API Server and all its methods
package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// API Server struct
type Server struct {
	port       int
	echoServer *echo.Echo
}

const gameRoute = "/games"

var (
	once   sync.Once
	server *Server
)

// GetServer returns an instance of API server
func GetServer(restPort int) *Server {
	once.Do(func() {
		server = createServer(restPort)
	})

	return server
}

// createServer method create and configure api server
func createServer(restPort int) *Server {
	apiServer := echo.New()
	apiServer.Use(middleware.Recover())
	apiServer.Pre(middleware.RemoveTrailingSlash())
	apiServer.GET("/health", healthCheck)

	return &Server{
		port:       restPort,
		echoServer: apiServer,
	}
}

// StartServer method launches the api server as a go routine
func (s *Server) StartServer() error {
	gamesAPIGroup := s.echoServer.Group(gameRoute)
	s.loadGamesAPIRoutes(gamesAPIGroup)

	//start rest server
	go func(port int) {
		addr := fmt.Sprintf(":%d", port)
		s.echoServer.Logger.Fatal(s.echoServer.Start(addr))
	}(s.port)

	return nil
}

// Shutdown method stops the api server
func (s *Server) Shutdown() {
	_ = s.echoServer.Shutdown(context.Background())
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "echo sever running")
}

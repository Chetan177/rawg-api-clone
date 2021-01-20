package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	port       int
	echoServer *echo.Echo
}

var (
	once   sync.Once
	server *Server
)

func GetServer(restPort int) *Server {
	once.Do(func() {
		server = createServer(restPort)
	})

	return server
}

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

func (s *Server) StartServer() error {

	//start rest server
	go func(port int) {
		addr := fmt.Sprintf(":%d", port)
		s.echoServer.Logger.Fatal(s.echoServer.Start(addr))
	}(s.port)

	return nil
}

func (s *Server) Shutdown() {
	_ = s.echoServer.Shutdown(context.Background())
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "echo sever running")
}

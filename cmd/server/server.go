package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"rawg/pkg/api"
)

// Main method to start the API server
func main() {
	sig := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	server := api.GetServer(9000)
	err := server.StartServer()
	if err != nil {
		log.Error("error starting api server: %+v", err)
	}

	go func() {
		sig := <-sig
		fmt.Println(sig)
		done <- true
	}()

	<-done
	server.Shutdown()
}

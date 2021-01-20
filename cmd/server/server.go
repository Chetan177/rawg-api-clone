package main

import (
	"fmt"
	"os"
	"rawg/pkg/api"

	"github.com/labstack/gommon/log"
)

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

package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"rawg/pkg/api"
	"rawg/pkg/storage/mongo"

	"github.com/labstack/gommon/log"
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
	// test
	dbConnection := mongo.NewStorage("mongodb://localhost:27017", "gamesdb", "games")
	m := &api.Game{}
	filter := bson.D{{"id", 3498}}
	err = dbConnection.Get(filter, m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" model %+v", m)

	go func() {
		sig := <-sig
		fmt.Println(sig)
		done <- true
	}()

	<-done
	server.Shutdown()
}

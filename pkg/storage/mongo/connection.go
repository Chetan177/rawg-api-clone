package mongo

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	address     string
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

var (
	once            sync.Once
	mongoConnection *Connection
)

func NewStorage(address, database, collection string) *Connection {
	once.Do(func() {
		mongoConnection = &Connection{
			address: address,
		}
		//mongodb://localhost:27017
		clientOptions := options.Client().ApplyURI(address)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatalf("unable to connect to mongo db %+v", err)
		}
		mongoConnection.mongoClient = client
		mongoConnection.collection = client.Database(database).Collection(collection)
	})
	return mongoConnection
}

func (m *Connection) isConnected() bool {
	err := m.mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

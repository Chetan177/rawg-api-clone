package storage

import "go.mongodb.org/mongo-driver/bson"

type Storage interface {
	Get(bson.D, interface{}) error
}

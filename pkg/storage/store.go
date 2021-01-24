package storage

import "go.mongodb.org/mongo-driver/bson"

type Storage interface {
	Get(bson.D, interface{}) error
	GetDocumentByID(int, interface{}) error
	GetGameByID(int, interface{}) error
}

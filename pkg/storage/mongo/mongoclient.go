package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *Connection) Get(filter bson.D, result interface{}) error {
	err := m.collection.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		return err
	}
	return err
}

func (m *Connection) GetDocumentByID(id int, result interface{}) error {
	filter := bson.D{{"_id", id}}
	err := m.collection.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		return err
	}
	return err
}

func (m *Connection) GetGameByID(id int, result interface{}) error {
	filter := bson.D{{"id", id}}
	err := m.collection.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		return err
	}
	return err
}


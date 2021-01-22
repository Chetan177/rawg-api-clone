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

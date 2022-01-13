package dbcollection

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Connect(database, collection string, client *mongo.Client) (*mongo.Collection, error) {
	coll := client.Database(database).Collection(collection)

	count, err := coll.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("count error: %w", err)
	}

	if count == 0 {
		if _, err = coll.InsertOne(context.TODO(),
			bson.M{"total":123,	"actions":bson.M{}, "countries":bson.M{}}); err != nil {
			return nil, fmt.Errorf("object insert error: %w", err)
		}
	}
	return coll, nil
}
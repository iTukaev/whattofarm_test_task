package dbclient

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to MongoDB with user, password and host.
// Return *mongo.Client, if Ping is OK.
func Connect(user, password, host string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s", user, password, host))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("MongoDB client creating error: %w", err)
	}

	if err = client.Connect(context.Background()); err != nil {
		return nil, fmt.Errorf("MongoDB client connecting error: %w", err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("MongoDB client ping error: %w", err)
	}

	return client, nil
}

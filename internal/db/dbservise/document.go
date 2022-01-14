package dbservise

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (s *service) GetDocumentID() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := s.client.Database(s.database).Collection(s.collection)

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("MongoDB documents count error: %w", err)
	}

	if count > 1 {
		return fmt.Errorf("MongoDB collection have more then ONE document, check database")
	}

	if count == 0 {
		result, err := collection.InsertOne(ctx,s.data)
		if err != nil {
			return fmt.Errorf("MongoDB object insert error: %w", err)
		}
		s.data.ID = result.InsertedID.(primitive.ObjectID)
		return nil
	}

	if err := collection.FindOne(ctx, bson.M{}).Decode(s.data); err != nil {
		return fmt.Errorf("MongoDB document ID getting error: %w", err)
	}

	return nil
}

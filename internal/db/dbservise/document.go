package dbservise

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// GetDocumentID writes ObjectID of MongoDB document to service.dada.ID.
// Return <nil>, if all OK.
// If MongoDB collection is empty, GetDocumentID create new empty document
// and write ID to service.dada.ID.
// If there are more than ONE documents, GetDocumentID is returned with error
func (s *service) GetDocumentID() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := s.client.Database(s.database).Collection(s.collection)

	err := collection.FindOne(ctx, bson.M{})

	if err != nil {
		return fmt.Errorf("MongoDB documents count error: %w", err)
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

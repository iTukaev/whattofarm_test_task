package dbservise

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type InputDocument struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

// GetDocumentID writes ObjectID of MongoDB document to Service.dada.ID.
// Return <nil>, if all OK.
// If MongoDB collection is empty, GetDocumentID create new empty document
// and write ID to Service.dada.ID.
// If there are more than ONE documents, GetDocumentID is returned with error
func (s *Service) GetDocumentID() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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

	input := &InputDocument{}
	if err := collection.FindOne(ctx, bson.M{}).Decode(input); err != nil {
		return fmt.Errorf("MongoDB document ID getting error: %w", err)
	}
	s.data.ID = input.ID

	return nil
}

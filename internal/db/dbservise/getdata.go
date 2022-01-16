package dbservise

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Payload is a structure for MongoDB document
type Payload struct {
	ID primitive.ObjectID `json:"_id"`
	Total int `json:"total"`
	Actions map[string]*SubCountries `json:"actions"`
	Countries map[string]*SubActions `json:"countries"`
}

// GetData return MongoDB's document as a JSON string
// and <nil> if all OK.
// Return error, if search or marshalling are incorrect
func (s *service) GetData() (string, error) {
	collection := s.client.Database(s.database).Collection(s.collection)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	payload := &Payload{
		Actions: make(map[string]*SubCountries),
		Countries: make(map[string]*SubActions),
	}

	err := collection.FindOne(ctx, bson.M{"_id":s.data.ID}).Decode(payload)
	if err != nil {
		return "", fmt.Errorf("MongoDB document getiing error: %w", err)
	}

	result, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("MongoDB document marshalling error: %w", err)
	}

	return string(result), nil
}
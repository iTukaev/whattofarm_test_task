package dbservise

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// InputGetData is a structure for MongoDB document
type InputGetData struct {
	ID primitive.ObjectID `json:"_id"`
	Total int `json:"total"`
	Actions map[string]*TotalCounter `json:"actions"`
	Countries map[string]*TotalCounter `json:"countries"`
}

// GetData return MongoDB's document as a JSON string
// and <nil> if all OK.
// Return error, if search or marshalling are incorrect
func (s *Service) GetData() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := s.client.Database(s.database).Collection(s.collection)

	input := &InputGetData{
		ID: s.data.ID,
		Actions: make(map[string]*TotalCounter),
		Countries: make(map[string]*TotalCounter),
	}

	err := collection.FindOne(ctx, bson.M{"_id":s.data.ID}).Decode(input)
	if err != nil {
		return nil, fmt.Errorf("MongoDB document getiing error: %w", err)
	}

	result, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("MongoDB document marshalling error: %w", err)
	}

	return result, nil
}

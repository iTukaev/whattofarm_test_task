package dbservise

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Payload struct {
	ID primitive.ObjectID `json:"_id"`
	Total int `json:"total"`
	Actions map[string]*TotalCounter `json:"actions"`
	Countries map[string]*TotalCounter `json:"countries"`
}

func (s *service) GetData() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	payload := &Payload{
		Actions: make(map[string]*TotalCounter),
		Countries: make(map[string]*TotalCounter),
	}

	collection := s.client.Database(s.database).Collection(s.collection)
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

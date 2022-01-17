package dbservise

import (
	"encoding/json"
	"fmt"
)

// Payload is a structure for MongoDB document
type Payload struct {
	Total int `json:"total"`
	Actions map[string]*SubCountries `json:"actions"`
	Countries map[string]*SubActions `json:"countries"`
}

// GetData return MongoDB's document as a JSON string
// and <nil> if all OK.
// Return error, if search or marshalling are incorrect
func (s *service) GetData() ([]byte, error) {
	//collection := s.client.Database(s.database).Collection(s.collection)
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	payload := &Payload{
		Actions: make(map[string]*SubCountries),
		Countries: make(map[string]*SubActions),
	}

	//err := collection.FindOne(ctx, bson.M{"_id":s.data.ID}).Decode(payload)
	//if err != nil {
	//	return "", fmt.Errorf("MongoDB document getiing error: %w", err)
	//}

	result, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("MongoDB document marshalling error: %w", err)
	}

	return result, nil
}
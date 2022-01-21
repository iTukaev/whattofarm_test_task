package dbservise

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// Update increment total, action and country counters.
// Return <nil>, if all OK.
// Return error, if MongoDB document updating finish with error.
func (s *Service) Update(action, country string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := s.client.Database(s.database).Collection(s.collection)

	actionsCount := fmt.Sprintf("actions.%s.total", action)
	countriesCount := fmt.Sprintf("countries.%s.total", country)
	filter := bson.M{"_id": s.data.ID}
	opts := bson.M{"$inc": bson.M{
		"total":1,
		actionsCount:1,
		countriesCount:1,
	}}

	err := collection.FindOneAndUpdate(ctx, filter, opts).Err()
	if err != nil {
		return fmt.Errorf("total count update error: %w", err)
	}

	return nil
}
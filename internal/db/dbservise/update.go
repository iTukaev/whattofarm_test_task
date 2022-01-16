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
func (s *service) Update(action, country string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := s.client.Database(s.database).Collection(s.collection)

	actionsCount := fmt.Sprintf("actions.%s.total", action)
	countriesCount := fmt.Sprintf("countries.%s.total", country)
	subActionsCount := fmt.Sprintf("countries.%s.actions.%s.total", country, action)
	subCountriesCount := fmt.Sprintf("actions.%s.countries.%s.total", action, country)
	filter := bson.M{"_id": s.data.ID}
	opts := bson.M{"$inc": bson.M{
		"total":1,
		actionsCount:1,
		subCountriesCount: 1,
		countriesCount:1,
		subActionsCount: 1,
	}}

	err := collection.FindOneAndUpdate(ctx, filter, opts).Err()
	if err != nil {
		return fmt.Errorf("total count update error: %w", err)
	}

	return nil
}
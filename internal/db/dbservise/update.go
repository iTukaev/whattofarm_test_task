package dbservise

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (s *service) Update(action, country string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	actionsCount := "actions." + action + ".total"
	countriesCount := "countries." + country + ".total"
	filter := bson.M{"_id": s.data.ID}
	opts := bson.M{"$inc": bson.M{"total":1, actionsCount:1, countriesCount:1}}

	err := s.client.Database(s.database).Collection(s.collection).
		FindOneAndUpdate(ctx, filter, opts).Err()
	if err != nil {
		return fmt.Errorf("total count update error: %w", err)
	}

	return nil
}
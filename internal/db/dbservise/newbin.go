package dbservise

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// NewBin creates new MongoDB's document with timestamp s.data.TimeStamp
//
func (s *service) NewBin(timestamp time.Time) error {
	if s.data.Total > 0 {
		collection := s.client.Database(s.database).Collection(s.collection)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

		s.data.Lock()
		result, err := collection.InsertOne(ctx, s.data)
		if err != nil {
			return fmt.Errorf("MongoDB new document don't created: %w", err)
		}
		s.data = NewDBStruct()
		s.data.ID = result.InsertedID.(primitive.ObjectID)
		s.data.TimeStamp.T = uint32(timestamp.Unix())
		s.data.Unlock()
	}
	return nil
}
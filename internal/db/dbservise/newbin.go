package dbservise

import (
	"context"
	"fmt"
	"time"
)

// NewBin creates new MongoDB's document and resets data structure
func (s *service) NewBin(timestamp int) error {
	if s.data.Total > 0 {
		collection := s.client.Database(s.database).Collection(s.collection)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

		s.data.Lock()
		_, err := collection.InsertOne(ctx, s.data)
		if err != nil {
			return fmt.Errorf("MongoDB new document don't created: %w", err)
		}
		s.data = NewDBStruct(timestamp)
		s.data.Unlock()
	}
	return nil
}
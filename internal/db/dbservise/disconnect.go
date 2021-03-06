package dbservise

import (
	"context"
	"fmt"
	"time"
)

// Disconnect client from MongoDB
func (s *Service) Disconnect(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return fmt.Errorf("MongoDB client disconnection error: %w", s.client.Disconnect(ctx))
}
package dbservise

import (
	"context"
	"fmt"
	"time"
)

func (s *service) Disconnect(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return fmt.Errorf("MongoDB client disconnection error: %w", s.client.Disconnect(ctx))
}
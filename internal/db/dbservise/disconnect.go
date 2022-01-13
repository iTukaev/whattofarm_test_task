package dbservise

import (
	"context"
	"time"
)

func (s *service) Disconnect(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return  s.client.Disconnect(ctx)
}

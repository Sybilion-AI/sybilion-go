package sybilion

import (
	"context"
	"time"

	api "go.sybilion.dev/sybilion/api"
)

// WaitForecast polls GET /api/v1/forecasts/{id} until the job is settled or the context is done.
func (c *Client) WaitForecast(ctx context.Context, jobID string, poll time.Duration) (*api.ApiV1ForecastsIdGet200Response, error) {
	if poll <= 0 {
		poll = 2 * time.Second
	}
	t := time.NewTicker(poll)
	defer t.Stop()
	for {
		j, err := c.GetForecast(ctx, jobID)
		if err != nil {
			return nil, err
		}
		if j.GetSettled() {
			return j, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-t.C:
		}
	}
}

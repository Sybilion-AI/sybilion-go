package sybilion

import (
	"context"

	api "go.sybilion.dev/sybilion/api"
)

// UsagePageFunc is called for each page of GET /api/v1/usage. Return false to stop iteration.
type UsagePageFunc func(ctx context.Context, page *api.ApiV1UsageGet200Response) (continueNext bool, err error)

// ForEachUsagePage walks usage from page 1 until the callback returns false or pages are exhausted.
func (c *Client) ForEachUsagePage(ctx context.Context, sort string, order string, limit int32, fn UsagePageFunc) error {
	if limit <= 0 {
		limit = 50
	}
	page := int32(1)
	for {
		resp, _, err := c.raw.DefaultAPI.ApiV1UsageGet(ctx).Page(page).Limit(limit).Sort(sort).Order(order).Execute()
		if err != nil {
			return parseAPIError(err)
		}
		cont, err := fn(ctx, resp)
		if err != nil {
			return err
		}
		if !cont {
			return nil
		}
		pg := resp.GetPagination()
		if int64(page) >= pg.GetTotalPages() {
			return nil
		}
		page++
	}
}

// JobsPageFunc is called for each page of GET /api/v1/jobs.
type JobsPageFunc func(ctx context.Context, page *api.ApiV1JobsGet200Response) (continueNext bool, err error)

// ForEachJobsPage walks job pages from page 1 until the callback returns false or pages are exhausted.
func (c *Client) ForEachJobsPage(ctx context.Context, sort string, order string, limit int32, fn JobsPageFunc) error {
	if limit <= 0 {
		limit = 50
	}
	page := int32(1)
	for {
		resp, _, err := c.raw.DefaultAPI.ApiV1JobsGet(ctx).Page(page).Limit(limit).Sort(sort).Order(order).Execute()
		if err != nil {
			return parseAPIError(err)
		}
		cont, err := fn(ctx, resp)
		if err != nil {
			return err
		}
		if !cont {
			return nil
		}
		pg := resp.GetPagination()
		if int64(page) >= pg.GetTotalPages() {
			return nil
		}
		page++
	}
}

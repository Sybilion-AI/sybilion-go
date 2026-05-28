package sybilion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	api "go.sybilion.dev/sybilion/api"
)

// AlertsRequest is the request body for POST /api/v1/alerts.
type AlertsRequest struct {
	Metadata        api.TimeseriesMetadata `json:"metadata"`
	ContextEnriched bool                   `json:"context_enriched"`
	Filters         *api.Filters           `json:"filters,omitempty"`
	DateFrom        string                 `json:"date_from,omitempty"`
	DateTo          string                 `json:"date_to,omitempty"`
}

// NewsItem is a news article attached to an AlertItem.
type NewsItem struct {
	Category    string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	PublishedAt string `json:"published_at,omitempty"`
	SourceName  string `json:"source_name,omitempty"`
	Title       string `json:"title,omitempty"`
	Trending    bool   `json:"trending,omitempty"`
	URL         string `json:"url,omitempty"`
}

// AlertItem is a single alert returned by POST /api/v1/alerts.
type AlertItem struct {
	Name      string     `json:"name,omitempty"`
	PctChange float64    `json:"pct_change,omitempty"`
	Trending  bool       `json:"trending,omitempty"`
	News      []NewsItem `json:"news,omitempty"`
}

type alertsResponse struct {
	Alerts []AlertItem `json:"alerts"`
}

// GetAlerts retrieves anomaly alerts for the supplied timeseries metadata (synchronous, billed).
func (c *Client) GetAlerts(ctx context.Context, req AlertsRequest) ([]AlertItem, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("sybilion: marshal alerts request: %w", err)
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/api/v1/alerts", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("sybilion: new alerts request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.token)
	}
	httpClient := c.raw.GetConfig().HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("sybilion: alerts http: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("sybilion: read alerts response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		var body struct {
			Error string `json:"error"`
		}
		if json.Unmarshal(data, &body) == nil && body.Error != "" {
			return nil, fmt.Errorf("%s", body.Error)
		}
		return nil, fmt.Errorf("%s", resp.Status)
	}
	var out alertsResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("sybilion: decode alerts response: %w", err)
	}
	return out.Alerts, nil
}

// Package sybilion is the official Go client for the Sybilion API.
//
// The hand-written wrapper provides ergonomic helpers (Bearer auth, forecast
// polling, pagination iterators) on top of the OpenAPI-generated client at
// go.sybilion.dev/sybilion/api. Most users only need:
//
//	c := sybilion.New(sybilion.Options{Token: os.Getenv("SYBILION_API_TOKEN")})
//	me, err := c.Me(ctx)
//
// Token can also be read from the SYBILION_API_TOKEN environment variable:
//
//	// export SYBILION_API_TOKEN=sk_ops_...
//	c := sybilion.New(sybilion.Options{})
//
// See https://sybilion.dev/docs/ for full guides and feature walkthroughs.
package sybilion

import (
	"context"
	"net/http"
	"os"
	"time"

	api "go.sybilion.dev/sybilion/api"
)

// Options configures the HTTP client and authentication.
type Options struct {
	// BaseURL is the API origin, e.g. https://api.example.com (no trailing slash required).
	// Empty BaseURL uses SYBILION_API_BASE_URL (if set), else DefaultPublicAPIBaseURL (defaults_gen.go).
	BaseURL string
	// Token is sent as Authorization: Bearer <token> (an API key sk_ops_... or a dashboard session token).
	// When empty, SYBILION_API_TOKEN is read from the environment.
	Token string
	// HTTPClient overrides the default net/http client (timeouts, transport, etc.).
	HTTPClient *http.Client
	// UserAgent overrides the default User-Agent header.
	UserAgent string
}

// Client is a thin wrapper around the openapi-generator DefaultAPI client.
type Client struct {
	raw     *api.APIClient
	baseURL string
	token   string
}

// New constructs a Client. Token resolution order: Options.Token, SYBILION_API_TOKEN env var.
// Base URL resolution: Options.BaseURL, SYBILION_API_BASE_URL env var, DefaultPublicAPIBaseURL.
func New(opts Options) *Client {
	cfg := api.NewConfiguration()
	base := resolveAPIBaseURL(opts.BaseURL)
	cfg.Servers = api.ServerConfigurations{{URL: base}}
	if opts.HTTPClient != nil {
		cfg.HTTPClient = opts.HTTPClient
	} else {
		cfg.HTTPClient = &http.Client{Timeout: 60 * time.Second}
	}
	token := opts.Token
	if token == "" {
		token = os.Getenv(EnvSybilionAPIToken)
	}
	if token != "" {
		cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	}
	if opts.UserAgent != "" {
		cfg.UserAgent = opts.UserAgent
	}
	return &Client{raw: api.NewAPIClient(cfg), baseURL: base, token: token}
}

// DefaultAPI returns the DefaultAPIService for advanced or paginated calls.
func (c *Client) DefaultAPI() *api.DefaultAPIService {
	return c.raw.DefaultAPI
}

// ── Account ───────────────────────────────────────────────────────────────────

// Me returns the authenticated account's info.
func (c *Client) Me(ctx context.Context) (*api.MeResponse, error) {
	resp, _, err := c.raw.DefaultAPI.ApiV1MeGet(ctx).Execute()
	return resp, parseAPIError(err)
}

// ── Catalog ───────────────────────────────────────────────────────────────────

// ListCategories lists available thematic categories.
func (c *Client) ListCategories(ctx context.Context) (*api.CatalogListResponse, error) {
	resp, _, err := c.raw.DefaultAPI.ApiV1CategoriesGet(ctx).Execute()
	return resp, parseAPIError(err)
}

// ListRegions lists available geographic regions.
func (c *Client) ListRegions(ctx context.Context) (*api.RegionListResponse, error) {
	resp, _, err := c.raw.DefaultAPI.ApiV1RegionsGet(ctx).Execute()
	return resp, parseAPIError(err)
}

// ── Forecasts ─────────────────────────────────────────────────────────────────

// SubmitForecast submits an async forecast job.
func (c *Client) SubmitForecast(ctx context.Context, req api.ForecastRequestV1) (*api.ApiV1ForecastsPost202Response, error) {
	resp, _, err := c.raw.DefaultAPI.ApiV1ForecastsPost(ctx).ForecastRequestV1(req).Execute()
	return resp, parseAPIError(err)
}

// GetForecast fetches the current status and metadata of a forecast job.
func (c *Client) GetForecast(ctx context.Context, id string) (*api.ApiV1ForecastsIdGet200Response, error) {
	resp, _, err := c.raw.DefaultAPI.ApiV1ForecastsIdGet(ctx, id).Execute()
	return resp, parseAPIError(err)
}

// GetForecastArtifact downloads a forecast artifact by name.
func (c *Client) GetForecastArtifact(ctx context.Context, id, name string) (*os.File, error) {
	resp, _, err := c.raw.DefaultAPI.ApiV1ForecastsIdArtifactsNameGet(ctx, id, name).Execute()
	return resp, parseAPIError(err)
}

// ── Drivers ───────────────────────────────────────────────────────────────────

// GetDrivers retrieves drivers ranked by explanatory power (synchronous, billed).
func (c *Client) GetDrivers(ctx context.Context, req api.RecommendRequestV1) (*http.Response, error) {
	resp, err := c.raw.DefaultAPI.ApiV1DriversPost(ctx).RecommendRequestV1(req).Execute()
	return resp, parseAPIError(err)
}

// ── Forecast helpers ──────────────────────────────────────────────────────────

// Forecasts groups forecast-related helpers.
type Forecasts struct {
	client *Client
}

// Forecasts returns forecast helpers including WaitForecast.
func (c *Client) Forecasts() *Forecasts {
	return &Forecasts{client: c}
}

// Wait polls GET /api/v1/forecasts/{id} until settled or context cancellation.
func (f *Forecasts) Wait(ctx context.Context, jobID string, poll time.Duration) (*api.ApiV1ForecastsIdGet200Response, error) {
	return f.client.WaitForecast(ctx, jobID, poll)
}

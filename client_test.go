package sybilion

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	api "go.sybilion.dev/sybilion/api"
)

func TestClient_AuthHeaderOnMe(t *testing.T) {
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/me" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		gotAuth = r.Header.Get("Authorization")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"user_id":             "00000000-0000-0000-0000-000000000001",
			"balance_eur_cents":   0,
			"available_eur_cents": 0,
			"api_usage_tier":      0,
			"lifetime_paid_cents": 0,
			"payment_count":       0,
			"has_ever_paid":       false,
			"euro_tranches":       []any{},
			"auto_recharge": map[string]any{
				"enabled":             false,
				"below_eur_cents":     0,
				"target_eur_cents":    0,
				"monthly_cap_cents":   0,
				"meter_cents":         0,
				"meter_month":         nil,
				"has_stripe_customer": false,
			},
		})
	}))
	defer srv.Close()

	c := New(Options{BaseURL: srv.URL, Token: "sk_ops_test"})
	_, err := c.Me(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if gotAuth != "Bearer sk_ops_test" {
		t.Fatalf("Authorization = %q, want Bearer sk_ops_test", gotAuth)
	}
}

func TestClient_TokenFromEnv(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/me" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"user_id":             "00000000-0000-0000-0000-000000000001",
			"balance_eur_cents":   0,
			"available_eur_cents": 0,
			"api_usage_tier":      0,
			"lifetime_paid_cents": 0,
			"payment_count":       0,
			"has_ever_paid":       false,
			"euro_tranches":       []any{},
			"auto_recharge": map[string]any{
				"enabled":             false,
				"below_eur_cents":     0,
				"target_eur_cents":    0,
				"monthly_cap_cents":   0,
				"meter_cents":         0,
				"meter_month":         nil,
				"has_stripe_customer": false,
			},
		})
	}))
	defer srv.Close()

	t.Setenv(EnvSybilionAPIToken, "sk_ops_env")
	c := New(Options{BaseURL: srv.URL})
	_, err := c.Me(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_TokenFromEnv_NotSet(t *testing.T) {
	os.Unsetenv(EnvSybilionAPIToken)
	c := New(Options{BaseURL: "http://localhost"})
	// No token set — client is constructed but requests will fail auth; just check it's created.
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestWaitForecast_Settled(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/forecasts/00000000-0000-0000-0000-000000000002" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		calls++
		settled := calls >= 2
		_ = json.NewEncoder(w).Encode(map[string]any{
			"job_id":          "00000000-0000-0000-0000-000000000002",
			"status":          "completed",
			"pipeline_type":   "forecast",
			"created_at":      time.Now().UTC().Format(time.RFC3339),
			"settled":         settled,
			"settled_at":      time.Now().UTC().Format(time.RFC3339),
			"eur_cents_final": 1,
		})
	}))
	defer srv.Close()

	c := New(Options{BaseURL: srv.URL, Token: "x"})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	j, err := c.Forecasts().Wait(ctx, "00000000-0000-0000-0000-000000000002", 10*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if !j.GetSettled() {
		t.Fatal("expected settled job")
	}
	if calls < 2 {
		t.Fatalf("expected at least 2 polls, got %d", calls)
	}
}

func TestParseAPIError_ExtractsErrorField(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]any{"error": "invalid token"})
	}))
	defer srv.Close()

	c := New(Options{BaseURL: srv.URL, Token: "bad"})
	_, err := c.Me(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid token" {
		t.Fatalf("want %q, got %q", "invalid token", err.Error())
	}
}

func TestParseAPIError_FallsBackToOriginal(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer srv.Close()

	c := New(Options{BaseURL: srv.URL, Token: "bad"})
	_, err := c.Me(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() == "" {
		t.Fatal("expected non-empty error message")
	}
}

func TestForecastRequestV1_AuxTimeseriesRoundTrip(t *testing.T) {
	aux := []map[string]float32{
		{"2015-01-01": 1.2, "2015-02-01": 1.4, "2015-03-01": 1.5},
		{"2015-01-01": 9.0, "2015-02-01": 8.7, "2015-03-01": 8.9},
	}
	req := api.ForecastRequestV1{
		PipelineVersion:    "v1",
		Frequency:          "monthly",
		RecencyFactor:      0.5,
		TimeseriesMetadata: api.TimeseriesMetadata{Title: "Monthly Brent Crude Oil Price Index"},
		Timeseries:         map[string]float32{"2015-01-01": 47.8, "2015-02-01": 58.1, "2015-03-01": 61.0},
		SoftHorizon:        api.PtrInt32(12),
		AuxTimeseries:      aux,
	}

	b, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if _, ok := raw["aux_timeseries"]; !ok {
		t.Fatalf("aux_timeseries missing from serialized request: %s", b)
	}

	var back api.ForecastRequestV1
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatalf("round-trip unmarshal: %v", err)
	}
	if len(back.AuxTimeseries) != 2 {
		t.Fatalf("expected 2 aux series, got %d", len(back.AuxTimeseries))
	}
}

func TestForecastRequestV1_AuxTimeseriesOmitted(t *testing.T) {
	req := api.ForecastRequestV1{
		PipelineVersion:    "v1",
		Frequency:          "monthly",
		RecencyFactor:      0.5,
		TimeseriesMetadata: api.TimeseriesMetadata{Title: "Monthly Brent Crude Oil Price Index"},
		Timeseries:         map[string]float32{"2015-01-01": 47.8},
		SoftHorizon:        api.PtrInt32(12),
	}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if bytesContains(b, "aux_timeseries") {
		t.Fatalf("aux_timeseries should be omitted when nil: %s", b)
	}
}

func bytesContains(b []byte, sub string) bool {
	return strings.Contains(string(b), sub)
}

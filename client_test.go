package sybilion

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
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

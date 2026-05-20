package sybilion

import "testing"

func TestResolveAPIBaseURL(t *testing.T) {
	t.Run("explicit wins over env", func(t *testing.T) {
		t.Setenv(EnvSybilionAPIBaseURL, "https://env.example")
		got := resolveAPIBaseURL("https://explicit.example/")
		want := "https://explicit.example"
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	})
	t.Run("env when explicit empty", func(t *testing.T) {
		t.Setenv(EnvSybilionAPIBaseURL, "https://from-env.example/")
		got := resolveAPIBaseURL("")
		want := "https://from-env.example"
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	})
	t.Run("default when env unset", func(t *testing.T) {
		t.Setenv(EnvSybilionAPIBaseURL, "")
		got := resolveAPIBaseURL("")
		if got != DefaultPublicAPIBaseURL {
			t.Fatalf("got %q, want %q", got, DefaultPublicAPIBaseURL)
		}
	})
}

package deezerauth

import (
	"reflect"
	"testing"
)

func TestNew_UsesEnvVars(t *testing.T) {
	// Arrange
	id := "env-client-id"
	secret := "env-client-secret"

	t.Setenv("DEEZER_ID", id)
	t.Setenv("DEEZER_SECRET", secret)

	// Act
	a := New()

	// Assert
	if a == nil || a.config == nil {
		t.Fatalf("New() returned nil or nil config")
	}

	if got, want := a.config.ClientID, id; got != want {
		t.Errorf("ClientID = %q, want %q", got, want)
	}

	if got, want := a.config.ClientSecret, secret; got != want {
		t.Errorf("ClientSecret = %q, want %q", got, want)
	}

	// Endpoint URLs should be set to constants
	if got, want := a.config.Endpoint.AuthURL, AuthURL; got != want {
		t.Errorf("Endpoint.AuthURL = %q, want %q", got, want)
	}
	if got, want := a.config.Endpoint.TokenURL, TokenURL; got != want {
		t.Errorf("Endpoint.TokenURL = %q, want %q", got, want)
	}

	// By default, optional fields should be empty
	if a.config.RedirectURL != "" {
		t.Errorf("RedirectURL = %q, want empty", a.config.RedirectURL)
	}
	if len(a.config.Scopes) != 0 {
		t.Errorf("Scopes = %v, want empty", a.config.Scopes)
	}
}

func TestNew_WithOptionsOverride(t *testing.T) {
	// Arrange env defaults that should be overridden
	t.Setenv("DEEZER_ID", "env-id")
	t.Setenv("DEEZER_SECRET", "env-secret")

	wantID := "opt-id"
	wantSecret := "opt-secret"
	wantRedirect := "https://example.com/callback"
	wantScopes := []string{"basic_access", "manage_library"}

	// Act
	a := New(
		WithClientID(wantID),
		WithClientSecret(wantSecret),
		WithRedirectURL(wantRedirect),
		WithScopes(wantScopes...),
	)

	// Assert
	if a == nil || a.config == nil {
		t.Fatalf("New() returned nil or nil config")
	}

	if got := a.config.ClientID; got != wantID {
		t.Errorf("ClientID = %q, want %q", got, wantID)
	}
	if got := a.config.ClientSecret; got != wantSecret {
		t.Errorf("ClientSecret = %q, want %q", got, wantSecret)
	}
	if got := a.config.RedirectURL; got != wantRedirect {
		t.Errorf("RedirectURL = %q, want %q", got, wantRedirect)
	}
	if got := a.config.Scopes; !reflect.DeepEqual(got, wantScopes) {
		t.Errorf("Scopes = %v, want %v", got, wantScopes)
	}

	// Endpoint constants should remain intact
	if a.config.Endpoint.AuthURL != AuthURL {
		t.Errorf("Endpoint.AuthURL = %q, want %q", a.config.Endpoint.AuthURL, AuthURL)
	}
	if a.config.Endpoint.TokenURL != TokenURL {
		t.Errorf("Endpoint.TokenURL = %q, want %q", a.config.Endpoint.TokenURL, TokenURL)
	}
}

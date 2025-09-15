package deezerauth

import (
	"os"

	"golang.org/x/oauth2"
)

const (
	AuthURL  = "https://connect.deezer.com/oauth/auth.php"
	TokenURL = "https://connect.deezer.com/oauth/access_token.php"
)

type Authenticator struct {
	config *oauth2.Config
}

type AuthenticatorOption func(a *Authenticator)

// WithClientID allows a client ID to be specified. Without this the value of the DEEZER_ID environment
// variable will be used.
func WithClientID(id string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.ClientID = id
	}
}

// WithClientSecret allows a client secret to be specified. Without this the value of the DEEZER_SECRET environment
// variable will be used.
func WithClientSecret(secret string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.ClientSecret = secret
	}
}

// WithScopes configures the oauth scopes that the client should request.
func WithScopes(scopes ...string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.Scopes = scopes
	}
}

// WithRedirectURL configures a redirect url for oauth flows. It must exactly match one of the
// URLs specified in your Deezer developer account.
func WithRedirectURL(url string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.RedirectURL = url
	}
}

func New(opts ...AuthenticatorOption) *Authenticator {
	cfg := &oauth2.Config{
		ClientID:     os.Getenv("DEEZER_ID"),
		ClientSecret: os.Getenv("DEEZER_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthURL,
			TokenURL: TokenURL,
		},
	}

	a := &Authenticator{
		config: cfg,
	}

	for _, opt := range opts {
		opt(a)
	}

	return a
}

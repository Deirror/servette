package oauth

import (
	"context"

	"golang.org/x/oauth2"
)

type Profile struct {
	Email      string `json:"email"`
	ExternalID string `json:"id"`
	Provider   string
}

type Provider interface {
	AuthCodeURL(state string) (string, error)
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	FetchProfile(ctx context.Context, token *oauth2.Token) (*Profile, error)
}

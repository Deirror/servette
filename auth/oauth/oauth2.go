package oauth

import (
	"golang.org/x/oauth2"
)

func NewOAuth2Config(cfg *Config, scopes []string, urls oauth2.Endpoint) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       scopes,
		Endpoint:     urls,
	}
}

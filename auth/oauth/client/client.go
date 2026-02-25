package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/Deirror/servette/auth/oauth"
	"github.com/Deirror/servette/auth/oauth/facebook"
	"github.com/Deirror/servette/auth/oauth/github"
	"github.com/Deirror/servette/auth/oauth/google"
	"github.com/Deirror/servette/encoding/json"
	"golang.org/x/oauth2"
)

type Client struct {
	conf     *oauth2.Config
	provider string
}

func NewClient(cfg *oauth.Config, provider string) (*Client, error) {
	provider = strings.ToLower(provider)

	var oauthCfg *oauth2.Config
	switch provider {
	case "google":
		oauthCfg = google.NewOAuth2Config(cfg)
	case "facebook":
		oauthCfg = facebook.NewOAuth2Config(cfg)
	case "github":
		oauthCfg = github.NewOAuth2Config(cfg)
	default:
		return nil, fmt.Errorf("unknown provider passed: %s", provider)
	}

	return &Client{
		conf:     oauthCfg,
		provider: provider,
	}, nil
}

func (c *Client) AuthCodeURL(state string) (string, error) {
	return c.conf.AuthCodeURL(state), nil
}

func (c *Client) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return c.conf.Exchange(ctx, code)
}

func (c *Client) FetchProfile(ctx context.Context, token *oauth2.Token) (*oauth.Profile, error) {
	url := oauth.UserInfoURLs[c.provider]

	client := c.conf.Client(ctx, token)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var profile oauth.Profile
	if err = json.DecodeInto(resp.Body, &profile); err != nil {
		return nil, err
	}

	profile.Provider = c.provider

	return &profile, nil
}

// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package payment

// Config holds configuration details for payment processing.
type Config struct {
	APIKey        string // Public API key for the payment provider
	SecretKey     string // Secret key for authenticating requests
	WebhookURL    string // URL to receive payment provider webhook callbacks
	WebhookSecret string // Secret used to verify webhook authenticity
}

func NewConfig(apiKey, secretKey, url, webhookSecret string) *Config {
	return &Config{
		APIKey:        apiKey,
		SecretKey:     secretKey,
		WebhookURL:    url,
		WebhookSecret: webhookSecret,
	}
}

// WithAPIKey sets the API key and returns the updated config.
func (c *Config) WithAPIKey(key string) *Config {
	c.APIKey = key
	return c
}

// WithSecretKey sets the secret key and returns the updated config.
func (c *Config) WithSecretKey(secret string) *Config {
	c.SecretKey = secret
	return c
}

// WithWebhookURL sets the webhook URL and returns the updated config.
func (c *Config) WithWebhookURL(url string) *Config {
	c.WebhookURL = url
	return c
}

// WithWebhookSecret sets the webhook secret and returns the updated config.
func (c *Config) WithWebhookSecret(secret string) *Config {
	c.WebhookSecret = secret
	return c
}

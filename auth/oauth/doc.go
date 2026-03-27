// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package oauth provides a lightweight abstraction over OAuth2 authentication
// flows for multiple providers (organizations).
//
// It wraps golang.org/x/oauth2 to simplify:
//
//   - Generating authorization URLs
//   - Exchanging authorization codes for tokens
//   - Fetching user profile information from provider-specific endpoints
//
// The package is designed around a provider (Org) concept, where each
// organization has predefined OAuth2 configuration and user info endpoints.
//
// # Core Concepts
//
// Client
//
//	Client is the main entry point for performing OAuth2 operations.
//	It holds a configured oauth2.Config and a selected provider (Org).
//
// Org
//
//	Org represents a specific OAuth provider (e.g. Google, GitHub, etc.).
//	Each Org maps to predefined configuration such as:
//	  - Auth and token URLs
//	  - Scopes
//	  - User info endpoint
//
// Profile
//
//	Profile represents the user data returned from the provider's
//	user info endpoint. It is populated via JSON decoding.
//
// # Flow
//
// Typical OAuth2 flow using this package:
//
//  1. Create a Client with configuration and provider (Org)
//  2. Redirect user to AuthCodeURL
//  3. Receive authorization code from callback
//  4. Exchange code for token
//  5. Fetch user profile using the token
//
// # Example
//
//	client, _ := oauth.NewClient(cfg, oauth.Google)
//
//	url, _ := client.AuthCodeURL("state")
//	// redirect user to url
//
//	token, _ := client.Exchange(ctx, code)
//
//	profile, _ := client.FetchProfile(ctx, token)
//
// # Notes
//
//   - This package assumes provider-specific configuration is defined
//     in an internal props map.
//   - JSON decoding is delegated to a custom encoding/json package.
//   - The HTTP client is derived from oauth2.Config, ensuring proper
//     token usage and authorization headers.
package oauth

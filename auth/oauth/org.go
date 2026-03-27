// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package oauth

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Org string

const (
	Google   Org = "google"
	Facebook Org = "facebook"
	GitHub   Org = "github"
)

func (o Org) IsGoogle() bool {
	return o == Google
}

func (o Org) IsFacebook() bool {
	return o == Facebook
}

func (o Org) IsGithub() bool {
	return o == GitHub
}

func ParseOrg(v string) (Org, error) {
	switch v {
	case string(Google):
		return Google, nil
	case string(Facebook):
		return Facebook, nil
	case string(GitHub):
		return GitHub, nil
	default:
		return "", fmt.Errorf("invalid org: %s", v)
	}
}

type Props struct {
	Scopes      []string
	UserInfoURL string
	Endpoint    oauth2.Endpoint
}

type OrgProps = map[Org]Props

var props OrgProps = OrgProps{
	Google: {
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		UserInfoURL: "https://www.googleapis.com/oauth2/v2/userinfo",
		Endpoint:    google.Endpoint,
	},
	Facebook: {
		Scopes:      []string{"email", "public_profile"},
		UserInfoURL: "https://graph.facebook.com/me?fields=id,name,email,picture",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/v17.0/dialog/oauth",
			TokenURL: "https://graph.facebook.com/v17.0/oauth/access_token",
		},
	},
	GitHub: {
		Scopes:      []string{"read:user", "user:email"},
		UserInfoURL: "https://api.github.com/user",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	},
}

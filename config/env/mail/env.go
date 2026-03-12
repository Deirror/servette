// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package mail

import (
	"github.com/Deirror/servette/config"
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/domain/mail"
	"github.com/Deirror/servette/env"
)

type MultiConfig = config.MultiConfig[mail.Config]

var suffixes = []string{
	"MAIL_HOST",
	"MAIL_PORT",
	"MAIL_USERNAME",
	"MAIL_PASSWORD",
	"MAIL_FROM",
}

// LoadConfig loads Mail Config values from environment variables,
// supporting an optional prefix.
func LoadConfig(prefix ...string) (*mail.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	host, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	port, err := env.Get(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	username, err := env.Get(pfx + suffixes[2])
	if err != nil {
		return nil, err
	}

	password, err := env.Get(pfx + suffixes[3])
	if err != nil {
		return nil, err
	}

	from, err := env.Get(pfx + suffixes[4])
	if err != nil {
		return nil, err
	}

	return mail.NewConfig(host, port, username, password, from), nil
}

// LoadMultiConfig loads multiple Configs by scanning env vars with mailer suffixes.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}

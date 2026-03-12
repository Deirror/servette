// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package payment

import (
	"github.com/Deirror/servette/config"
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/domain/payment"
	"github.com/Deirror/servette/env"
)

type MultiConfig = config.MultiConfig[payment.Config]

var suffixes = []string{
	"PAYMENT_API_KEY",
	"PAYMENT_SECRET_KEY",
	"PAYMENT_WEBHOOK_URL",
	"PAYMENT_WEBHOOK_SECRET",
}

// LoadConfig loads Payment Config from environment variables,
// optionally with a prefix.
func LoadConfig(prefix ...string) (*payment.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	apiKey, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	secretKey, err := env.Get(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	webhookURL, err := env.Get(pfx + suffixes[2])
	if err != nil {
		return nil, err
	}

	webhookSecret, err := env.Get(pfx + suffixes[3])
	if err != nil {
		return nil, err
	}

	return payment.NewConfig(apiKey, secretKey, webhookURL, webhookSecret), nil
}

// LoadMultiConfig loads multiple Payment Configs by scanning env vars with payment suffixes.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}

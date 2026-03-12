// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// Loads environment variables from the given files.
// It is intented to be used in development mode.
// If in dev mode, make sure to load the env vars first, then call other funcs.
func Load(filenames ...string) error {
	if err := godotenv.Load(filenames...); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}

// A wrapper func around os.Getenv, but handles error with more text.
func Get(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return val, nil
}

// A wrapper func around godotenv read func, handling errors more precisely.
func GetAll(filenames ...string) (map[string]string, error) {
	if len(filenames) == 0 {
		filenames = []string{".env"}
	}

	for _, filename := range filenames {
		if fileExists(filename) {
			kvps, err := godotenv.Read(filenames...)
			if err != nil {
				return nil, fmt.Errorf("failed to read .env: %w", err)
			}
			return kvps, nil
		}
	}

	envVars := make(map[string]string)
	for _, entry := range os.Environ() {
		parts := strings.SplitN(entry, "=", 2)
		if len(parts) == 2 {
			envVars[parts[0]] = parts[1]
		}
	}

	if len(envVars) == 0 {
		return nil, errors.New("no environment variables found (neither .env nor OS)")
	}

	return envVars, nil
}

// Same as Get, but with default value.
func GetValOrDefault(key, defaultVal string) string {
	val, err := Get(key)
	if err != nil {
		return defaultVal
	}
	return val
}

// Gets key's value, calling GetEnv and parses to bool.
// Assumes it is in a specific format - text or a binary representation.
func ParseBool(key string) (bool, error) {
	val, err := Get(key)
	if err != nil {
		return false, err
	}

	switch strings.ToLower(val) {
	case "true", "1", "yes", "y":
		return true, nil
	case "false", "0", "no", "n":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean value for %s: %s", key, val)
	}
}

// Gets and parses key's value to int.
func ParseInt(key string) (int, error) {
	val, err := Get(key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("invalid integer value for %s: %s", key, val)
	}
	return i, nil
}

// Gets and parses key's value to duration.
func ParseTimeDuration(key string) (time.Duration, error) {
	val, err := Get(key)
	if err != nil {
		return 0, err
	}

	dur, err := time.ParseDuration(val)
	if err != nil {
		return 0, fmt.Errorf("invalid duration value for %s: %s", key, val)
	}

	return dur, nil
}

// checks if file is valid and not dir.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

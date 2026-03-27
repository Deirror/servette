// Package translation provides utilities for managing multilingual
// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style

// license that can be found in the LICENSE file.
// content and performing runtime translation lookups in Go applications.
//
// It includes structures and helpers to:
//   - Load translation files (.json) into language bundles
//   - Flatten nested translation keys for easy access
//   - Resolve the appropriate language dynamically
//   - Retrieve translated strings via a simple interface
//
// # Core Types
//
//   - Translator: Combines a translation Bundle and a language Resolver
//     to provide context-aware translations.
//
//   - Bundle: Stores translations for multiple languages. Each language
//     maps keys to translated strings.
//
// - I18n: Provides convenient access to translations for a single language.
//
// # Loading Translations
//
// Use LoadBundle to load all JSON translation files from a directory:
//
//	bundle, err := translation.LoadBundle("translations")
//
// Files should be named by language code, e.g.:
//
//	en.json
//	es.json
//
// Nested JSON objects are flattened into dot-separated keys:
//
//	{
//	  "greeting": {
//	    "hello": "Hello"
//	  }
//	}
//
// becomes "greeting.hello": "Hello"
//
// # Using a Translator
//
// Create a Translator using Emplace, specifying the default and supported languages:
//
//	tr, err := translation.Emplace("translations", "en", "en", "es")
//
// Retrieve translated strings using the Translator or I18n:
//
//	i18n := bundle.ForLang("es")
//	fmt.Println(i18n.T("greeting.hello")) // Output: "Hola"
//
// # Notes
//
//   - If a translation key is missing, I18n returns the key itself
//   - Flattening allows deep JSON structures to be accessed via simple string keys
//   - Translator integrates a Resolver to dynamically select the correct language
//   - JSON decoding uses the servette/json package for type-safe operations
//
// This package is intended for applications needing runtime, file-based
// internationalization (i18n) and consistent language management.
package translation

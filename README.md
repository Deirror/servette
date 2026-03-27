# servette

`servette` is a lightweight Go backend toolkit for building web services with less boilerplate. It collects common helpers for application lifecycle management, configuration loading, environment access, structured logging, HTTP transport, JSON encoding, JWT handling, password hashing, localization, and reusable service clients.

## What it is for

The package is designed to help you wire together the repetitive parts of a Go service without forcing a large framework on top of your codebase. The main idea is to keep the building blocks small, composable, and ready to use.

## Highlights

- application lifecycle orchestration with `app`
- environment-driven configuration with `config` and `config/env`
- password hashing with `crypto`
- JWT token generation, validation, and cookie helpers with `auth/jwt`
- OAuth client helpers with `auth/oauth`
- JSON helpers for HTTP handlers and typed decoding with `encoding/json`
- structured logging helpers with `logger`
- URL and project-path helpers with `path`
- toast notifications through HTTP headers with `toast`
- translation bundles and language resolution with `translation`
- transport and HTTP protocol helpers with `transport`
- ready-to-use external service clients under `domain`

## Repository layout

- `app` — application lifecycle runner abstraction
- `auth` — authentication helpers such as JWT and OAuth
- `config` — configuration loading and composition
- `crypto` — password hashing and verification
- `domain` — ready-to-use service clients for external systems
- `encoding/json` — JSON encode/decode helpers
- `env` — environment variable loading and parsing
- `logger` — structured logging helpers
- `path` — project root and path utilities
- `toast` — server-triggered toast notifications
- `translation` — bundles, resolvers, and i18n helpers
- `transport` — HTTP transport and protocol helpers

## Package overview

### `app`
The `app` package provides a small runner abstraction for coordinating long-running components. It is built around a `Runner` interface with `Start` and `Shutdown` methods, so you can run servers, workers, or any other background processes through one entry point.

### `config` and `config/env`
These packages load application configuration from environment variables. They are designed around prefix-based configuration so the same app can be deployed in multiple environments or instances without changing code.

### `env`
The `env` package wraps common environment-variable tasks such as loading `.env` files, reading values, and parsing booleans, integers, and durations.

### `auth`
Authentication helpers are split into focused packages:
- `jwt` for token creation, validation, and cookie management
- `oauth` for OAuth2 authorization flows and provider profile fetching

### `crypto`
The `crypto` package provides bcrypt password hashing and password verification helpers.

### `encoding/json`
This package adds convenience wrappers around Go’s standard JSON support, including HTTP response helpers and generic decode functions.

### `logger`
The `logger` package wraps `log/slog` with environment-aware initialization and helper functions for function-call logging with request IDs and optional timing.

### `path`
Path utilities help with discovering the project root and building URL-style paths.

### `toast`
The `toast` package sends JSON payloads through the `HX-Trigger` header so the frontend can show success, error, info, or warning toasts.

### `translation`
The translation packages load language bundles from JSON files, resolve the active language, and return translated strings with a simple `T(key)` lookup.

### `transport`
Transport helpers define HTTP-oriented request/response conventions and protocol-level utilities for passing data through handlers.

### `domain`
The `domain` tree contains preconfigured clients for external systems such as databases, caches, and other integrations.

## Example

```go
appRunner := app.New(logger, httpServer, backgroundWorker)

if err := appRunner.Run(ctx); err != nil {
    logger.Error("app stopped", "err", err)
}
```

```go
cfg, err := appenv.LoadConfig("WEB")
if err != nil {
    return err
}

log := logger.New(cfg.Mode)
```

```go
hash, err := crypto.HashPassword("my-secret-password")
if err != nil {
    return err
}
```

```go
toast.Success(w, "Saved successfully")
```

## Requirements

- Go `1.25.4`

## Installation

```bash
go get github.com/Deirror/servette
```

## Philosophy

`servette` aims to remove repetitive setup code while staying close to standard Go patterns. Each package is intentionally small, focused, and easy to compose with the rest of your codebase.

## Contributing

Contributions are welcome. Please open an issue or pull request.

## License

MIT

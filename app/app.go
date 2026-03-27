// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package app

import (
	"context"
	"log/slog"
	"sync"
)

// Runner represents a long-running component managed by the App.
// Start should block until the component stops or the context is canceled.
// Shutdown should attempt to gracefully stop the component.
type Runner interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

// App coordinates the lifecycle of multiple Runner instances.
type App struct {
	log *slog.Logger

	runners []Runner
}

func New(log *slog.Logger, runners ...Runner) *App {
	return &App{
		log:     log,
		runners: runners,
	}
}

// Run starts all runners and blocks until either:
//   - the provided context is canceled, or
//   - any runner returns an error
//
// After that, it triggers a graceful shutdown of all runners and waits
// for them to finish before returning.
//
// Note: The first runner error will cause the app to begin shutdown.
func (a *App) Run(ctx context.Context) error {
	var wg sync.WaitGroup

	// Shared error channel
	errCh := make(chan error, len(a.runners))

	a.log.Info("app starting")

	// Start runners
	a.start(ctx, &wg, errCh)

	// Wait for stop signal or runner error
	select {
	case <-ctx.Done():
		a.log.Info("app shutdown requested")
	case err := <-errCh:
		a.log.Error("app stopping due to runner error", slog.Any(Error, err))
	}

	a.log.Info("app shutting down")

	// Stop runners
	a.shutdown(ctx, &wg)

	wg.Wait()
	return nil
}

// start launches all runners concurrently.
//
// Each runner is started in its own goroutine. If a runner returns an error,
// it is sent to errCh to trigger application shutdown.
func (a *App) start(ctx context.Context, wg *sync.WaitGroup, errCh chan<- error) {
	for i, runner := range a.runners {
		wg.Go(func() {
			a.log.Info("starting", slog.Int(RunnerKey, i))
			if err := runner.Start(ctx); err != nil {
				a.log.Error("exited",
					slog.Int(RunnerKey, i),
					slog.Any(Error, err),
				)
				errCh <- err
			}
		})
	}
}

// shutdown stops all runners concurrently.
//
// It attempts to gracefully shut down each runner. Errors are logged but
// do not stop the shutdown process.
//
// This function waits for all shutdown routines to complete before returning.
func (a *App) shutdown(ctx context.Context, wg *sync.WaitGroup) {
	for i, runner := range a.runners {
		wg.Go(func() {
			a.log.Info("stopping", slog.Int(RunnerKey, i))
			if err := runner.Shutdown(ctx); err != nil {
				a.log.Error("shutdown failed",
					slog.Int(RunnerKey, i),
					slog.Any(Error, err),
				)
			}
		})
	}
	wg.Wait()
}

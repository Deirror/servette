// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package logger

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/Deirror/servette/env"
	"github.com/Deirror/servette/transport"
)

// Inits the go standart logger, based on env mode.
func New(mode env.Mode) *slog.Logger {
	var h slog.Handler
	if mode.IsDev() {
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	} else {
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}
	logger := slog.New(h)
	return logger
}

// LogFunc logs a function call with optional extra attributes.
// It records the request ID from context and any error passed.
// Timing is NOT recorded here.
// Use this when you only want to log function entry or results without measuring duration.
func LogFunc(
	ctx context.Context,
	logger *slog.Logger,
	funcName string,
	err error,
	extraAttrs ...slog.Attr,
) {
	attrs := []slog.Attr{
		slog.Any(transport.ReqID, ctx.Value(transport.ReqID)),
		slog.Any(Error, err),
	}
	attrs = append(attrs, extraAttrs...)

	level := slog.LevelInfo
	if err != nil {
		level = slog.LevelError
	}

	logger.LogAttrs(ctx, level, funcName, attrs...)
}

// LogFuncWithTiming logs a function call including the elapsed time since `begin`.
// It records the request ID from context, the duration, any error, and optional extra attributes.
// Use this when you want to measure and log the time taken by a function.
func LogFuncWithTiming(
	ctx context.Context,
	logger *slog.Logger,
	funcName string,
	begin time.Time,
	err error,
	extraAttrs ...slog.Attr,
) {
	attrs := []slog.Attr{
		slog.Any(transport.ReqID, ctx.Value(transport.ReqID)),
		slog.Duration(Duration, time.Since(begin)),
		slog.Any(Error, err),
	}
	attrs = append(attrs, extraAttrs...)

	level := slog.LevelInfo
	if err != nil {
		level = slog.LevelError
	}

	logger.LogAttrs(ctx, level, funcName, attrs...)
}

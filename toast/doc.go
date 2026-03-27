// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package toast provides utilities to trigger client-side toast notifications
// via HTTP headers, typically in web applications using HTMX or similar frameworks.
//
// It standardizes sending structured notifications from the server to the client
// by setting a JSON-encoded value in the `HX-Trigger` HTTP header.
//
// # Features
//
//   - Trigger success, error, info, or warning messages
//   - Uses a consistent payload structure for front-end consumption
//   - Minimal API: just call Success, Error, Info, or Warning with a message
//
// # Usage
//
// Trigger a success toast:
//
//	toast.Success(w, "Operation completed successfully")
//
// Trigger an error toast:
//
//	toast.Error(w, "Something went wrong")
//
// # Payload Structure
//
// The header contains a JSON object with the following format:
//
//	{
//	  "showToast": {
//	    "message": "your message here",
//	    "type": "success|error|info|warning"
//	  }
//	}
//
// This allows client-side listeners to easily parse and display the notification.
//
// # Notes
//
//   - The package assumes the client is listening to the `HX-Trigger` header
//   - The payload is marshaled as JSON; errors during marshaling are ignored
//   - Use this package for server-initiated, one-off notifications per request
package toast

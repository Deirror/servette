// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package header

type Key = string

const (
	Accept             Key = "Accept"
	AcceptLanguage     Key = "Accept-Language"
	Authorization      Key = "Authorization"
	CacheControl       Key = "Cache-Control"
	Connection         Key = "Connection"
	Cookie             Key = "Cookie"
	ContentLength      Key = "Content-Length"
	ContentType        Key = "Content-Type"
	Origin             Key = "Origin"
	Host               Key = "Host"
	UserAgent          Key = "User-Agent"
	XForwardedFor      Key = "X-Forwarded-For"
	XRealIP            Key = "X-Real-IP"
	SetCookie          Key = "Set-Cookie"
	Location           Key = "Location"
	RetryAfter         Key = "Retry-After"
	Server             Key = "Server"
	TextHTML           Key = "text/html"
	ApplicationJSON    Key = "application/json"
	HXPushURL          Key = "HX-Push-Url"
	HXTrigger          Key = "HX-Trigger"
	HXTriggerAfterSwap Key = "HX-Trigger-After-Swap"
)

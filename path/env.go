// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package pathx

import "github.com/Deirror/servette/env"

// AppendAppMode adds .(app mode) on the env file path.
// It assumes only prod and dev, since we cannot know
// if it is staging or dev, we fallback to dev only.
func AppendAppMode(isProd bool, envPath string) string {
	if !isProd {
		return envPath + "." + string(env.Dev)
	}
	return envPath + "." + string(env.Prod)
}

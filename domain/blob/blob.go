// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package blob

import "context"

type Storer interface {
	Download(ctx context.Context, bucket, path string) ([]byte, error)
	Upload(ctx context.Context, bucket, path string, data []byte) error
	CreateSignedURL(ctx context.Context, bucket, path string, expiresInSec int, forUpload, upsert bool) (string, string, error)
}

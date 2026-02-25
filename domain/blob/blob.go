package blob

import "context"

type Storer interface {
	Download(ctx context.Context, bucket, path string) ([]byte, error)
	Upload(ctx context.Context, bucket, path string, data []byte) error
	CreateSignedURL(ctx context.Context, bucket, path string, expiresInSec int, forUpload, upsert bool) (string, string, error)
}

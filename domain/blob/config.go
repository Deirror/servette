package blob

// Config holds the configuration required to connect to a blob storage provider.
type Config struct {
	ProjectURL string // Base project URL of the blob storage service
	APIKey     string // API key for authentication with the blob storage provider
	Bucket     string // Target bucket name for storing or retrieving objects
}

func NewConfig(url, apiKey, bucket string) *Config {
	return &Config{
		ProjectURL: url,
		APIKey:     apiKey,
		Bucket:     bucket,
	}
}

// WithProjectURL sets the ProjectURL and returns the updated Config.
func (c *Config) WithProjectURL(url string) *Config {
	c.ProjectURL = url
	return c
}

// WithAPIKey sets the APIKey and returns the updated Config.
func (c *Config) WithAPIKey(apiKey string) *Config {
	c.APIKey = apiKey
	return c
}

// WithBucket sets the Bucket and returns the updated Config.
func (c *Config) WithBucket(bucket string) *Config {
	c.Bucket = bucket
	return c
}

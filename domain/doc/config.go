package doc

// Config holds configuration details for connecting to a document database.
type Config struct {
	StoreURL string // used such as a MongoDB connection string (example: mongodb+srv://.../dbName)
	Database string // which database to connect to
}

func NewConfig(storeURL, db string) *Config {
	return &Config{
		StoreURL: storeURL,
		Database: db,
	}
}

// WithStoreURL sets a new document store URL and returns the updated Config.
func (c *Config) WithStoreURL(url string) *Config {
	c.StoreURL = url
	return c
}

// WithDatabse sets a new document databse and returns the updated Config.
func (c *Config) WithDatabse(db string) *Config {
	c.Database = db
	return c
}

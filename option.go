package deckv

// Config holds the configuration for Deckv instance.
type Config struct {
	ConfFilePath string
	Storage      DataStorage
}

// Option defines a function type for configuring Deckv instances.
type Option func(*Config)

// WithConfFilePath sets the path to the blocklist configuration file.
func WithConfFilePath(path string) Option {
	return func(c *Config) {
		c.ConfFilePath = path
	}
}

// WithStorage sets the storage backend to use for the blocklist.
func WithStorage(storage DataStorage) Option {
	return func(c *Config) {
		c.Storage = storage
	}
}

func withDefaults(c *Config) {
	if c.Storage == nil {
		c.Storage = &InMemStorage{}
	}
}

package cli

// Config holds the state of the CLI application
type Config struct {
	Next     string
	Previous string
}

// NewConfig creates a new Config instance
func NewConfig() *Config {
	return &Config{}
}

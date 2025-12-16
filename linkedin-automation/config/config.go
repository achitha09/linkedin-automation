package config

// Config represents application configuration values.
// Values can be loaded from environment variables or config files.
type Config struct {
	LinkedInEmail    string
	LinkedInPassword string
	DailyLimit       int
	UserAgent        string
}

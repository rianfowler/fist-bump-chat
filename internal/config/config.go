package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	v *viper.Viper
}

func New() *Configuration {
	v := viper.New()

	// Set default configurations if needed. For example:
	// v.SetDefault("key", "defaultValue")
	v.SetDefault("GOOGLE_REDIRECT_URL", "http://localhost:8080/auth/callback/google")

	// Enable reading from environment variables
	v.AutomaticEnv()

	// For local development, read from .env file
	v.SetConfigType("env")
	v.SetConfigFile(".env")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore if using environment variables in production
		} else {
			// Some other error occurred while reading the config
			panic(err)
		}
	}

	return &Configuration{v: v}
}

func (c *Configuration) GetString(key string) string {
	return c.v.GetString(key)
}

// Similarly, add methods for other data types as needed
// like GetInt, GetBool, etc.

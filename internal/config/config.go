package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	v *viper.Viper
}

type ConfigDefault map[string]interface{}

var devDefaults = ConfigDefault{
	"SIGNIN_ENABLED":      false,
	"GOOGLE_REDIRECT_URL": "http://localhost:8080/auth/callback/google",
}

var prodDefaults = ConfigDefault{
	"SIGNIN_ENABLED":      true,
	"GOOGLE_REDIRECT_URL": "http://www.fistbump.io/auth/callback/google",
}

func New() *Configuration {
	v := viper.New()

	// Enable reading from environment variables
	v.AutomaticEnv()

	v.GetString("PROD")

	defaults := devDefaults

	if v.GetString("PROD") == "true" {
		defaults = prodDefaults
	}

	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	// Set default configurations if needed. For example:
	// v.SetDefault("key", "defaultValue")
	// v.SetDefault("GOOGLE_REDIRECT_URL", "http://localhost:8080/auth/callback/google")

	// For local development, read from .env file
	v.SetConfigType("env")
	v.SetConfigFile(".env")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore if using environment variables in production
			// TODO: add error handling
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

func (c *Configuration) GetBool(key string) bool {
	return c.v.GetBool(key)
}

// Similarly, add methods for other data types as needed
// like GetInt, GetBool, etc.

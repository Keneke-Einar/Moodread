package config

import (
	advancedConfig "github.com/num30/config"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"sync"
)

type Config struct {
	Env string `mapstructure:"ENV"`

	Database DatabaseConfig `mapstructure:",squash"`
	S3       S3Config       `mapstructure:",squash"`
	Server   Server         `mapstructure:",squash"`
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		viper.SetConfigFile(".env") // Load from .env file if available
		viper.AutomaticEnv()        // Read from system environment variables
		bindAllEnvVars()            // Automatically bind all env variables

		// Try loading .env file but do NOT fail if it doesn't exist
		if _, err := os.Stat(".env"); err == nil {
			err := viper.ReadInConfig()
			if err != nil {
				log.Printf("Warning: Failed to read config file: %v. Falling back to env variables.", err)
			}
		}

		// Unmarshal environment variables into Config struct
		var config Config
		if err := viper.Unmarshal(&config); err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		if err := advancedConfig.NewConfReader("zeep-config").Read(&config); err != nil {
			log.Fatalf("Error validating config: %v", err)
		}

		cfg = &config
	})

	return cfg
}

func bindAllEnvVars() {
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		key := pair[0]
		_ = viper.BindEnv(key)
	}
}

func MustGetConfig() *Config {
	if cfg == nil {
		log.Fatal("Config not initialized. Call LoadConfig() first.")
	}
	return cfg
}

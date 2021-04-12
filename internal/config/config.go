package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rl404/hayasui/internal/errors"
)

// Config is config for hayasui.
type Config struct {
	// Discord token.
	Token string `envconfig:"TOKEN"`
	// Discord command prefix.
	Prefix string `envconfig:"PREFIX" default:">"`
	// API host.
	APIHost string `envconfig:"API_HOST"`
	// Link to entry page.
	LinkHost string `envconfig:"LINK_HOST" default:"https://myanimelist.net"`
	// Redis config.
	Redis redisConfig `envconfig:"REDIS"`
}

type redisConfig struct {
	// Redis address with format `host:port`.
	Address string `envconfig:"ADDRESS"`
	// Redis password if exists.
	Password string `envconfig:"PASSWORD"`
	// Caching time duration (in seconds).
	Time int `envconfig:"TIME" default:"86400"`
}

const envPath = "../../.env"
const envPrefix = "HYS"

// GetConfig to read and parse env.
func GetConfig() (cfg Config, err error) {
	// Load .env file if exists.
	godotenv.Load(envPath)

	// Convert env to struct.
	if err = envconfig.Process(envPrefix, &cfg); err != nil {
		return cfg, err
	}

	// Validate config.
	if cfg.Token == "" {
		return cfg, errors.ErrRequiredToken
	}

	if cfg.Prefix == "" {
		return cfg, errors.ErrRequiredPrefix
	}

	if cfg.APIHost == "" {
		return cfg, errors.ErrRequiredAPI
	}

	if cfg.Redis.Address == "" {
		return cfg, errors.ErrRequiredRedis
	}

	if cfg.Redis.Time <= 0 {
		return cfg, errors.ErrInvalidCacheTime
	}

	return cfg, nil
}

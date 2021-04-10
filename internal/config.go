package internal

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	// Discord token.
	Token string `envconfig:"TOKEN"`
	// Discord command prefix.
	Prefix string `envconfig:"PREFIX" default:">"`
	// API host.
	ApiHost string `envconfig:"API_HOST"`
	// Link to entry page.
	LinkHost string `envconfig:"LINK_HOST"`
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
func GetConfig() (cfg config, err error) {
	// Load .env file if exists.
	godotenv.Load(envPath)

	// Convert env to struct.
	if err = envconfig.Process(envPrefix, &cfg); err != nil {
		return cfg, err
	}

	// Validate config.
	if cfg.Token == "" {
		return cfg, errRequiredToken
	}

	if cfg.Prefix == "" {
		return cfg, errRequiredPrefix
	}

	if cfg.ApiHost == "" {
		return cfg, errRequiredAPI
	}

	if cfg.Redis.Address == "" {
		return cfg, errRequiredRedis
	}

	if cfg.Redis.Time <= 0 {
		return cfg, errInvalidCacheTime
	}

	return cfg, nil
}

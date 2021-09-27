package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rl404/hayasui/internal/errors"
)

// Config is config for hayasui.
type Config struct {
	// Discord token.
	Token string `envconfig:"TOKEN" required:"true"`
	// Discord command prefix.
	Prefix string `envconfig:"PREFIX" required:"true" default:">"`
	// Link to entry page.
	RedirectHost string `envconfig:"REDIRECT_HOST" required:"true" default:"https://anilist.co"`
	// Redis config.
	Redis redisConfig `envconfig:"REDIS"`
}

type redisConfig struct {
	// Redis address with format `host:port`.
	Address string `envconfig:"ADDRESS" required:"true" default:"localhost:6379"`
	// Redis password if exists.
	Password string `envconfig:"PASSWORD"`
	// Caching time duration (in seconds).
	Time int `envconfig:"TIME" required:"true" default:"86400"`
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

	if cfg.Redis.Time <= 0 {
		return cfg, errors.ErrInvalidCacheTime
	}

	return cfg, nil
}

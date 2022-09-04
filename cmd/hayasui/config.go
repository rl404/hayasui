package main

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rl404/fairy/cache"
	"github.com/rl404/fairy/log"
	"github.com/rl404/hayasui/internal/errors"
	"github.com/rl404/hayasui/internal/utils"
)

type config struct {
	Discord  discordConfig  `envconfig:"DISCORD"`
	Cache    cacheConfig    `envconfig:"CACHE"`
	Log      logConfig      `envconfig:"LOG"`
	Newrelic newrelicConfig `envconfig:"NEWRELIC"`
}

type discordConfig struct {
	Token  string `envconfig:"TOKEN" validate:"required" mod:"trim"`
	Prefix string `envconfig:"PREFIX" validate:"required" mod:"trim,default=>"`
}

type cacheConfig struct {
	Dialect  string        `envconfig:"DIALECT" validate:"required,oneof=redis inmemory memcache" mod:"default=inmemory,no_space,lcase"`
	Address  string        `envconfig:"ADDRESS"`
	Password string        `envconfig:"PASSWORD"`
	Time     time.Duration `envconfig:"TIME" validate:"required,gt=0" mod:"default=24h"`
}

type logConfig struct {
	Type  log.LogType  `envconfig:"TYPE" default:"2"`
	Level log.LogLevel `envconfig:"LEVEL" default:"-1"`
	JSON  bool         `envconfig:"JSON" default:"false"`
	Color bool         `envconfig:"COLOR" default:"true"`
}

type newrelicConfig struct {
	Name       string `envconfig:"NAME" default:"hayasui"`
	LicenseKey string `envconfig:"LICENSE_KEY"`
}

const envPath = "../../.env"
const envPrefix = "HYS"

var cacheType = map[string]cache.CacheType{
	"redis":    cache.Redis,
	"inmemory": cache.InMemory,
	"memcache": cache.Memcache,
}

func getConfig() (*config, error) {
	var cfg config

	// Load .env file.
	_ = godotenv.Load(envPath)

	// Convert env to struct.
	if err := envconfig.Process(envPrefix, &cfg); err != nil {
		return nil, err
	}

	if cfg.Cache.Time <= 0 {
		return nil, errors.ErrInvalidCacheTime
	}

	// Validate.
	if err := utils.Validate(&cfg); err != nil {
		return nil, err
	}

	// Init global log.
	if err := utils.InitLog(cfg.Log.Type, cfg.Log.Level, cfg.Log.JSON, cfg.Log.Color); err != nil {
		return nil, err
	}

	return &cfg, nil
}

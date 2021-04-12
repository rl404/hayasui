package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rl404/hayasui/internal/api"
	_cache "github.com/rl404/hayasui/internal/api/cache"
	"github.com/rl404/hayasui/internal/api/http"
	"github.com/rl404/hayasui/internal/cache"
	"github.com/rl404/hayasui/internal/config"
	"github.com/rl404/hayasui/internal/discord"
	"github.com/rl404/hayasui/internal/handler"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "hayasui",
		Short: "Discord bot for anime/manga/character/people database.",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "bot",
		Short: "Run bot",
		RunE: func(*cobra.Command, []string) error {
			return bot()
		},
	})

	if err := cmd.Execute(); err != nil {
		log.Println(err)
	}
}

func bot() error {
	// Get config.
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	// Init redis.
	redis, err := cache.NewCache(cfg.Redis.Address, cfg.Redis.Password, time.Duration(cfg.Redis.Time)*time.Second)
	if err != nil {
		return err
	}
	defer redis.Close()

	// Init api.
	var service api.API
	service = http.New(cfg.APIHost)
	service = _cache.New(redis, service)

	// Init discord.
	d, err := discord.New(cfg.Token)
	if err != nil {
		return err
	}
	defer d.Close()

	// Init handler.
	mh := handler.NewMessageHandler(service, redis, cfg.Prefix, cfg.LinkHost)
	rh := handler.NewReactionHandler(service, redis, cfg.LinkHost)

	// Add handler.
	d.AddMessageHandler(mh.Handler())
	d.AddReactionHandler(rh.Handler())

	// Run bot.
	if err = d.Run(); err != nil {
		return err
	}

	log.Println("hayasui is running...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-quit

	return nil
}

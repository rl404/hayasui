package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rl404/hayasui/internal"
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
	cfg, err := internal.GetConfig()
	if err != nil {
		return err
	}

	// Init api.
	api, err := internal.NewAPI(cfg.ApiHost)
	if err != nil {
		return err
	}

	// Init redis.
	redis, err := internal.NewCache(cfg.Redis.Address, cfg.Redis.Password, time.Duration(cfg.Redis.Time)*time.Second)
	if err != nil {
		return err
	}
	defer redis.Close()

	// Init discord.
	discord, err := internal.NewDiscord(cfg.Token)
	if err != nil {
		return err
	}

	// Init handler.
	mh := internal.NewMessageHandler(api, redis, cfg.Prefix, cfg.LinkHost)
	rh := internal.NewReactionHandler(api, redis, cfg.LinkHost)

	// Add handler.
	discord.AddMessageHandler(mh.Handler())
	discord.AddReactionHandler(rh.Handler())

	// Run bot.
	if err = discord.Run(); err != nil {
		return err
	}
	defer discord.Close()

	log.Println("hayasui is running...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-quit

	return nil
}

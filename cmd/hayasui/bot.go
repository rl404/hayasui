package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/fairy/cache"
	_nr "github.com/rl404/fairy/log/newrelic"
	nrCache "github.com/rl404/fairy/monitoring/newrelic/cache"
	_bot "github.com/rl404/hayasui/internal/delivery/bot"
	animeRepository "github.com/rl404/hayasui/internal/domain/anime/repository"
	animeAnilist "github.com/rl404/hayasui/internal/domain/anime/repository/anilist"
	animeCache "github.com/rl404/hayasui/internal/domain/anime/repository/cache"
	discordRepository "github.com/rl404/hayasui/internal/domain/discord/repository"
	discordClient "github.com/rl404/hayasui/internal/domain/discord/repository/client"
	reactionRepository "github.com/rl404/hayasui/internal/domain/reaction/repository"
	reactionCache "github.com/rl404/hayasui/internal/domain/reaction/repository/cache"
	templateRepository "github.com/rl404/hayasui/internal/domain/template/repository"
	templateClient "github.com/rl404/hayasui/internal/domain/template/repository/client"
	"github.com/rl404/hayasui/internal/service"
	"github.com/rl404/hayasui/internal/utils"
)

func bot() error {
	// Get config.
	cfg, err := getConfig()
	if err != nil {
		return err
	}

	// Init newrelic.
	nrApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.Newrelic.Name),
		newrelic.ConfigLicense(cfg.Newrelic.LicenseKey),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		utils.Error(err.Error())
	} else {
		defer nrApp.Shutdown(10 * time.Second)
		utils.AddLog(_nr.NewFromNewrelicApp(nrApp, _nr.ErrorLevel))
		utils.Info("newrelic initialized")
	}

	// Init cache.
	c, err := cache.New(cacheType[cfg.Cache.Dialect], cfg.Cache.Address, cfg.Cache.Password, cfg.Cache.Time)
	if err != nil {
		return err
	}
	c = nrCache.New(cfg.Cache.Dialect, cfg.Cache.Address, c)
	utils.Info("cache initialized")
	defer c.Close()

	// Init discord.
	var discord discordRepository.Repository
	discord, err = discordClient.New(cfg.Discord.Token)
	if err != nil {
		return err
	}
	utils.Info("discord initialized")

	// Init template.
	var template templateRepository.Repository = templateClient.New(cfg.Discord.Prefix)
	utils.Info("template initialized")

	// Init anime.
	var anime animeRepository.Repository
	anime = animeAnilist.New()
	anime = animeCache.New(c, anime)
	utils.Info("anime initialized")

	// Init reaction.
	var reaction reactionRepository.Repository = reactionCache.New(c)
	utils.Info("reaction initialized")

	// Init service.
	service := service.New(discord, template, anime, reaction)
	utils.Info("service initialized")

	// Init bot.
	bot := _bot.New(service, cfg.Discord.Prefix)
	bot.RegisterHandler(nrApp)
	utils.Info("bot initialized")

	// Run bot.
	if err := bot.Run(); err != nil {
		return err
	}
	utils.Info("hayasui is running...")
	defer bot.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	return nil
}

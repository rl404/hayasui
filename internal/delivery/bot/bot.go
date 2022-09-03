package bot

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/hayasui/internal/errors"
	"github.com/rl404/hayasui/internal/service"
	"github.com/rl404/hayasui/internal/utils"
)

// Bot contains functions for bot.
type Bot struct {
	service service.Service
	prefix  string
}

// New to create new bot.
func New(service service.Service, prefix string) *Bot {
	return &Bot{
		service: service,
		prefix:  prefix,
	}
}

// Run to run bot.
func (b *Bot) Run() error {
	return b.service.Run()
}

// Stop to stop bot.
func (b *Bot) Stop() error {
	return b.service.Stop()
}

// RegisterReadyHandler to register ready handler.
func (b *Bot) RegisterHandler(nrApp *newrelic.Application) {
	b.service.RegisterReadyHandler(b.readyHandler())
	b.service.RegisterMessageHandler(b.messageHandler(nrApp))
	b.service.RegisterReactionHandler(b.reactionHandler(nrApp))
}

func (b *Bot) log(ctx context.Context) {
	if rvr := recover(); rvr != nil {
		errors.Wrap(ctx, fmt.Errorf("%v", rvr), fmt.Errorf("%s", debug.Stack()))
	}

	errStack := errors.Get(ctx)
	if len(errStack) > 0 {
		utils.Error(strings.Join(errStack, ","))
	}
}

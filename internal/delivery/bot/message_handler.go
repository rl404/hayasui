package bot

import (
	"context"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/fairy/errors/stack"
)

func (b *Bot) messageHandler(nrApp *newrelic.Application) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		ctx := stack.Init(context.Background())
		defer b.log(ctx)

		// Ignore all messages created by the bot itself.
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Command and prefix check.
		if b.prefixCheck(m.Content) {
			return
		}

		// Remove prefix.
		m.Content = b.cleanPrefix(m.Content)

		// Get arguments.
		r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`)
		args := r.FindAllString(m.Content, -1)

		tx := nrApp.StartTransaction("Command " + args[0])
		defer tx.End()

		ctx = newrelic.NewContext(ctx, tx)

		switch args[0] {
		case "ping":
			stack.Wrap(ctx, b.service.HandlePing(ctx, m))
		case "help", "h":
			stack.Wrap(ctx, b.service.HandleHelp(ctx, m))
		case "search", "s":
			stack.Wrap(ctx, b.service.HandleSearch(ctx, m, args))
		case "anime", "a":
			stack.Wrap(ctx, b.service.HandleAnime(ctx, m, args))
		case "manga", "m":
			stack.Wrap(ctx, b.service.HandleManga(ctx, m, args))
		case "character", "char", "c":
			stack.Wrap(ctx, b.service.HandleCharacter(ctx, m, args))
		case "people", "p":
			stack.Wrap(ctx, b.service.HandlePeople(ctx, m, args))
		}
	}
}

func (b *Bot) prefixCheck(cmd string) bool {
	return len(cmd) <= len(b.prefix) || cmd[:len(b.prefix)] != b.prefix
}

func (b *Bot) cleanPrefix(cmd string) string {
	return strings.TrimSpace(cmd[len(b.prefix):])
}

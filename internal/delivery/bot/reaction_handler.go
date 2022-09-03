package bot

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/hayasui/internal/errors"
)

func (b *Bot) reactionHandler(nrApp *newrelic.Application) func(*discordgo.Session, *discordgo.MessageReactionAdd) {
	return func(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
		ctx := errors.Init(context.Background())
		defer b.log(ctx)

		// Ignore all messages created by the bot itself.
		if m.UserID == s.State.User.ID {
			return
		}

		cmd, err := b.service.GetReaction(ctx, m)
		if err != nil {
			return
		}

		tx := nrApp.StartTransaction("Reaction " + m.Emoji.ID)
		defer tx.End()

		ctx = newrelic.NewContext(ctx, tx)

		switch cmd.Command {
		case "search", "s":
			errors.Wrap(ctx, b.service.HandleSearchReaction(ctx, m, *cmd))
		case "anime", "a":
			errors.Wrap(ctx, b.service.HandleAnimeReaction(ctx, m, *cmd))
		case "manga", "m":
			errors.Wrap(ctx, b.service.HandleMangaReaction(ctx, m, *cmd))
		case "character", "char", "c":
			errors.Wrap(ctx, b.service.HandleCharacterReaction(ctx, m, *cmd))
		case "people", "p":
			errors.Wrap(ctx, b.service.HandlePeopleReaction(ctx, m, *cmd))
		}
	}
}

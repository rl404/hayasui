package service

import (
	"context"

	"github.com/bwmarrin/discordgo"
	reactionEntity "github.com/rl404/hayasui/internal/domain/reaction/entity"
	"github.com/rl404/hayasui/internal/errors"
)

// HandleSearch to handle search.
func (s *service) HandleSearch(ctx context.Context, m *discordgo.MessageCreate, args []string) error {
	if len(args) <= 2 {
		return errors.Wrap(ctx, s.handleInvalid(ctx, m.ChannelID))
	}

	switch args[1] {
	case "anime", "a":
		return errors.Wrap(ctx, s.handleSearchAnime(ctx, m, args))
	case "manga", "m":
		return errors.Wrap(ctx, s.handleSearchManga(ctx, m, args))
	case "character", "char", "c":
		return errors.Wrap(ctx, s.handleSearchCharacter(ctx, m, args))
	case "people", "p":
		return errors.Wrap(ctx, s.handleSearchPeople(ctx, m, args))
	}

	return errors.Wrap(ctx, s.handleInvalid(ctx, m.ChannelID))
}

// HandleSearchReaction to handle search reaction.
func (s *service) HandleSearchReaction(ctx context.Context, m *discordgo.MessageReactionAdd, cmd reactionEntity.Command) error {
	switch cmd.Arg {
	case "anime", "a":
		return errors.Wrap(ctx, s.handleSearchAnimeReaction(ctx, m, cmd))
	case "manga", "m":
		return errors.Wrap(ctx, s.handleSearchMangaReaction(ctx, m, cmd))
	case "character", "char", "c":
		return errors.Wrap(ctx, s.handleSearchCharacterReaction(ctx, m, cmd))
	case "people", "p":
		return errors.Wrap(ctx, s.handleSearchPeopleReaction(ctx, m, cmd))
	}

	return errors.Wrap(ctx, s.handleInvalid(ctx, m.ChannelID))
}

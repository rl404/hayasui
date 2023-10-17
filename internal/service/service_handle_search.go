package service

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/rl404/fairy/errors/stack"
	reactionEntity "github.com/rl404/hayasui/internal/domain/reaction/entity"
)

// HandleSearch to handle search.
func (s *service) HandleSearch(ctx context.Context, m *discordgo.MessageCreate, args []string) error {
	if len(args) <= 2 {
		return stack.Wrap(ctx, s.handleInvalid(ctx, m.ChannelID))
	}

	switch args[1] {
	case "anime", "a":
		return stack.Wrap(ctx, s.handleSearchAnime(ctx, m, args))
	case "manga", "m":
		return stack.Wrap(ctx, s.handleSearchManga(ctx, m, args))
	case "character", "char", "c":
		return stack.Wrap(ctx, s.handleSearchCharacter(ctx, m, args))
	case "people", "p":
		return stack.Wrap(ctx, s.handleSearchPeople(ctx, m, args))
	}

	return stack.Wrap(ctx, s.handleInvalid(ctx, m.ChannelID))
}

// HandleSearchReaction to handle search reaction.
func (s *service) HandleSearchReaction(ctx context.Context, m *discordgo.MessageReactionAdd, cmd reactionEntity.Command) error {
	switch cmd.Arg {
	case "anime", "a":
		return stack.Wrap(ctx, s.handleSearchAnimeReaction(ctx, m, cmd))
	case "manga", "m":
		return stack.Wrap(ctx, s.handleSearchMangaReaction(ctx, m, cmd))
	case "character", "char", "c":
		return stack.Wrap(ctx, s.handleSearchCharacterReaction(ctx, m, cmd))
	case "people", "p":
		return stack.Wrap(ctx, s.handleSearchPeopleReaction(ctx, m, cmd))
	}

	return stack.Wrap(ctx, s.handleInvalid(ctx, m.ChannelID))
}

package service

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/rl404/fairy/errors/stack"
	reactionEntity "github.com/rl404/hayasui/internal/domain/reaction/entity"
	"github.com/rl404/hayasui/internal/domain/template/entity"
)

// Run to run discord bot.
func (s *service) Run() error {
	return s.discord.Run()
}

// Stop to stop discord bot.
func (s *service) Stop() error {
	return s.discord.Close()
}

// RegisterReadyHandler to register discord ready handler.
func (s *service) RegisterReadyHandler(fn func(*discordgo.Session, *discordgo.Ready)) {
	s.discord.AddReadyHandler(fn)
}

// RegisterMessageHandler to register discord message handler.
func (s *service) RegisterMessageHandler(fn func(*discordgo.Session, *discordgo.MessageCreate)) {
	s.discord.AddMessageHandler(fn)
}

// RegisterReactionHandler to register discord reaction handler.
func (s *service) RegisterReactionHandler(fn func(*discordgo.Session, *discordgo.MessageReactionAdd)) {
	s.discord.AddReactionHandler(fn)
}

// HandlePing to handle ping.
func (s *service) HandlePing(ctx context.Context, m *discordgo.MessageCreate) error {
	_, err := s.discord.SendMessage(ctx, m.ChannelID, "pong")
	return stack.Wrap(ctx, err)
}

// HandleHelp to handle help.
func (s *service) HandleHelp(ctx context.Context, m *discordgo.MessageCreate) error {
	_, err := s.discord.SendMessageEmbed(ctx, m.ChannelID, s.template.GetHelp())
	return stack.Wrap(ctx, err)
}

func (s *service) handleInvalid(ctx context.Context, channelID string) error {
	_, err := s.discord.SendMessage(ctx, channelID, s.template.GetInvalid())
	return stack.Wrap(ctx, err)
}

func (s *service) handleInvalidID(ctx context.Context, channelID string) error {
	_, err := s.discord.SendMessage(ctx, channelID, entity.MsgInvalidID)
	return stack.Wrap(ctx, err)
}

func (s *service) handleError(ctx context.Context, channelID string, err error) error {
	if _, err2 := s.discord.SendMessage(ctx, channelID, err.Error()); err2 != nil {
		return stack.Wrap(ctx, err2)
	}
	return stack.Wrap(ctx, err)
}

// GetReaction to get discord reaction.
func (s *service) GetReaction(ctx context.Context, m *discordgo.MessageReactionAdd) (*reactionEntity.Command, error) {
	cmd, err := s.reaction.GetCommand(ctx, m.MessageID)
	if err != nil {
		return nil, stack.Wrap(ctx, err)
	}

	// Remove user reaction.
	if err := s.discord.RemoveMessageReaction(ctx, m.ChannelID, m.MessageID, m.Emoji.Name, m.UserID); err != nil {
		return nil, stack.Wrap(ctx, err)
	}

	return cmd, nil
}

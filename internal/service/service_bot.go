package service

import (
	"context"

	"github.com/bwmarrin/discordgo"
	reactionEntity "github.com/rl404/hayasui/internal/domain/reaction/entity"
	"github.com/rl404/hayasui/internal/domain/template/entity"
	"github.com/rl404/hayasui/internal/errors"
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
	return errors.Wrap(ctx, err)
}

// HandleHelp to handle help.
func (s *service) HandleHelp(ctx context.Context, m *discordgo.MessageCreate) error {
	_, err := s.discord.SendMessageEmbed(ctx, m.ChannelID, s.template.GetHelp())
	return errors.Wrap(ctx, err)
}

func (s *service) handleInvalid(ctx context.Context, channelID string) error {
	_, err := s.discord.SendMessage(ctx, channelID, s.template.GetInvalid())
	return errors.Wrap(ctx, err)
}

func (s *service) handleInvalidID(ctx context.Context, channelID string) error {
	_, err := s.discord.SendMessage(ctx, channelID, entity.MsgInvalidID)
	return errors.Wrap(ctx, err)
}

func (s *service) handleError(ctx context.Context, channelID string, err error) error {
	if _, err2 := s.discord.SendMessage(ctx, channelID, err.Error()); err2 != nil {
		return errors.Wrap(ctx, err2)
	}
	return errors.Wrap(ctx, err)
}

// GetReaction to get discord reaction.
func (s *service) GetReaction(ctx context.Context, m *discordgo.MessageReactionAdd) (*reactionEntity.Command, error) {
	defer func() {
		// Remove user reaction.
		if err := s.discord.RemoveMessageReaction(ctx, m.ChannelID, m.MessageID, m.Emoji.Name, m.UserID); err != nil {
			errors.Wrap(ctx, err)
		}
	}()

	cmd, err := s.reaction.GetCommand(ctx, m.MessageID)
	if err != nil {
		return nil, errors.Wrap(ctx, err)
	}

	return cmd, nil
}

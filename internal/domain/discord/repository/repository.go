package repository

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

// Repository contains functions for discord domain.
type Repository interface {
	Run() error
	Close() error

	AddReadyHandler(func(*discordgo.Session, *discordgo.Ready))
	AddMessageHandler(func(*discordgo.Session, *discordgo.MessageCreate))
	AddReactionHandler(func(*discordgo.Session, *discordgo.MessageReactionAdd))

	SendMessage(ctx context.Context, channelID, content string) (string, error)
	SendMessageEmbed(ctx context.Context, channelID string, content *discordgo.MessageEmbed) (string, error)
	EditMessageEmbed(ctx context.Context, channelID, messageID string, content *discordgo.MessageEmbed) (string, error)

	AddMessageReaction(ctx context.Context, channelID string, messageID string, content string) error
	RemoveMessageReaction(ctx context.Context, channelID, emojiID, emojiName, userID string) error
}

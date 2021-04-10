package internal

import "github.com/bwmarrin/discordgo"

// Discord contains all basic discord function.
type Discord interface {
	AddMessageHandler(func(*discordgo.Session, *discordgo.MessageCreate))
	AddReactionHandler(func(*discordgo.Session, *discordgo.MessageReactionAdd))
	Run() error
	Close() error
}

type discord struct {
	client *discordgo.Session
}

// NewDiscord to create new discord client.
func NewDiscord(token string) (Discord, error) {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return &discord{
		client: client,
	}, nil
}

// AddMessageHandler to add message handler.
func (d *discord) AddMessageHandler(handler func(*discordgo.Session, *discordgo.MessageCreate)) {
	d.client.AddHandler(handler)
}

// AddReactionHandler to add reaction handler.
func (d *discord) AddReactionHandler(handler func(*discordgo.Session, *discordgo.MessageReactionAdd)) {
	d.client.AddHandler(handler)
}

// Run to login and start discord bot.
func (d *discord) Run() error {
	return d.client.Open()
}

// Close to stop discord bot.
func (d *discord) Close() error {
	return d.client.Close()
}

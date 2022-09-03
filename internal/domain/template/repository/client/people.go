package client

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rl404/hayasui/internal/domain/template/entity"
	"github.com/rl404/hayasui/internal/utils"
)

// GetPeople to get people template.
func (c *Client) GetPeople(data entity.People, info int8) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Name,
		URL:   data.URL,
		Color: entity.ColorPurple,
	}

	if info == entity.InfoSimple {
		msg.Description = utils.Ellipsis(c.emptyCheck(data.About), 500)
		msg.Thumbnail = &discordgo.MessageEmbedThumbnail{
			URL: data.Image,
		}
		return msg
	}

	msg.Image = &discordgo.MessageEmbedImage{
		URL: data.Image,
	}

	msg.Fields = []*discordgo.MessageEmbedField{
		{
			Name:   "Alternative Names",
			Value:  c.emptyCheck(strings.Join(data.AlternativeNames, ", ")),
			Inline: true,
		},
		{
			Name:   "Favorite",
			Value:  utils.Thousands(data.Favorite),
			Inline: true,
		},
		{
			Name:   "Birthday",
			Value:  c.emptyCheck(data.Birthday.ToStr()),
			Inline: true,
		},
		{
			Name:  "About",
			Value: utils.Ellipsis(c.emptyCheck(data.About), 1000),
		},
	}

	return msg
}

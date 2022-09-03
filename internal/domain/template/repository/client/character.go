package client

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rl404/hayasui/internal/domain/template/entity"
	"github.com/rl404/hayasui/internal/utils"
)

// GetCharacter to get character template.
func (c *Client) GetCharacter(data entity.Character, info int8) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Name,
		URL:   data.URL,
		Color: entity.ColorOrange,
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
			Name:   "Nicknames",
			Value:  c.emptyCheck(strings.Join(data.Nicknames, ", ")),
			Inline: true,
		},
		{
			Name:   "Japanese",
			Value:  c.emptyCheck(data.NameJapanese),
			Inline: true,
		},
		{
			Name:   "Favorite",
			Value:  utils.Thousands(data.Favorite),
			Inline: true,
		},
		{
			Name:  "About",
			Value: utils.Ellipsis(c.emptyCheck(data.About), 1000),
		},
	}

	return msg
}

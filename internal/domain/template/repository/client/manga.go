package client

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rl404/hayasui/internal/domain/template/entity"
	"github.com/rl404/hayasui/internal/utils"
)

// GetManga to get manga template.
func (c *Client) GetManga(data entity.Manga, info int8) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Title,
		URL:   data.URL,
		Color: entity.ColorGreen,
	}

	if info == entity.InfoSimple {
		msg.Description = utils.Ellipsis(c.emptyCheck(data.Synopsis), 500)
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
			Name:   "English",
			Value:  c.emptyCheck(data.TitleEnglish),
			Inline: true,
		},
		{
			Name:   "Japanese",
			Value:  c.emptyCheck(data.TitleJapanese),
			Inline: true,
		},
		{
			Name:   "Synonyms",
			Value:  c.emptyCheck(strings.Join(data.TitleSynonyms, ", ")),
			Inline: true,
		},
		{
			Name:  "Synopsis",
			Value: utils.Ellipsis(c.emptyCheck(data.Synopsis), 1000),
		},
		{
			Name:   "Score",
			Value:  fmt.Sprintf("%.2f", data.Score),
			Inline: true,
		},
		{
			Name:   "Member",
			Value:  utils.Thousands(data.Member),
			Inline: true,
		},
		{
			Name:   "Favorite",
			Value:  utils.Thousands(data.Favorite),
			Inline: true,
		},
		{
			Name:   "Type",
			Value:  c.emptyCheck(data.Type),
			Inline: true,
		},
		{
			Name:   "Status",
			Value:  c.emptyCheck(data.Status),
			Inline: true,
		},
		{
			Name:   "Chapter",
			Value:  utils.Thousands(data.Chapter),
			Inline: true,
		},
		{
			Name:   "Ranking",
			Value:  c.emptyCheck(data.Ranking),
			Inline: true,
		},
		{
			Name:   "Airing Start",
			Value:  c.emptyCheck(data.StartDate.ToStr()),
			Inline: true,
		},
		{
			Name:   "Airing End",
			Value:  c.emptyCheck(data.EndDate.ToStr()),
			Inline: true,
		},
	}

	return msg
}

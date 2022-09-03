package client

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/rl404/hayasui/internal/domain/template/entity"
	"github.com/rl404/hayasui/internal/utils"
)

// GetSearch to get search template.
func (c *Client) GetSearch(data []entity.Search, _type entity.SearchType, info int8, page, lastPage int) *discordgo.MessageEmbed {
	body := "```"
	if len(data) == 0 {
		body += "No result."
	} else {
		switch info {
		case entity.InfoSimple:
			body += c.getTableHeader([]string{"ID", _type.GetHeader()}, []int{6, 45}) + "\n"
			for _, d := range data {
				body += c.getTableRow([]string{strconv.Itoa(d.ID), utils.Ellipsis(d.Name, 45)}, []int{6, 45}) + "\n"
			}
		case entity.InfoMore:
			body += c.getTableHeader([]string{"ID", _type.GetHeader(), "Type"}, []int{6, 35, 7}) + "\n"
			for _, d := range data {
				body += c.getTableRow([]string{strconv.Itoa(d.ID), utils.Ellipsis(d.Name, 35), d.Type}, []int{6, 35, 7}) + "\n"
			}
		case entity.InfoAll:
			body += c.getTableHeader([]string{"ID", _type.GetHeader(), "Type", "Score"}, []int{6, 29, 7, 5}) + "\n"
			for _, d := range data {
				body += c.getTableRow([]string{strconv.Itoa(d.ID), utils.Ellipsis(d.Name, 29), d.Type, fmt.Sprintf("%.2f", d.Score)}, []int{6, 29, 7, 5}) + "\n"
			}
		}
	}
	body += "```"

	return &discordgo.MessageEmbed{
		Title:       _type.ToTitle() + " Search Results",
		Color:       _type.GetColor(),
		Description: body,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("page: %v / %v", page, lastPage),
		},
	}

}

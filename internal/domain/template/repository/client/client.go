package client

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rl404/hayasui/internal/domain/template/entity"
	"github.com/rl404/hayasui/internal/utils"
)

// Client is template client.
type Client struct {
	prefix string
}

// New to create new template client.
func New(prefix string) *Client {
	return &Client{
		prefix: prefix,
	}
}

func (c *Client) clean(str string) string {
	return strings.ReplaceAll(str, "{{prefix}}", c.prefix)
}

func (c *Client) emptyCheck(str string) string {
	if str == "" {
		return "-"
	}
	return str
}

// GetHelp to get help template.
func (c *Client) GetHelp() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Help",
		Description: entity.MsgHelpContent,
		Color:       entity.ColorGreyLight,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  c.clean(entity.MsgSearchCmd),
				Value: c.clean(entity.MsgSearchContent),
			},
			{
				Name:  c.clean(entity.MsgAnimeCmd),
				Value: c.clean(entity.MsgAnimeContent),
			},
			{
				Name:  c.clean(entity.MsgMangaCmd),
				Value: c.clean(entity.MsgMangaContent),
			},
			{
				Name:  c.clean(entity.MsgCharCmd),
				Value: c.clean(entity.MsgCharContent),
			},
			{
				Name:  c.clean(entity.MsgPeopleCmd),
				Value: c.clean(entity.MsgPeopleContent),
			},
		},
	}
}

// GetInvalid to get invalid template.
func (c *Client) GetInvalid() string {
	return c.clean(entity.MsgInvalid)
}

func (c *Client) getTableHeader(titles []string, len []int) (str string) {
	var dash []string
	for i := range titles {
		titles[i] = utils.PadRight(titles[i], len[i], " ")
		dash = append(dash, utils.PadRight("", len[i], "-"))
	}
	return fmt.Sprintf("%s\n%s", strings.Join(titles, " | "), strings.Join(dash, " | "))
}

func (c *Client) getTableRow(data []string, len []int) (str string) {
	for i := range data {
		data[i] = utils.PadRight(data[i], len[i], " ")
	}
	return strings.Join(data, " | ")
}

package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rl404/hayasui/internal/constant"
	"github.com/rl404/hayasui/internal/model"
	"github.com/rl404/hayasui/internal/utils"
)

// Templater contains function to prepare message template.
type Templater interface {
	GetHelp() *discordgo.MessageEmbed
	GetSearch(data []model.DataSearch, page int, _type string, count int) *discordgo.MessageEmbed
	GetAnime(data *model.DataAnimeManga, info bool) *discordgo.MessageEmbed
	GetManga(data *model.DataAnimeManga, info bool) *discordgo.MessageEmbed
	GetCharacter(data *model.DataCharPeople, info bool) *discordgo.MessageEmbed
	GetPeople(data *model.DataCharPeople, info bool) *discordgo.MessageEmbed
}

type template struct {
	linkHost string
}

func newTemplate(linkHost string) Templater {
	return &template{
		linkHost: linkHost,
	}
}

var color = map[string]int{
	constant.TypeAnime:     constant.ColorBlue,
	constant.TypeManga:     constant.ColorGreen,
	constant.TypeCharacter: constant.ColorOrange,
	constant.TypePeople:    constant.ColorPurple,
}

func (t *template) getTableHeader() string {
	return fmt.Sprintf("%s | %s\n%s | %s",
		utils.PadRight("ID", 6, " "),
		utils.PadRight("Name", 45, " "),
		utils.PadRight("", 6, "-"),
		utils.PadRight("", 45, "-"))
}

func (t *template) getTableBody(id int, name string) string {
	return fmt.Sprintf("%s | %s",
		utils.PadRight(strconv.Itoa(id), 6, " "),
		utils.PadRight(utils.Ellipsis(name, 45), 45, " "))
}

// GetHelp to get help message template.
func (t *template) GetHelp() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Help",
		Description: constant.MsgHelpCmd,
		Color:       constant.ColorGreyLight,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  constant.MsgSearchCmd,
				Value: constant.MsgSearchContent,
			},
			{
				Name:  constant.MsgAnimeCmd,
				Value: constant.MsgAnimeContent,
			},
			{
				Name:  constant.MsgMangaCmd,
				Value: constant.MsgMangaContent,
			},
			{
				Name:  constant.MsgCharCmd,
				Value: constant.MsgCharContent,
			},
			{
				Name:  constant.MsgPeopleCmd,
				Value: constant.MsgPeopleContent,
			},
		},
	}
}

// GetSearch to get search result message template.
func (t *template) GetSearch(data []model.DataSearch, page int, _type string, cnt int) *discordgo.MessageEmbed {
	body := "```"
	if len(data) > 0 {
		body += t.getTableHeader() + "\n"
		for _, d := range data {
			if _type == constant.TypeAnime || _type == constant.TypeManga {
				body += t.getTableBody(d.ID, d.Title) + "\n"
			} else {
				body += t.getTableBody(d.ID, d.Name) + "\n"
			}
		}
	} else {
		body += "No result."
	}
	body += "```"

	return &discordgo.MessageEmbed{
		Title:       strings.Title(_type) + " Search Results",
		Color:       color[_type],
		Description: body,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("page: %v / %v", page, (cnt/constant.DataPerPage)+1),
		},
	}
}

// GetAnime to get anime data message template.
func (t *template) GetAnime(data *model.DataAnimeManga, info bool) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Title,
		URL:   utils.GenerateLink(t.linkHost, constant.TypeAnime, data.ID),
		Color: constant.ColorBlue,
	}

	if !info {
		msg.Description = utils.Ellipsis(utils.EmptyCheck(data.Synopsis), 500)
		msg.Thumbnail = &discordgo.MessageEmbedThumbnail{
			URL: data.Image,
		}
	} else {
		msg.Image = &discordgo.MessageEmbedImage{
			URL: data.Image,
		}
		msg.Fields = []*discordgo.MessageEmbedField{
			{
				Name:   "English",
				Value:  utils.EmptyCheck(data.AltTitles.English),
				Inline: true,
			},
			{
				Name:   "Japanese",
				Value:  utils.EmptyCheck(data.AltTitles.Japanese),
				Inline: true,
			},
			{
				Name:   "Synonym",
				Value:  utils.EmptyCheck(data.AltTitles.Synonym),
				Inline: true,
			},
			{
				Name:  "Synopsis",
				Value: utils.Ellipsis(utils.EmptyCheck(data.Synopsis), 1000),
			},
			{
				Name:   "Rank",
				Value:  "#" + utils.Thousands(data.Rank),
				Inline: true,
			},
			{
				Name:   "Score",
				Value:  fmt.Sprintf("%.2f", data.Score),
				Inline: true,
			},
			{
				Name:   "Poplarity",
				Value:  "#" + utils.Thousands(data.Popularity),
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
				Value:  constant.AnimeTypes[data.Type],
				Inline: true,
			},
			{
				Name:   "Status",
				Value:  constant.AnimeStatuses[data.Status],
				Inline: true,
			},
			{
				Name:   "Episode",
				Value:  utils.Thousands(data.Episode),
				Inline: true,
			},
			{
				Name:   "Airing Start",
				Value:  utils.DateToStr(data.Airing.Start),
				Inline: true,
			},
		}
	}

	return msg
}

// GetManga to get manga data message template.
func (t *template) GetManga(data *model.DataAnimeManga, info bool) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Title,
		URL:   utils.GenerateLink(t.linkHost, constant.TypeManga, data.ID),
		Color: constant.ColorGreen,
	}

	if !info {
		msg.Description = utils.Ellipsis(utils.EmptyCheck(data.Synopsis), 500)
		msg.Thumbnail = &discordgo.MessageEmbedThumbnail{
			URL: data.Image,
		}
	} else {
		msg.Image = &discordgo.MessageEmbedImage{
			URL: data.Image,
		}
		msg.Fields = []*discordgo.MessageEmbedField{
			{
				Name:   "English",
				Value:  utils.EmptyCheck(data.AltTitles.English),
				Inline: true,
			},
			{
				Name:   "Japanese",
				Value:  utils.EmptyCheck(data.AltTitles.Japanese),
				Inline: true,
			},
			{
				Name:   "Synonym",
				Value:  utils.EmptyCheck(data.AltTitles.Synonym),
				Inline: true,
			},
			{
				Name:  "Synopsis",
				Value: utils.Ellipsis(utils.EmptyCheck(data.Synopsis), 1000),
			},
			{
				Name:   "Rank",
				Value:  "#" + utils.Thousands(data.Rank),
				Inline: true,
			},
			{
				Name:   "Score",
				Value:  fmt.Sprintf("%.2f", data.Score),
				Inline: true,
			},
			{
				Name:   "Poplarity",
				Value:  "#" + utils.Thousands(data.Popularity),
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
				Value:  constant.MangaTypes[data.Type],
				Inline: true,
			},
			{
				Name:   "Status",
				Value:  constant.MangaStatuses[data.Status],
				Inline: true,
			},
			{
				Name:   "Chapter",
				Value:  utils.Thousands(data.Chapter),
				Inline: true,
			},
			{
				Name:   "Publishing Start",
				Value:  utils.DateToStr(data.Publishing.Start),
				Inline: true,
			},
		}
	}

	return msg
}

// GetCharacter to get character data message template.
func (t *template) GetCharacter(data *model.DataCharPeople, info bool) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Name,
		URL:   utils.GenerateLink(t.linkHost, constant.TypeCharacter, data.ID),
		Color: constant.ColorOrange,
	}

	if !info {
		msg.Description = utils.Ellipsis(utils.EmptyCheck(data.About), 500)
		msg.Thumbnail = &discordgo.MessageEmbedThumbnail{
			URL: data.Image,
		}
	} else {
		msg.Image = &discordgo.MessageEmbedImage{
			URL: data.Image,
		}
		msg.Fields = []*discordgo.MessageEmbedField{
			{
				Name:   "Nicknames",
				Value:  utils.EmptyCheck(strings.Join(data.Nicknames, ", ")),
				Inline: true,
			},
			{
				Name:   "Japanese",
				Value:  utils.EmptyCheck(data.JapaneseName),
				Inline: true,
			},
			{
				Name:   "Favorite",
				Value:  utils.Thousands(data.Favorite),
				Inline: true,
			},
			{
				Name:  "About",
				Value: utils.Ellipsis(utils.EmptyCheck(data.About), 1000),
			},
		}
	}

	return msg
}

// GetPeople to get people data message template.
func (t *template) GetPeople(data *model.DataCharPeople, info bool) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Name,
		URL:   utils.GenerateLink(t.linkHost, constant.TypePeople, data.ID),
		Color: constant.ColorPurple,
	}

	if !info {
		msg.Description = utils.Ellipsis(utils.EmptyCheck(data.More), 500)
		msg.Thumbnail = &discordgo.MessageEmbedThumbnail{
			URL: data.Image,
		}
	} else {
		msg.Image = &discordgo.MessageEmbedImage{
			URL: data.Image,
		}
		msg.Fields = []*discordgo.MessageEmbedField{
			{
				Name:   "Given Name",
				Value:  utils.EmptyCheck(data.GivenName),
				Inline: true,
			},
			{
				Name:   "Family Name",
				Value:  utils.EmptyCheck(data.FamilyName),
				Inline: true,
			},
			{
				Name:   "Alternative Names",
				Value:  utils.EmptyCheck(strings.Join(data.AlternativeNames, ", ")),
				Inline: true,
			},
			{
				Name:   "Favorite",
				Value:  utils.Thousands(data.Favorite),
				Inline: true,
			},
			{
				Name:   "Birthday",
				Value:  utils.DateToStr(data.Birthday),
				Inline: true,
			},
			{
				Name:   "Website",
				Value:  utils.EmptyCheck(data.Website),
				Inline: true,
			},
			{
				Name:  "More",
				Value: utils.Ellipsis(utils.EmptyCheck(data.More), 1000),
			},
		}
	}

	return msg
}

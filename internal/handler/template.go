package handler

import (
	"fmt"
	"strconv"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/bwmarrin/discordgo"
	"github.com/rl404/hayasui/internal/constant"
	"github.com/rl404/hayasui/internal/model"
	"github.com/rl404/hayasui/internal/utils"
)

// Templater contains function to prepare message template.
type Templater interface {
	GetHelp() *discordgo.MessageEmbed
	GetSearchAnime(data []model.DataSearchAnimeManga, cmd model.Command) *discordgo.MessageEmbed
	GetSearchManga(data []model.DataSearchAnimeManga, cmd model.Command) *discordgo.MessageEmbed
	GetSearchCharacter(data []model.DataSearchCharPeople, cmd model.Command) *discordgo.MessageEmbed
	GetSearchPeople(data []model.DataSearchCharPeople, cmd model.Command) *discordgo.MessageEmbed
	GetAnime(data *model.DataAnimeManga, info bool) *discordgo.MessageEmbed
	GetManga(data *model.DataAnimeManga, info bool) *discordgo.MessageEmbed
	GetCharacter(data *model.DataCharPeople, info bool) *discordgo.MessageEmbed
	GetPeople(data *model.DataCharPeople, info bool) *discordgo.MessageEmbed
}

type template struct {
	linkHost  string
	converter *md.Converter
}

func newTemplate(linkHost string) Templater {
	return &template{
		linkHost:  linkHost,
		converter: md.NewConverter("", true, nil),
	}
}

var color = map[string]int{
	constant.TypeAnime:     constant.ColorBlue,
	constant.TypeManga:     constant.ColorGreen,
	constant.TypeCharacter: constant.ColorOrange,
	constant.TypePeople:    constant.ColorPurple,
}

func (t *template) getTableHeader(titles []string, len []int) (str string) {
	var dash []string
	for i := range titles {
		titles[i] = utils.PadRight(titles[i], len[i], " ")
		dash = append(dash, utils.PadRight("", len[i], "-"))
	}
	return fmt.Sprintf("%s\n%s", strings.Join(titles, " | "), strings.Join(dash, " | "))
}

func (t *template) getTableRow(data []string, len []int) (str string) {
	for i := range data {
		data[i] = utils.PadRight(data[i], len[i], " ")
	}
	return strings.Join(data, " | ")
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

// GetSearchAnime to get anime search result message template.
func (t *template) GetSearchAnime(data []model.DataSearchAnimeManga, cmd model.Command) *discordgo.MessageEmbed {
	body := "```"
	if len(data) == 0 {
		body += "No result."
	} else {
		switch cmd.Type {
		case 0:
			body += t.getTableHeader([]string{"ID", "Title"}, []int{6, 45}) + "\n"
			for _, d := range data {
				body += t.getTableRow([]string{strconv.Itoa(d.ID), utils.Ellipsis(d.Title, 45)}, []int{6, 45}) + "\n"
			}
		case 1:
			body += t.getTableHeader([]string{"ID", "Title", "Type"}, []int{6, 35, 7}) + "\n"
			for _, d := range data {
				body += t.getTableRow([]string{strconv.Itoa(d.ID), utils.Ellipsis(d.Title, 35), constant.MediaTypes[d.Type]}, []int{6, 35, 7}) + "\n"
			}
		case 2:
			body += t.getTableHeader([]string{"ID", "Title", "Type", "Score"}, []int{6, 29, 7, 5}) + "\n"
			for _, d := range data {
				body += t.getTableRow([]string{strconv.Itoa(d.ID), utils.Ellipsis(d.Title, 29), constant.MediaTypes[d.Type], strconv.Itoa(d.Score)}, []int{6, 29, 7, 5}) + "\n"
			}
		}
	}
	body += "```"

	return &discordgo.MessageEmbed{
		Title:       "Anime Search Results",
		Color:       color[constant.TypeAnime],
		Description: body,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("page: %v / %v", cmd.Page, cmd.LastPage),
		},
	}
}

// GetSearchManga to get manga search result message template.
func (t *template) GetSearchManga(data []model.DataSearchAnimeManga, cmd model.Command) *discordgo.MessageEmbed {
	body := "```"
	if len(data) == 0 {
		body += "No result."
	} else {
		switch cmd.Type {
		case 0:
			body += t.getTableHeader([]string{"ID", "Title"}, []int{6, 45}) + "\n"
			for _, d := range data {
				body += t.getTableRow([]string{strconv.Itoa(d.ID), utils.Ellipsis(d.Title, 45)}, []int{6, 45}) + "\n"
			}
		case 1:
			body += t.getTableHeader([]string{"ID", "Title", "Type"}, []int{6, 35, 7}) + "\n"
			for _, d := range data {
				body += t.getTableRow([]string{strconv.Itoa(d.ID), utils.Ellipsis(d.Title, 35), constant.MediaTypes[d.Type]}, []int{6, 35, 7}) + "\n"
			}
		case 2:
			body += t.getTableHeader([]string{"ID", "Title", "Type", "Score"}, []int{6, 29, 7, 5}) + "\n"
			for _, d := range data {
				body += t.getTableRow([]string{strconv.Itoa(d.ID), utils.Ellipsis(d.Title, 29), constant.MediaTypes[d.Type], strconv.Itoa(d.Score)}, []int{6, 29, 7, 5}) + "\n"
			}
		}
	}
	body += "```"

	return &discordgo.MessageEmbed{
		Title:       "Manga Search Results",
		Color:       color[constant.TypeManga],
		Description: body,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("page: %v / %v", cmd.Page, cmd.LastPage),
		},
	}
}

// GetSearchCharacter to get character search result message template.
func (t *template) GetSearchCharacter(data []model.DataSearchCharPeople, cmd model.Command) *discordgo.MessageEmbed {
	body := "```"
	if len(data) == 0 {
		body += "No result."
	} else {
		body += t.getTableHeader([]string{"ID", "Name"}, []int{6, 45}) + "\n"
		for _, d := range data {
			body += t.getTableRow([]string{strconv.Itoa(d.ID), utils.Ellipsis(d.Name, 45)}, []int{6, 45}) + "\n"
		}
	}
	body += "```"

	return &discordgo.MessageEmbed{
		Title:       "Character Search Results",
		Color:       color[constant.TypeCharacter],
		Description: body,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("page: %v / %v", cmd.Page, cmd.LastPage),
		},
	}
}

// GetSearchPeople to get people search result message template.
func (t *template) GetSearchPeople(data []model.DataSearchCharPeople, cmd model.Command) *discordgo.MessageEmbed {
	body := "```"
	if len(data) == 0 {
		body += "No result."
	} else {
		body += t.getTableHeader([]string{"ID", "Name"}, []int{6, 45}) + "\n"
		for _, d := range data {
			body += t.getTableRow([]string{strconv.Itoa(d.ID), utils.Ellipsis(d.Name, 45)}, []int{6, 45}) + "\n"
		}
	}
	body += "```"

	return &discordgo.MessageEmbed{
		Title:       "People Search Results",
		Color:       color[constant.TypePeople],
		Description: body,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("page: %v / %v", cmd.Page, cmd.LastPage),
		},
	}
}

// GetAnime to get anime data message template.
func (t *template) GetAnime(data *model.DataAnimeManga, info bool) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Title.Romaji,
		URL:   utils.GenerateLink(t.linkHost, constant.TypeAnime, data.ID),
		Color: constant.ColorBlue,
	}

	if !info {
		msg.Description = t.toMD(utils.Ellipsis(utils.EmptyCheck(data.Synopsis), 500))
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
				Value:  utils.EmptyCheck(data.Title.English),
				Inline: true,
			},
			{
				Name:   "Native",
				Value:  utils.EmptyCheck(data.Title.Native),
				Inline: true,
			},
			{
				Name:  "Synopsis",
				Value: t.toMD(utils.Ellipsis(utils.EmptyCheck(data.Synopsis), 1000)),
			},
			{
				Name:   "Score",
				Value:  strconv.Itoa(data.Score),
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
				Value:  constant.MediaTypes[data.Type],
				Inline: true,
			},
			{
				Name:   "Status",
				Value:  constant.MediaStatuses[data.Status],
				Inline: true,
			},
			{
				Name:   "Episode",
				Value:  utils.Thousands(data.Episode),
				Inline: true,
			},
			{
				Name:   "Ranking",
				Value:  utils.EmptyCheck(strings.Join(data.Rankings, "\n")),
				Inline: true,
			},
			{
				Name:   "Airing Start",
				Value:  utils.DateToStr(data.Airing.Start),
				Inline: true,
			},
			{
				Name:   "Airing End",
				Value:  utils.DateToStr(data.Airing.End),
				Inline: true,
			},
		}
	}

	return msg
}

// GetManga to get manga data message template.
func (t *template) GetManga(data *model.DataAnimeManga, info bool) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Title.Romaji,
		URL:   utils.GenerateLink(t.linkHost, constant.TypeManga, data.ID),
		Color: constant.ColorGreen,
	}

	if !info {
		msg.Description = t.toMD(utils.Ellipsis(utils.EmptyCheck(data.Synopsis), 500))
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
				Value:  utils.EmptyCheck(data.Title.English),
				Inline: true,
			},
			{
				Name:   "Native",
				Value:  utils.EmptyCheck(data.Title.Native),
				Inline: true,
			},
			{
				Name:  "Synopsis",
				Value: t.toMD(utils.Ellipsis(utils.EmptyCheck(data.Synopsis), 1000)),
			},
			{
				Name:   "Score",
				Value:  strconv.Itoa(data.Score),
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
				Value:  constant.MediaTypes[data.Type],
				Inline: true,
			},
			{
				Name:   "Status",
				Value:  constant.MediaStatuses[data.Status],
				Inline: true,
			},
			{
				Name:   "Chapter",
				Value:  utils.Thousands(data.Chapter),
				Inline: true,
			},
			{
				Name:   "Ranking",
				Value:  utils.EmptyCheck(strings.Join(data.Rankings, "\n")),
				Inline: true,
			},
			{
				Name:   "Publishing Start",
				Value:  utils.DateToStr(data.Publishing.Start),
				Inline: true,
			},
			{
				Name:   "Publishing End",
				Value:  utils.DateToStr(data.Publishing.End),
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
				Name:  "More",
				Value: utils.Ellipsis(utils.EmptyCheck(data.More), 800),
			},
		}
	}

	return msg
}

func (t *template) toMD(str string) string {
	s, _ := t.converter.ConvertString(str)
	return s
}

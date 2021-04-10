package internal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func repeat(str string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(str, n)
}

func padLeft(str string, l int, p string) string {
	return repeat(p, l-len(str)) + str
}

func padRight(str string, l int, p string) string {
	return str + repeat(p, l-len(str))
}

func ellipsis(str string, length int) string {
	if len(str) > length {
		return str[:length-3] + "..."
	}
	return str
}

func getTableHeader() string {
	return fmt.Sprintf("%s | %s\n%s | %s",
		padRight("ID", 6, " "),
		padRight("Name", 45, " "),
		padRight("", 6, "-"),
		padRight("", 45, "-"))
}

func getTableBody(id int, name string) string {
	return fmt.Sprintf("%s | %s",
		padRight(strconv.Itoa(id), 6, " "),
		padRight(ellipsis(name, 45), 45, " "))
}

func getRedisKey(id string) string {
	return "hayasui:" + id
}

func getLink(host string, path ...interface{}) string {
	for _, p := range path {
		host += fmt.Sprintf("/%v", p)
	}
	return host
}

func getSearchTemplate(data []searchData, page int, t string, color int, cnt int) *discordgo.MessageEmbed {
	var body string
	for _, d := range data {
		if t == anime || t == manga {
			body += getTableBody(d.ID, d.Title) + "\n"
		} else {
			body += getTableBody(d.ID, d.Name) + "\n"
		}
	}

	return &discordgo.MessageEmbed{
		Title: strings.Title(t) + " Search Results",
		Color: color,
		Description: fmt.Sprintf("```%s\n%s\n```",
			getTableHeader(),
			body,
		),
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("page: %v / %v", page, (cnt/dataPerPage)+1),
		},
	}
}

func getAnimeTemplate(data *animeMangaData, link string, info bool) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Title,
		URL:   getLink(link, "anime", data.ID),
		Color: blue,
	}

	if !info {
		msg.Description = ellipsis(emptyCheck(data.Synopsis), 500)
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
				Value:  emptyCheck(data.AltTitles.English),
				Inline: true,
			},
			{
				Name:   "Japanese",
				Value:  emptyCheck(data.AltTitles.Japanese),
				Inline: true,
			},
			{
				Name:   "Synonym",
				Value:  emptyCheck(data.AltTitles.Synonym),
				Inline: true,
			},
			{
				Name:  "Synopsis",
				Value: ellipsis(emptyCheck(data.Synopsis), 1000),
			},
			{
				Name:   "Rank",
				Value:  "#" + thousands(data.Rank),
				Inline: true,
			},
			{
				Name:   "Score",
				Value:  fmt.Sprintf("%.2f", data.Score),
				Inline: true,
			},
			{
				Name:   "Poplarity",
				Value:  "#" + thousands(data.Popularity),
				Inline: true,
			},
			{
				Name:   "Member",
				Value:  thousands(data.Member),
				Inline: true,
			},
			{
				Name:   "Favorite",
				Value:  thousands(data.Favorite),
				Inline: true,
			},
			{
				Name:   "Type",
				Value:  animeTypes[data.Type],
				Inline: true,
			},
			{
				Name:   "Status",
				Value:  animeStatuses[data.Status],
				Inline: true,
			},
			{
				Name:   "Episode",
				Value:  thousands(data.Episode),
				Inline: true,
			},
			{
				Name:   "Airing Start",
				Value:  toDate(data.Airing.Start),
				Inline: true,
			},
		}
	}

	return msg
}

func getMangaTemplate(data *animeMangaData, link string, info bool) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Title,
		URL:   getLink(link, "manga", data.ID),
		Color: green,
	}

	if !info {
		msg.Description = ellipsis(emptyCheck(data.Synopsis), 500)
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
				Value:  emptyCheck(data.AltTitles.English),
				Inline: true,
			},
			{
				Name:   "Japanese",
				Value:  emptyCheck(data.AltTitles.Japanese),
				Inline: true,
			},
			{
				Name:   "Synonym",
				Value:  emptyCheck(data.AltTitles.Synonym),
				Inline: true,
			},
			{
				Name:  "Synopsis",
				Value: ellipsis(emptyCheck(data.Synopsis), 1000),
			},
			{
				Name:   "Rank",
				Value:  "#" + thousands(data.Rank),
				Inline: true,
			},
			{
				Name:   "Score",
				Value:  fmt.Sprintf("%.2f", data.Score),
				Inline: true,
			},
			{
				Name:   "Poplarity",
				Value:  "#" + thousands(data.Popularity),
				Inline: true,
			},
			{
				Name:   "Member",
				Value:  thousands(data.Member),
				Inline: true,
			},
			{
				Name:   "Favorite",
				Value:  thousands(data.Favorite),
				Inline: true,
			},
			{
				Name:   "Type",
				Value:  mangaTypes[data.Type],
				Inline: true,
			},
			{
				Name:   "Status",
				Value:  mangaStatuses[data.Status],
				Inline: true,
			},
			{
				Name:   "Chapter",
				Value:  thousands(data.Chapter),
				Inline: true,
			},
			{
				Name:   "Publishing Start",
				Value:  toDate(data.Publishing.Start),
				Inline: true,
			},
		}
	}

	return msg
}

func getCharacterTemplate(data *charPeopleData, link string, info bool) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Name,
		URL:   getLink(link, "character", data.ID),
		Color: orange,
	}

	if !info {
		msg.Description = ellipsis(emptyCheck(data.About), 500)
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
				Value:  emptyCheck(strings.Join(data.Nicknames, ", ")),
				Inline: true,
			},
			{
				Name:   "Japanese",
				Value:  emptyCheck(data.JapaneseName),
				Inline: true,
			},
			{
				Name:   "Favorite",
				Value:  thousands(data.Favorite),
				Inline: true,
			},
			{
				Name:  "About",
				Value: ellipsis(emptyCheck(data.About), 1000),
			},
		}
	}

	return msg
}

func getPeopleTemplate(data *charPeopleData, link string, info bool) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Title: data.Name,
		URL:   getLink(link, "people", data.ID),
		Color: purple,
	}

	if !info {
		msg.Description = ellipsis(emptyCheck(data.More), 500)
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
				Value:  emptyCheck(data.GivenName),
				Inline: true,
			},
			{
				Name:   "Family Name",
				Value:  emptyCheck(data.FamilyName),
				Inline: true,
			},
			{
				Name:   "Alternative Names",
				Value:  emptyCheck(strings.Join(data.AlternativeNames, ", ")),
				Inline: true,
			},
			{
				Name:   "Favorite",
				Value:  thousands(data.Favorite),
				Inline: true,
			},
			{
				Name:   "Birthday",
				Value:  toDate(data.Birthday),
				Inline: true,
			},
			{
				Name:   "Website",
				Value:  emptyCheck(data.Website),
				Inline: true,
			},
			{
				Name:  "More",
				Value: ellipsis(emptyCheck(data.More), 1000),
			},
		}
	}

	return msg
}

var months = []string{
	"",
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func toDate(d date) string {
	if d.Year != 0 {
		if d.Month != 0 {
			if d.Day != 0 {
				return fmt.Sprintf("%v %s %v", d.Day, months[d.Month][:3], d.Year)
			} else {
				return fmt.Sprintf("%s %v", months[d.Month][:3], d.Year)
			}
		} else {
			return strconv.Itoa(d.Year)
		}
	} else {
		return "-"
	}
}

func thousands(num int) string {
	str := strconv.Itoa(num)
	l_str := len(str)
	digits := l_str
	if num < 0 {
		digits--
	}
	commas := (digits+2)/3 - 1
	l_buf := l_str + commas
	var sbuf [32]byte // pre allocate buffer at stack rather than make([]byte,n)
	buf := sbuf[0:l_buf]
	// copy str from the end
	for s_i, b_i, c3 := l_str-1, l_buf-1, 0; ; {
		buf[b_i] = str[s_i]
		if s_i == 0 {
			return string(buf)
		}
		s_i--
		b_i--
		// insert comma every 3 chars
		c3++
		if c3 == 3 && (s_i > 0 || num > 0) {
			buf[b_i] = ','
			b_i--
			c3 = 0
		}
	}
}

func emptyCheck(str string) string {
	if str == "" {
		return "-"
	}
	return str
}

package internal

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type reactionHandler struct {
	api      API
	cache    Cacher
	linkHost string
}

// NewReactionHandler to create new discord reaction handler.
func NewReactionHandler(api API, c Cacher, lh string) *reactionHandler {
	return &reactionHandler{
		api:      api,
		cache:    c,
		linkHost: lh,
	}
}

// Handler to get handler function.
func (h *reactionHandler) Handler() func(*discordgo.Session, *discordgo.MessageReactionAdd) {
	return func(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
		// Ignore bot reaction.
		if s.State.User.ID == m.UserID {
			return
		}

		var cmd cacheModel
		if err := h.cache.Get(getRedisKey(m.MessageID), &cmd); err != nil {
			log.Println(err)
			return
		}

		switch cmd.Commands[0] {
		case "search", "s":
			switch cmd.Commands[1] {
			case "anime", "a":
				h.handleSearchAnime(s, m, cmd)
			case "manga", "m":
				h.handleSearchManga(s, m, cmd)
			case "character", "char", "c":
				h.handleSearchCharacter(s, m, cmd)
			case "people", "p":
				h.handleSearchPeople(s, m, cmd)
			}
		case "anime", "a":
			h.handleGetAnime(s, m, cmd)
		case "manga", "m":
			h.handleGetManga(s, m, cmd)
		case "character", "char", "c":
			h.handleGetCharacter(s, m, cmd)
		case "people", "p":
			h.handleGetPeople(s, m, cmd)
		}

		// Remove user reaction.
		if err := s.MessageReactionRemove(m.ChannelID, m.MessageID, m.Emoji.Name, m.UserID); err != nil {
			log.Println(err)
		}
	}
}

func (h *reactionHandler) handleSearchAnime(s *discordgo.Session, m *discordgo.MessageReactionAdd, c cacheModel) {
	switch m.Emoji.Name {
	case arrowStart:
		if c.Page == 1 {
			return
		}
		c.Page = 1
	case arrowLeft:
		if c.Page == 1 {
			return
		}
		c.Page--
	case arrowRight:
		if c.Page == c.LastPage {
			return
		}
		c.Page++
	case arrowEnd:
		if c.Page == c.LastPage {
			return
		}
		c.Page = c.LastPage
	default:
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchAnime(strings.Join(c.Commands[2:], " "), c.Page)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	_, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, getSearchTemplate(data, c.Page, anime, blue, cnt))
	if err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(m.MessageID), cacheModel{
		Commands: c.Commands,
		Page:     c.Page,
		LastPage: (cnt / dataPerPage) + 1,
	}); err != nil {
		log.Println(err)
	}
}

func (h *reactionHandler) handleSearchManga(s *discordgo.Session, m *discordgo.MessageReactionAdd, c cacheModel) {
	switch m.Emoji.Name {
	case arrowStart:
		if c.Page == 1 {
			return
		}
		c.Page = 1
	case arrowLeft:
		if c.Page == 1 {
			return
		}
		c.Page--
	case arrowRight:
		if c.Page == c.LastPage {
			return
		}
		c.Page++
	case arrowEnd:
		if c.Page == c.LastPage {
			return
		}
		c.Page = c.LastPage
	default:
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchManga(strings.Join(c.Commands[2:], " "), c.Page)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	_, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, getSearchTemplate(data, c.Page, manga, green, cnt))
	if err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(m.MessageID), cacheModel{
		Commands: c.Commands,
		Page:     c.Page,
		LastPage: (cnt / dataPerPage) + 1,
	}); err != nil {
		log.Println(err)
	}
}

func (h *reactionHandler) handleSearchCharacter(s *discordgo.Session, m *discordgo.MessageReactionAdd, c cacheModel) {
	switch m.Emoji.Name {
	case arrowStart:
		if c.Page == 1 {
			return
		}
		c.Page = 1
	case arrowLeft:
		if c.Page == 1 {
			return
		}
		c.Page--
	case arrowRight:
		if c.Page == c.LastPage {
			return
		}
		c.Page++
	case arrowEnd:
		if c.Page == c.LastPage {
			return
		}
		c.Page = c.LastPage
	default:
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchCharacter(strings.Join(c.Commands[2:], " "), c.Page)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	_, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, getSearchTemplate(data, c.Page, character, orange, cnt))
	if err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(m.MessageID), cacheModel{
		Commands: c.Commands,
		Page:     c.Page,
		LastPage: (cnt / dataPerPage) + 1,
	}); err != nil {
		log.Println(err)
	}
}

func (h *reactionHandler) handleSearchPeople(s *discordgo.Session, m *discordgo.MessageReactionAdd, c cacheModel) {
	switch m.Emoji.Name {
	case arrowStart:
		if c.Page == 1 {
			return
		}
		c.Page = 1
	case arrowLeft:
		if c.Page == 1 {
			return
		}
		c.Page--
	case arrowRight:
		if c.Page == c.LastPage {
			return
		}
		c.Page++
	case arrowEnd:
		if c.Page == c.LastPage {
			return
		}
		c.Page = c.LastPage
	default:
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchPeople(strings.Join(c.Commands[2:], " "), c.Page)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	_, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, getSearchTemplate(data, c.Page, people, purple, cnt))
	if err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(m.MessageID), cacheModel{
		Commands: c.Commands,
		Page:     c.Page,
		LastPage: (cnt / dataPerPage) + 1,
	}); err != nil {
		log.Println(err)
	}
}

func (h *reactionHandler) handleGetAnime(s *discordgo.Session, m *discordgo.MessageReactionAdd, c cacheModel) {
	if m.Emoji.Name != info {
		return
	}

	c.Info = !c.Info

	// Get data.
	id, _ := strconv.Atoi(c.Commands[1])
	data, _, err := h.api.GetAnime(id)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	if _, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, getAnimeTemplate(data, h.linkHost, c.Info)); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(m.MessageID), cacheModel{
		Commands: c.Commands,
		Info:     c.Info,
	}); err != nil {
		log.Println(err)
	}
}

func (h *reactionHandler) handleGetManga(s *discordgo.Session, m *discordgo.MessageReactionAdd, c cacheModel) {
	if m.Emoji.Name != info {
		return
	}

	c.Info = !c.Info

	// Get data.
	id, _ := strconv.Atoi(c.Commands[1])
	data, _, err := h.api.GetManga(id)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	if _, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, getMangaTemplate(data, h.linkHost, c.Info)); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(m.MessageID), cacheModel{
		Commands: c.Commands,
		Info:     c.Info,
	}); err != nil {
		log.Println(err)
	}
}

func (h *reactionHandler) handleGetCharacter(s *discordgo.Session, m *discordgo.MessageReactionAdd, c cacheModel) {
	if m.Emoji.Name != info {
		return
	}

	c.Info = !c.Info

	// Get data.
	id, _ := strconv.Atoi(c.Commands[1])
	data, _, err := h.api.GetCharacter(id)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	if _, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, getCharacterTemplate(data, h.linkHost, c.Info)); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(m.MessageID), cacheModel{
		Commands: c.Commands,
		Info:     c.Info,
	}); err != nil {
		log.Println(err)
	}
}

func (h *reactionHandler) handleGetPeople(s *discordgo.Session, m *discordgo.MessageReactionAdd, c cacheModel) {
	if m.Emoji.Name != info {
		return
	}

	c.Info = !c.Info

	// Get data.
	id, _ := strconv.Atoi(c.Commands[1])
	data, _, err := h.api.GetPeople(id)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	if _, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, getPeopleTemplate(data, h.linkHost, c.Info)); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(m.MessageID), cacheModel{
		Commands: c.Commands,
		Info:     c.Info,
	}); err != nil {
		log.Println(err)
	}
}

package internal

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type messageHandler struct {
	api      API
	cache    Cacher
	prefix   string
	linkHost string
}

// NewMessageHandler to create new discord message handler.
func NewMessageHandler(api API, c Cacher, prefix string, lh string) *messageHandler {
	return &messageHandler{
		api:      api,
		cache:    c,
		prefix:   prefix,
		linkHost: lh,
	}
}

// Handler to get handler function.
func (h *messageHandler) Handler() func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore all messages created by the bot itself.
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Command and prefix check.
		if h.prefixCheck(m.Content) {
			return
		}

		// Remove prefix.
		m.Content = h.cleanPrefix(m.Content)

		// Get arguments.
		r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`)
		args := r.FindAllString(m.Content, -1)

		switch args[0] {
		case "ping":
			h.handlePing(s, m)
		case "help", "h":
			h.handleHelp(s, m)
		case "search", "s":
			if len(args) <= 2 {
				h.handleInvalid(s, m, invalidContent)
				return
			}
			switch args[1] {
			case "anime", "a":
				h.handleSearchAnime(s, m, args)
			case "manga", "m":
				h.handleSearchManga(s, m, args)
			case "character", "char", "c":
				h.handleSearchCharacter(s, m, args)
			case "people", "p":
				h.handleSearchPeople(s, m, args)
			default:
				h.handleInvalid(s, m, invalidContent)
				return
			}
		case "anime", "a":
			if len(args) == 1 || len(args) > 2 {
				h.handleInvalid(s, m, invalidContent)
				return
			}
			h.handleGetAnime(s, m, args)
		case "manga", "m":
			if len(args) == 1 || len(args) > 2 {
				h.handleInvalid(s, m, invalidContent)
				return
			}
			h.handleGetManga(s, m, args)
		case "character", "char", "c":
			if len(args) == 1 || len(args) > 2 {
				h.handleInvalid(s, m, invalidContent)
				return
			}
			h.handleGetCharacter(s, m, args)
		case "people", "p":
			if len(args) == 1 || len(args) > 2 {
				h.handleInvalid(s, m, invalidContent)
				return
			}
			h.handleGetPeople(s, m, args)
		default:
			return
		}
	}
}

func (h *messageHandler) prefixCheck(cmd string) bool {
	return len(cmd) <= len(h.prefix) || cmd[:len(h.prefix)] != h.prefix
}

func (h *messageHandler) cleanPrefix(cmd string) string {
	return strings.TrimSpace(cmd[len(h.prefix):])
}

func (h *messageHandler) handleInvalid(s *discordgo.Session, m *discordgo.MessageCreate, str string) {
	if _, err := s.ChannelMessageSend(m.ChannelID, str); err != nil {
		log.Println(err)
	}
}

func (h *messageHandler) handlePing(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSend(m.ChannelID, "pong"); err != nil {
		log.Println(err)
	}
}

func (h *messageHandler) handleHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Help",
		Description: helpContent,
		Color:       greyLight,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  searchCmd,
				Value: searchContent,
			},
			{
				Name:  animeCmd,
				Value: animeContent,
			},
			{
				Name:  mangaCmd,
				Value: mangaContent,
			},
			{
				Name:  charCmd,
				Value: charContent,
			},
			{
				Name:  peopleCmd,
				Value: peopleContent,
			},
		},
	}); err != nil {
		log.Println(err)
	}
}

func (h *messageHandler) handleSearchAnime(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args[2]) < 3 {
		h.handleInvalid(s, m, search3Letter)
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchAnime(strings.Join(args[2:], " "), 1)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, getSearchTemplate(data, 1, anime, blue, cnt))
	if err != nil {
		log.Println(err)
		return
	}

	// Add pagination reaction.
	if cnt > len(data) {
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowStart); err != nil {
			log.Println(err)
		}
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowLeft); err != nil {
			log.Println(err)
		}
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowRight); err != nil {
			log.Println(err)
		}
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowEnd); err != nil {
			log.Println(err)
		}
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(msg.ID), cacheModel{
		Commands: args,
		Page:     1,
		LastPage: (cnt / dataPerPage) + 1,
	}); err != nil {
		log.Println(err)
	}
}

func (h *messageHandler) handleSearchManga(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args[2]) < 3 {
		h.handleInvalid(s, m, search3Letter)
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchManga(strings.Join(args[2:], " "), 1)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, getSearchTemplate(data, 1, manga, green, cnt))
	if err != nil {
		log.Println(err)
		return
	}
	if err != nil {
		log.Println(err)
		return
	}

	// Add pagination reaction.
	if cnt > len(data) {
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowStart); err != nil {
			log.Println(err)
		}
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowLeft); err != nil {
			log.Println(err)
		}
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowRight); err != nil {
			log.Println(err)
		}
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowEnd); err != nil {
			log.Println(err)
		}
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(msg.ID), cacheModel{
		Commands: args,
		Page:     1,
		LastPage: (cnt / dataPerPage) + 1,
	}); err != nil {
		log.Println(err)
	}
}

func (h *messageHandler) handleSearchCharacter(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args[2]) < 3 {
		h.handleInvalid(s, m, search3Letter)
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchCharacter(strings.Join(args[2:], " "), 1)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, getSearchTemplate(data, 1, character, orange, cnt))
	if err != nil {
		log.Println(err)
		return
	}
	if err != nil {
		log.Println(err)
		return
	}

	// Add pagination reaction.
	if cnt > len(data) {
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowStart); err != nil {
			log.Println(err)
		}
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowLeft); err != nil {
			log.Println(err)
		}
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowRight); err != nil {
			log.Println(err)
		}
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowEnd); err != nil {
			log.Println(err)
		}
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(msg.ID), cacheModel{
		Commands: args,
		Page:     1,
		LastPage: (cnt / dataPerPage) + 1,
	}); err != nil {
		log.Println(err)
	}
}

func (h *messageHandler) handleSearchPeople(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args[2]) < 3 {
		h.handleInvalid(s, m, search3Letter)
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchPeople(strings.Join(args[2:], " "), 1)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, getSearchTemplate(data, 1, people, purple, cnt))
	if err != nil {
		log.Println(err)
		return
	}
	if err != nil {
		log.Println(err)
		return
	}

	// Add pagination reaction.
	if cnt > len(data) {
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowStart); err != nil {
			log.Println(err)
		}
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowLeft); err != nil {
			log.Println(err)
		}
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowRight); err != nil {
			log.Println(err)
		}
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, arrowEnd); err != nil {
			log.Println(err)
		}
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(msg.ID), cacheModel{
		Commands: args,
		Page:     1,
		LastPage: (cnt / dataPerPage) + 1,
	}); err != nil {
		log.Println(err)
	}
}

func (h *messageHandler) handleGetAnime(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	id, err := strconv.Atoi(args[1])
	if err != nil || id <= 0 {
		h.handleInvalid(s, m, invalidID)
		return
	}

	// Get data.
	data, code, err := h.api.GetAnime(id)
	if err != nil {
		log.Println(err)
		if code == http.StatusNotFound {
			h.handleInvalid(s, m, invalidID)
			return
		}
		return
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, getAnimeTemplate(data, h.linkHost, false))
	if err != nil {
		log.Println(err)
		return
	}

	// Add reaction.
	if err = s.MessageReactionAdd(m.ChannelID, msg.ID, info); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(msg.ID), cacheModel{
		Commands: args,
		Info:     false,
	}); err != nil {
		log.Println(err)
	}
}

func (h *messageHandler) handleGetManga(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	id, err := strconv.Atoi(args[1])
	if err != nil || id <= 0 {
		h.handleInvalid(s, m, invalidID)
		return
	}

	// Get data.
	data, code, err := h.api.GetManga(id)
	if err != nil {
		log.Println(err)
		if code == http.StatusNotFound {
			h.handleInvalid(s, m, invalidID)
			return
		}
		return
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, getMangaTemplate(data, h.linkHost, false))
	if err != nil {
		log.Println(err)
		return
	}

	// Add reaction.
	if err = s.MessageReactionAdd(m.ChannelID, msg.ID, info); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(msg.ID), cacheModel{
		Commands: args,
		Info:     false,
	}); err != nil {
		log.Println(err)
	}
}

func (h *messageHandler) handleGetCharacter(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	id, err := strconv.Atoi(args[1])
	if err != nil || id <= 0 {
		h.handleInvalid(s, m, invalidID)
		return
	}

	// Get data.
	data, code, err := h.api.GetCharacter(id)
	if err != nil {
		log.Println(err)
		if code == http.StatusNotFound {
			h.handleInvalid(s, m, invalidID)
			return
		}
		return
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, getCharacterTemplate(data, h.linkHost, false))
	if err != nil {
		log.Println(err)
		return
	}

	// Add reaction.
	if err = s.MessageReactionAdd(m.ChannelID, msg.ID, info); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(msg.ID), cacheModel{
		Commands: args,
		Info:     false,
	}); err != nil {
		log.Println(err)
	}
}

func (h *messageHandler) handleGetPeople(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	id, err := strconv.Atoi(args[1])
	if err != nil || id <= 0 {
		h.handleInvalid(s, m, invalidID)
		return
	}

	// Get data.
	data, code, err := h.api.GetPeople(id)
	if err != nil {
		log.Println(err)
		if code == http.StatusNotFound {
			h.handleInvalid(s, m, invalidID)
			return
		}
		return
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, getPeopleTemplate(data, h.linkHost, false))
	if err != nil {
		log.Println(err)
		return
	}

	// Add reaction.
	if err = s.MessageReactionAdd(m.ChannelID, msg.ID, info); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(getRedisKey(msg.ID), cacheModel{
		Commands: args,
		Info:     false,
	}); err != nil {
		log.Println(err)
	}
}

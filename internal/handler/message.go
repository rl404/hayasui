package handler

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rl404/hayasui/internal/api"
	"github.com/rl404/hayasui/internal/cache"
	"github.com/rl404/hayasui/internal/constant"
	"github.com/rl404/hayasui/internal/model"
)

// MessageHandler for handling message.
type MessageHandler struct {
	api      api.API
	cache    cache.Cacher
	prefix   string
	template Templater
}

// NewMessageHandler to create new discord message handler.
func NewMessageHandler(api api.API, c cache.Cacher, prefix string, lh string) *MessageHandler {
	return &MessageHandler{
		api:      api,
		cache:    c,
		prefix:   prefix,
		template: newTemplate(lh),
	}
}

// Handler to get handler function.
func (h *MessageHandler) Handler() func(*discordgo.Session, *discordgo.MessageCreate) {
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
				h.handleInvalid(s, m, constant.MsgInvalid)
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
				h.handleInvalid(s, m, constant.MsgInvalid)
				return
			}
		case "anime", "a":
			if len(args) == 1 || len(args) > 2 {
				h.handleInvalid(s, m, constant.MsgInvalid)
				return
			}
			h.handleGetAnime(s, m, args)
		case "manga", "m":
			if len(args) == 1 || len(args) > 2 {
				h.handleInvalid(s, m, constant.MsgInvalid)
				return
			}
			h.handleGetManga(s, m, args)
		case "character", "char", "c":
			if len(args) == 1 || len(args) > 2 {
				h.handleInvalid(s, m, constant.MsgInvalid)
				return
			}
			h.handleGetCharacter(s, m, args)
		case "people", "p":
			if len(args) == 1 || len(args) > 2 {
				h.handleInvalid(s, m, constant.MsgInvalid)
				return
			}
			h.handleGetPeople(s, m, args)
		default:
			return
		}
	}
}

func (h *MessageHandler) prefixCheck(cmd string) bool {
	return len(cmd) <= len(h.prefix) || cmd[:len(h.prefix)] != h.prefix
}

func (h *MessageHandler) cleanPrefix(cmd string) string {
	return strings.TrimSpace(cmd[len(h.prefix):])
}

func (h *MessageHandler) handleInvalid(s *discordgo.Session, m *discordgo.MessageCreate, str string) {
	if _, err := s.ChannelMessageSend(m.ChannelID, str); err != nil {
		log.Println(err)
	}
}

func (h *MessageHandler) handlePing(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSend(m.ChannelID, "pong"); err != nil {
		log.Println(err)
	}
}

func (h *MessageHandler) handleHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, h.template.GetHelp()); err != nil {
		log.Println(err)
	}
}

func (h *MessageHandler) handleSearchAnime(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args[2]) < 3 {
		h.handleInvalid(s, m, constant.MsgSearch3Letter)
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchAnime(strings.Join(args[2:], " "), 1)
	if err != nil {
		log.Println(err)
		return
	}

	// Command model.
	cmd := model.Command{
		Commands: args,
		Page:     1,
		LastPage: (cnt / constant.DataPerPage) + 1,
		Type:     0,
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, h.template.GetSearchAnime(data, cmd))
	if err != nil {
		log.Println(err)
		return
	}

	// Add pagination reaction.
	for _, r := range constant.ReactionPaginationWithInfo {
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, r); err != nil {
			log.Println(err)
		}
	}

	// Save to redis.
	if err = h.cache.Set(cmd, "msg", msg.ID); err != nil {
		log.Println(err)
	}
}

func (h *MessageHandler) handleSearchManga(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args[2]) < 3 {
		h.handleInvalid(s, m, constant.MsgSearch3Letter)
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchManga(strings.Join(args[2:], " "), 1)
	if err != nil {
		log.Println(err)
		return
	}

	// Command model.
	cmd := model.Command{
		Commands: args,
		Page:     1,
		LastPage: (cnt / constant.DataPerPage) + 1,
		Type:     0,
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, h.template.GetSearchManga(data, cmd))
	if err != nil {
		log.Println(err)
		return
	}

	// Add pagination reaction.
	for _, r := range constant.ReactionPaginationWithInfo {
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, r); err != nil {
			log.Println(err)
		}
	}

	// Save to redis.
	if err = h.cache.Set(cmd, "msg", msg.ID); err != nil {
		log.Println(err)
	}
}

func (h *MessageHandler) handleSearchCharacter(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args[2]) < 3 {
		h.handleInvalid(s, m, constant.MsgSearch3Letter)
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchCharacter(strings.Join(args[2:], " "), 1)
	if err != nil {
		log.Println(err)
		return
	}

	// Command model.
	cmd := model.Command{
		Commands: args,
		Page:     1,
		LastPage: (cnt / constant.DataPerPage) + 1,
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, h.template.GetSearchCharacter(data, cmd))
	if err != nil {
		log.Println(err)
		return
	}

	// Add pagination reaction.
	for _, r := range constant.ReactionPagination {
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, r); err != nil {
			log.Println(err)
		}
	}

	// Save to redis.
	if err = h.cache.Set(cmd, "msg", msg.ID); err != nil {
		log.Println(err)
	}
}

func (h *MessageHandler) handleSearchPeople(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args[2]) < 3 {
		h.handleInvalid(s, m, constant.MsgSearch3Letter)
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchPeople(strings.Join(args[2:], " "), 1)
	if err != nil {
		log.Println(err)
		return
	}

	// Command model.
	cmd := model.Command{
		Commands: args,
		Page:     1,
		LastPage: (cnt / constant.DataPerPage) + 1,
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, h.template.GetSearchPeople(data, cmd))
	if err != nil {
		log.Println(err)
		return
	}

	// Add pagination reaction.
	for _, r := range constant.ReactionPagination {
		if err = s.MessageReactionAdd(m.ChannelID, msg.ID, r); err != nil {
			log.Println(err)
		}
	}

	// Save to redis.
	if err = h.cache.Set(cmd, "msg", msg.ID); err != nil {
		log.Println(err)
	}
}

func (h *MessageHandler) handleGetAnime(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	id, err := strconv.Atoi(args[1])
	if err != nil || id <= 0 {
		h.handleInvalid(s, m, constant.MsgInvalidID)
		return
	}

	// Get data.
	data, err := h.api.GetAnime(id)
	if err != nil {
		h.handleInvalid(s, m, err.Error())
		return
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, h.template.GetAnime(data, false))
	if err != nil {
		log.Println(err)
		return
	}

	// Add reaction.
	if err = s.MessageReactionAdd(m.ChannelID, msg.ID, constant.ReactionInfo); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(model.Command{
		Commands: args,
		Info:     false,
	}, "msg", msg.ID); err != nil {
		log.Println(err)
	}
}

func (h *MessageHandler) handleGetManga(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	id, err := strconv.Atoi(args[1])
	if err != nil || id <= 0 {
		h.handleInvalid(s, m, constant.MsgInvalidID)
		return
	}

	// Get data.
	data, err := h.api.GetManga(id)
	if err != nil {
		h.handleInvalid(s, m, err.Error())
		return
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, h.template.GetManga(data, false))
	if err != nil {
		log.Println(err)
		return
	}

	// Add reaction.
	if err = s.MessageReactionAdd(m.ChannelID, msg.ID, constant.ReactionInfo); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(model.Command{
		Commands: args,
		Info:     false,
	}, "msg", msg.ID); err != nil {
		log.Println(err)
	}
}

func (h *MessageHandler) handleGetCharacter(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	id, err := strconv.Atoi(args[1])
	if err != nil || id <= 0 {
		h.handleInvalid(s, m, constant.MsgInvalidID)
		return
	}

	// Get data.
	data, err := h.api.GetCharacter(id)
	if err != nil {
		h.handleInvalid(s, m, err.Error())
		return
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, h.template.GetCharacter(data, false))
	if err != nil {
		log.Println(err)
		return
	}

	// Add reaction.
	if err = s.MessageReactionAdd(m.ChannelID, msg.ID, constant.ReactionInfo); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(model.Command{
		Commands: args,
		Info:     false,
	}, "msg", msg.ID); err != nil {
		log.Println(err)
	}
}

func (h *MessageHandler) handleGetPeople(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	id, err := strconv.Atoi(args[1])
	if err != nil || id <= 0 {
		h.handleInvalid(s, m, constant.MsgInvalidID)
		return
	}

	// Get data.
	data, err := h.api.GetPeople(id)
	if err != nil {
		h.handleInvalid(s, m, err.Error())
		return
	}

	// Send message.
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, h.template.GetPeople(data, false))
	if err != nil {
		log.Println(err)
		return
	}

	// Add reaction.
	if err = s.MessageReactionAdd(m.ChannelID, msg.ID, constant.ReactionInfo); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(model.Command{
		Commands: args,
		Info:     false,
	}, "msg", msg.ID); err != nil {
		log.Println(err)
	}
}

package handler

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rl404/hayasui/internal/api"
	"github.com/rl404/hayasui/internal/cache"
	"github.com/rl404/hayasui/internal/constant"
	"github.com/rl404/hayasui/internal/model"
)

// ReactionHandler for handling reaction.
type ReactionHandler struct {
	api      api.API
	cache    cache.Cacher
	template Templater
}

// NewReactionHandler to create new discord reaction handler.
func NewReactionHandler(api api.API, c cache.Cacher, lh string) *ReactionHandler {
	return &ReactionHandler{
		api:      api,
		cache:    c,
		template: newTemplate(lh),
	}
}

// Handler to get handler function.
func (h *ReactionHandler) Handler() func(*discordgo.Session, *discordgo.MessageReactionAdd) {
	return func(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
		// Ignore bot reaction.
		if s.State.User.ID == m.UserID {
			return
		}

		var cmd model.Command
		if err := h.cache.Get(&cmd, "msg", m.MessageID); err != nil {
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

func (h *ReactionHandler) handleSearchAnime(s *discordgo.Session, m *discordgo.MessageReactionAdd, c model.Command) {
	switch m.Emoji.Name {
	case constant.ReactionArrowStart:
		if c.Page == 1 {
			return
		}
		c.Page = 1
	case constant.ReactionArrowLeft:
		if c.Page == 1 {
			return
		}
		c.Page--
	case constant.ReactionArrowRight:
		if c.Page == c.LastPage {
			return
		}
		c.Page++
	case constant.ReactionArrowEnd:
		if c.Page == c.LastPage {
			return
		}
		c.Page = c.LastPage
	case constant.ReactionInfo:
		c.Type++
		if c.Type > 2 {
			c.Type = 0
		}
	default:
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchAnime(strings.Join(c.Commands[2:], " "), c.Page)
	if err != nil {
		log.Println(err)
		return
	}

	// Command model.
	c = model.Command{
		Commands: c.Commands,
		Page:     c.Page,
		LastPage: (cnt / constant.DataPerPage) + 1,
		Type:     c.Type,
	}

	// Send message.
	if _, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, h.template.GetSearchAnime(data, c)); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(c, "msg", m.MessageID); err != nil {
		log.Println(err)
	}
}

func (h *ReactionHandler) handleSearchManga(s *discordgo.Session, m *discordgo.MessageReactionAdd, c model.Command) {
	switch m.Emoji.Name {
	case constant.ReactionArrowStart:
		if c.Page == 1 {
			return
		}
		c.Page = 1
	case constant.ReactionArrowLeft:
		if c.Page == 1 {
			return
		}
		c.Page--
	case constant.ReactionArrowRight:
		if c.Page == c.LastPage {
			return
		}
		c.Page++
	case constant.ReactionArrowEnd:
		if c.Page == c.LastPage {
			return
		}
		c.Page = c.LastPage
	case constant.ReactionInfo:
		c.Type++
		if c.Type > 2 {
			c.Type = 0
		}
	default:
		return
	}

	// Get data.
	data, cnt, err := h.api.SearchManga(strings.Join(c.Commands[2:], " "), c.Page)
	if err != nil {
		log.Println(err)
		return
	}

	// Command model.
	c = model.Command{
		Commands: c.Commands,
		Page:     c.Page,
		LastPage: (cnt / constant.DataPerPage) + 1,
		Type:     c.Type,
	}

	// Send message.
	if _, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, h.template.GetSearchManga(data, c)); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(c, "msg", m.MessageID); err != nil {
		log.Println(err)
	}
}

func (h *ReactionHandler) handleSearchCharacter(s *discordgo.Session, m *discordgo.MessageReactionAdd, c model.Command) {
	switch m.Emoji.Name {
	case constant.ReactionArrowStart:
		if c.Page == 1 {
			return
		}
		c.Page = 1
	case constant.ReactionArrowLeft:
		if c.Page == 1 {
			return
		}
		c.Page--
	case constant.ReactionArrowRight:
		if c.Page == c.LastPage {
			return
		}
		c.Page++
	case constant.ReactionArrowEnd:
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

	// Command model.
	c = model.Command{
		Commands: c.Commands,
		Page:     c.Page,
		LastPage: (cnt / constant.DataPerPage) + 1,
	}

	// Send message.
	if _, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, h.template.GetSearchCharacter(data, c)); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(c, "msg", m.MessageID); err != nil {
		log.Println(err)
	}
}

func (h *ReactionHandler) handleSearchPeople(s *discordgo.Session, m *discordgo.MessageReactionAdd, c model.Command) {
	switch m.Emoji.Name {
	case constant.ReactionArrowStart:
		if c.Page == 1 {
			return
		}
		c.Page = 1
	case constant.ReactionArrowLeft:
		if c.Page == 1 {
			return
		}
		c.Page--
	case constant.ReactionArrowRight:
		if c.Page == c.LastPage {
			return
		}
		c.Page++
	case constant.ReactionArrowEnd:
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

	// Command model.
	c = model.Command{
		Commands: c.Commands,
		Page:     c.Page,
		LastPage: (cnt / constant.DataPerPage) + 1,
		Type:     c.Type,
	}

	// Send message.
	if _, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, h.template.GetSearchPeople(data, c)); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(c, "msg", m.MessageID); err != nil {
		log.Println(err)
	}
}

func (h *ReactionHandler) handleGetAnime(s *discordgo.Session, m *discordgo.MessageReactionAdd, c model.Command) {
	if m.Emoji.Name != constant.ReactionInfo {
		return
	}

	c.Info = !c.Info

	// Get data.
	id, _ := strconv.Atoi(c.Commands[1])
	data, err := h.api.GetAnime(id)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	if _, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, h.template.GetAnime(data, c.Info)); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(model.Command{
		Commands: c.Commands,
		Info:     c.Info,
	}, "msg", m.MessageID); err != nil {
		log.Println(err)
	}
}

func (h *ReactionHandler) handleGetManga(s *discordgo.Session, m *discordgo.MessageReactionAdd, c model.Command) {
	if m.Emoji.Name != constant.ReactionInfo {
		return
	}

	c.Info = !c.Info

	// Get data.
	id, _ := strconv.Atoi(c.Commands[1])
	data, err := h.api.GetManga(id)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	if _, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, h.template.GetManga(data, c.Info)); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(model.Command{
		Commands: c.Commands,
		Info:     c.Info,
	}, "msg", m.MessageID); err != nil {
		log.Println(err)
	}
}

func (h *ReactionHandler) handleGetCharacter(s *discordgo.Session, m *discordgo.MessageReactionAdd, c model.Command) {
	if m.Emoji.Name != constant.ReactionInfo {
		return
	}

	c.Info = !c.Info

	// Get data.
	id, _ := strconv.Atoi(c.Commands[1])
	data, err := h.api.GetCharacter(id)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	if _, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, h.template.GetCharacter(data, c.Info)); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(model.Command{
		Commands: c.Commands,
		Info:     c.Info,
	}, "msg", m.MessageID); err != nil {
		log.Println(err)
	}
}

func (h *ReactionHandler) handleGetPeople(s *discordgo.Session, m *discordgo.MessageReactionAdd, c model.Command) {
	if m.Emoji.Name != constant.ReactionInfo {
		return
	}

	c.Info = !c.Info

	// Get data.
	id, _ := strconv.Atoi(c.Commands[1])
	data, err := h.api.GetPeople(id)
	if err != nil {
		log.Println(err)
		return
	}

	// Send message.
	if _, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, h.template.GetPeople(data, c.Info)); err != nil {
		log.Println(err)
		return
	}

	// Save to redis.
	if err = h.cache.Set(model.Command{
		Commands: c.Commands,
		Info:     c.Info,
	}, "msg", m.MessageID); err != nil {
		log.Println(err)
	}
}

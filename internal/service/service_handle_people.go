package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	reactionEntity "github.com/rl404/hayasui/internal/domain/reaction/entity"
	"github.com/rl404/hayasui/internal/domain/template/entity"
	"github.com/rl404/hayasui/internal/errors"
)

// HandlePeople to handle people.
func (s *service) HandlePeople(ctx context.Context, m *discordgo.MessageCreate, args []string) error {
	if len(args) != 2 {
		return errors.Wrap(ctx, s.handleInvalid(ctx, m.ChannelID))
	}

	id, err := strconv.Atoi(args[1])
	if err != nil || id <= 0 {
		return errors.Wrap(ctx, s.handleInvalidID(ctx, m.ChannelID))
	}

	// Get data.
	data, err := s.anime.GetPeople(ctx, id)
	if err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	// Send message.
	msg, err := s.discord.SendMessageEmbed(ctx, m.ChannelID, s.template.GetPeople(entity.People{
		URL:   data.URL,
		Name:  data.Name,
		About: data.About,
		Image: data.Image,
	}, entity.InfoSimple))
	if err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	// Add reaction.
	if err := s.discord.AddMessageReaction(ctx, m.ChannelID, msg, entity.ReactionInfo); err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	// Save reaction.
	if err := s.reaction.SetCommand(ctx, msg, reactionEntity.Command{
		Command: args[0],
		ID:      data.ID,
		Info:    entity.InfoSimple,
	}); err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	return nil
}

// HandlePeopleReaction to handle people reaction.
func (s *service) HandlePeopleReaction(ctx context.Context, m *discordgo.MessageReactionAdd, cmd reactionEntity.Command) error {
	if m.Emoji.Name != entity.ReactionInfo {
		return nil
	}

	cmd.Info++
	if cmd.Info > entity.InfoMore {
		cmd.Info = entity.InfoSimple
	}

	// Get data.
	data, err := s.anime.GetPeople(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	// Edit message.
	msg, err := s.discord.EditMessageEmbed(ctx, m.ChannelID, m.MessageID, s.template.GetPeople(entity.People{
		URL:              data.URL,
		Name:             data.Name,
		AlternativeNames: data.AlternativeNames,
		Birthday: entity.Date{
			Year:  data.Birthday.Year,
			Month: data.Birthday.Month,
			Day:   data.Birthday.Day,
		},
		Favorite: data.Favorite,
		About:    data.About,
		Image:    data.Image,
	}, cmd.Info))
	if err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	// Save reaction.
	if err := s.reaction.SetCommand(ctx, msg, cmd); err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	return nil
}

func (s *service) handleSearchPeople(ctx context.Context, m *discordgo.MessageCreate, args []string) error {
	// Get data.
	data, cnt, err := s.anime.SearchPeople(ctx, strings.Join(args[2:], " "), 1)
	if err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	templateDatas := make([]entity.Search, len(data))
	for i, d := range data {
		templateDatas[i] = entity.Search{
			ID:   d.ID,
			Name: d.Name,
		}
	}

	lastPage := (cnt / entity.DataPerPage) + 1

	// Send message.
	msg, err := s.discord.SendMessageEmbed(ctx, m.ChannelID, s.template.GetSearch(templateDatas, entity.TypePeople, entity.InfoSimple, 1, lastPage))
	if err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	// Add reaction.
	for _, r := range entity.ReactionPagination {
		if err := s.discord.AddMessageReaction(ctx, m.ChannelID, msg, r); err != nil {
			return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
		}
	}

	// Save reaction.
	if err := s.reaction.SetCommand(ctx, msg, reactionEntity.Command{
		Command:  args[0],
		Arg:      args[1],
		Query:    strings.Join(args[2:], " "),
		Page:     1,
		LastPage: lastPage,
		Info:     entity.InfoSimple,
	}); err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	return nil
}

func (s *service) handleSearchPeopleReaction(ctx context.Context, m *discordgo.MessageReactionAdd, cmd reactionEntity.Command) error {
	switch m.Emoji.Name {
	case entity.ReactionArrowStart:
		if cmd.Page == 1 {
			return nil
		}
		cmd.Page = 1
	case entity.ReactionArrowLeft:
		if cmd.Page == 1 {
			return nil
		}
		cmd.Page--
	case entity.ReactionArrowRight:
		if cmd.Page == cmd.LastPage {
			return nil
		}
		cmd.Page++
	case entity.ReactionArrowEnd:
		if cmd.Page == cmd.LastPage {
			return nil
		}
		cmd.Page = cmd.LastPage
	default:
		return nil
	}

	// Get data.
	data, cnt, err := s.anime.SearchPeople(ctx, cmd.Query, cmd.Page)
	if err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	templateDatas := make([]entity.Search, len(data))
	for i, d := range data {
		templateDatas[i] = entity.Search{
			ID:   d.ID,
			Name: d.Name,
		}
	}

	cmd.LastPage = (cnt / entity.DataPerPage) + 1

	// Edit message.
	msg, err := s.discord.EditMessageEmbed(ctx, m.ChannelID, m.MessageID, s.template.GetSearch(templateDatas, entity.TypePeople, cmd.Info, cmd.Page, cmd.LastPage))
	if err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	// Save reaction.
	if err := s.reaction.SetCommand(ctx, msg, cmd); err != nil {
		return errors.Wrap(ctx, s.handleError(ctx, m.ChannelID, errors.Wrap(ctx, err)))
	}

	return nil
}

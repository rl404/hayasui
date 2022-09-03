package service

import (
	"context"

	"github.com/bwmarrin/discordgo"

	animeRepository "github.com/rl404/hayasui/internal/domain/anime/repository"
	discordRepository "github.com/rl404/hayasui/internal/domain/discord/repository"
	"github.com/rl404/hayasui/internal/domain/reaction/entity"
	reactionRepository "github.com/rl404/hayasui/internal/domain/reaction/repository"
	templateRepository "github.com/rl404/hayasui/internal/domain/template/repository"
)

// Service contains functions for service.
type Service interface {
	Run() error
	Stop() error

	RegisterReadyHandler(func(*discordgo.Session, *discordgo.Ready))
	RegisterMessageHandler(func(*discordgo.Session, *discordgo.MessageCreate))
	RegisterReactionHandler(func(*discordgo.Session, *discordgo.MessageReactionAdd))

	HandlePing(ctx context.Context, m *discordgo.MessageCreate) error
	HandleHelp(ctx context.Context, m *discordgo.MessageCreate) error
	HandleSearch(ctx context.Context, m *discordgo.MessageCreate, args []string) error
	HandleAnime(ctx context.Context, m *discordgo.MessageCreate, args []string) error
	HandleManga(ctx context.Context, m *discordgo.MessageCreate, args []string) error
	HandleCharacter(ctx context.Context, m *discordgo.MessageCreate, args []string) error
	HandlePeople(ctx context.Context, m *discordgo.MessageCreate, args []string) error

	GetReaction(ctx context.Context, m *discordgo.MessageReactionAdd) (*entity.Command, error)

	HandleSearchReaction(ctx context.Context, m *discordgo.MessageReactionAdd, cmd entity.Command) error
	HandleAnimeReaction(ctx context.Context, m *discordgo.MessageReactionAdd, cmd entity.Command) error
	HandleMangaReaction(ctx context.Context, m *discordgo.MessageReactionAdd, cmd entity.Command) error
	HandleCharacterReaction(ctx context.Context, m *discordgo.MessageReactionAdd, cmd entity.Command) error
	HandlePeopleReaction(ctx context.Context, m *discordgo.MessageReactionAdd, cmd entity.Command) error
}

type service struct {
	discord  discordRepository.Repository
	template templateRepository.Repository
	anime    animeRepository.Repository
	reaction reactionRepository.Repository
}

// New to create new service.
func New(
	discord discordRepository.Repository,
	template templateRepository.Repository,
	anime animeRepository.Repository,
	reaction reactionRepository.Repository,
) Service {
	return &service{
		discord:  discord,
		template: template,
		anime:    anime,
		reaction: reaction,
	}
}

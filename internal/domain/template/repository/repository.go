package repository

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rl404/hayasui/internal/domain/template/entity"
)

// Repository contains functions for template domain.
type Repository interface {
	GetHelp() *discordgo.MessageEmbed
	GetInvalid() string
	GetAnime(data entity.Anime, info int8) *discordgo.MessageEmbed
	GetManga(data entity.Manga, info int8) *discordgo.MessageEmbed
	GetCharacter(data entity.Character, info int8) *discordgo.MessageEmbed
	GetPeople(data entity.People, info int8) *discordgo.MessageEmbed
	GetSearch(data []entity.Search, _type entity.SearchType, info int8, page, lastPage int) *discordgo.MessageEmbed
}

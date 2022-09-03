package entity

// Color.
const (
	ColorGreyLight = 12370112
	ColorBlue      = 3447003
	ColorGreen     = 3066993
	ColorOrange    = 15105570
	ColorPurple    = 10181046
)

// Emoticon reaction.
const (
	ReactionArrowLeft  = "◀️"
	ReactionArrowRight = "▶️"
	ReactionArrowStart = "⏪"
	ReactionArrowEnd   = "⏩"
	ReactionInfo       = "ℹ️"
)

// ReactionPagination is reaction list for pagination.
var ReactionPagination = []string{
	ReactionArrowStart,
	ReactionArrowLeft,
	ReactionArrowRight,
	ReactionArrowEnd,
}

// ReactionPaginationWithInfo is reaction list for pagination with info.
var ReactionPaginationWithInfo = []string{
	ReactionArrowStart,
	ReactionArrowLeft,
	ReactionInfo,
	ReactionArrowRight,
	ReactionArrowEnd,
}

// Message response.
const (
	MsgInvalid       = "Invalid command. See **{{prefix}}help** for more information."
	MsgSearch3Letter = "Search at least 3 letters."
	MsgInvalidID     = "Invalid ID."
	MsgNotFound      = "Entry not found."
	MsgHelpCmd       = "{{prefix}}help"
	MsgHelpContent   = "Hayasui is a bot that help you get anime/manga/character/people information with interactive message."
	MsgSearchCmd     = "Search anime/manga/character/people"
	MsgSearchContent = "```{{prefix}}search <anime|manga|character|people> <query...>``` ```{{prefix}}search anime naruto\n{{prefix}}search manga one piece\n{{prefix}}search character ichigo\n{{prefix}}search people kana```"
	MsgAnimeCmd      = "Get anime"
	MsgAnimeContent  = "```{{prefix}}anime <anime_id>``` ```{{prefix}}anime 1```"
	MsgMangaCmd      = "Get manga"
	MsgMangaContent  = "```{{prefix}}manga <manga_id>``` ```{{prefix}}manga 1```"
	MsgCharCmd       = "Get character"
	MsgCharContent   = "```{{prefix}}character <character_id>``` ```{{prefix}}character 1```"
	MsgPeopleCmd     = "Get people"
	MsgPeopleContent = "```{{prefix}}people <people_id>``` ```{{prefix}}people 1```"
)

// DataPerPage is limit data count per page.
const DataPerPage = 10

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

// Available option of Info type.
const (
	InfoSimple int8 = iota
	InfoMore
	InfoAll
)

// SearchType
type SearchType int8

// Available option of SearchType type.
const (
	TypeAnime SearchType = iota
	TypeManga
	TypeCharacter
	TypePeople
)

// ToTitle to convert search type to title.
func (s SearchType) ToTitle() string {
	switch s {
	case TypeAnime:
		return "Anime"
	case TypeManga:
		return "Manga"
	case TypeCharacter:
		return "Character"
	case TypePeople:
		return "people"
	default:
		return ""
	}
}

// GetColor to get search type color.
func (s SearchType) GetColor() int {
	switch s {
	case TypeAnime:
		return ColorBlue
	case TypeManga:
		return ColorGreen
	case TypeCharacter:
		return ColorOrange
	case TypePeople:
		return ColorPurple
	default:
		return 0
	}
}

// GetHeader to get search type header.
func (s SearchType) GetHeader() string {
	switch s {
	case TypeAnime, TypeManga:
		return "Title"
	case TypeCharacter, TypePeople:
		return "Name"
	default:
		return ""
	}
}

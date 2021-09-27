package constant

// Type.
const (
	TypeAnime     = "anime"
	TypeManga     = "manga"
	TypeCharacter = "character"
	TypePeople    = "staff"
)

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
	MsgInvalid       = "Invalid command. See **>help** for more information."
	MsgSearch3Letter = "Search at least 3 letters."
	MsgInvalidID     = "Invalid ID."
	MsgNotFound      = "Entry not found."
	MsgHelpCmd       = ">help"
	MsgHelpContent   = "Hayasui is a bot that help you get anime/manga/character/people information with interactive message."
	MsgSearchCmd     = "Search anime/manga/character/people"
	MsgSearchContent = "```>search <anime|manga|character|people> <query...>``` ```>search anime naruto\n>search manga one piece\n>search character ichigo\n>search people kana```"
	MsgAnimeCmd      = "Get anime"
	MsgAnimeContent  = "```>anime <anime_id>``` ```>anime 1```"
	MsgMangaCmd      = "Get manga"
	MsgMangaContent  = "```>manga <manga_id>``` ```>manga 1```"
	MsgCharCmd       = "Get character"
	MsgCharContent   = "```>character <character_id>``` ```>character 1```"
	MsgPeopleCmd     = "Get people"
	MsgPeopleContent = "```>people <people_id>``` ```>people 1```"
)

// DataPerPage is limit data count per page.
const DataPerPage = 10

// MediaTypes is types of anime & manga.
var MediaTypes = map[string]string{
	"TV":       "TV",
	"TV_SHORT": "Short",
	"MOVIE":    "Movie",
	"SPECIAL":  "Special",
	"OVA":      "OVA",
	"ONA":      "ONA",
	"MUSIC":    "Music",
	"MANGA":    "Manga",
	"NOVEL":    "Novel",
	"ONE_SHOT": "OneShot",
}

// MediaStatuses is status of anime & manga.
var MediaStatuses = map[string]string{
	"FINISHED":         "Finished",
	"RELEASING":        "Releasing",
	"NOT_YET_RELEASED": "Not Yet Released",
	"CANCELLED":        "Cancelled",
	"HIATUS":           "Hiatus",
}

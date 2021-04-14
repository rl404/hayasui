package constant

// Type.
const (
	TypeAnime     = "anime"
	TypeManga     = "manga"
	TypeCharacter = "character"
	TypePeople    = "people"
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
	ReactionArrowLeft  = "⬅"
	ReactionArrowRight = "➡"
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

// AnimeTypes is types of anime.
var AnimeTypes = map[int]string{
	1: "TV",
	2: "OVA",
	3: "Movie",
	4: "Special",
	5: "ONA",
	6: "Music",
}

// MangaTypes is types of manga.
var MangaTypes = map[int]string{
	1: "Manga",
	2: "Light Novel",
	3: "One-shot",
	4: "Doujinshi",
	5: "Manhwa",
	6: "Manhua",
	7: "OEL",
	8: "Novel",
}

var MangaTypesShort = map[int]string{
	1: "Manga",
	2: "LN",
	3: "Oneshot",
	4: "Doujin",
	5: "Manhwa",
	6: "Manhua",
	7: "OEL",
	8: "Novel",
}

// AnimeStatuses is status of anime.
var AnimeStatuses = map[int]string{
	1: "Currently Airing",
	2: "Finished Airing",
	3: "Not yet aired",
}

// MangaStatuses is status of manga.
var MangaStatuses = map[int]string{
	1: "Publishing",
	2: "Finished",
	3: "Not yet published",
	4: "On Hiatus",
	5: "Discontinued",
}

package internal

// Type.
const (
	anime     = "anime"
	manga     = "manga"
	character = "character"
	people    = "people"
)

// Color.
const (
	greyLight = 12370112
	blue      = 3447003
	green     = 3066993
	orange    = 15105570
	purple    = 10181046
)

// Emoticon.
const (
	arrowStart = "⏪"
	arrowLeft  = "⬅"
	arrowRight = "➡"
	arrowEnd   = "⏩"
	info       = "ℹ️"
)

// Response.
const (
	invalidContent = "Invalid command. See **>help** for more information."
	search3Letter  = "Search at least 3 letters."
	invalidID      = "Invalid ID."
	notFound       = "Entry not found."
	helpCmd        = ">help"
	helpContent    = "Hayasui is a bot that help you get anime/manga/character/people information with interactive message."
	searchCmd      = "Search anime/manga/character/people"
	searchContent  = "```>search <anime|manga|character|people> <query...>``` ```>search anime naruto\n>search manga one piece\n>search character ichigo\n>search people kana```"
	animeCmd       = "Get anime"
	animeContent   = "```>anime <anime_id>``` ```>anime 1```"
	mangaCmd       = "Get manga"
	mangaContent   = "```>manga <manga_id>``` ```>manga 1```"
	charCmd        = "Get character"
	charContent    = "```>character <character_id>``` ```>character 1```"
	peopleCmd      = "Get people"
	peopleContent  = "```>people <people_id>``` ```>people 1```"
)

const (
	dataPerPage = 10
)

var animeTypes = map[int]string{
	1: "TV",
	2: "OVA",
	3: "Movie",
	4: "Special",
	5: "ONA",
	6: "Music",
}

var mangaTypes = map[int]string{
	1: "Manga",
	2: "Light Novel",
	3: "One-shot",
	4: "Doujinshi",
	5: "Manhwa",
	6: "Manhua",
	7: "OEL",
	8: "Novel",
}

var animeStatuses = map[int]string{
	1: "Currently Airing",
	2: "Finished Airing",
	3: "Not yet aired",
}

var mangaStatuses = map[int]string{
	1: "Publishing",
	2: "Finished",
	3: "Not yet published",
	4: "On Hiatus",
	5: "Discontinued",
}

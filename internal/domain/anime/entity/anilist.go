package entity

// AnilistTypes is types of anime & manga.
var AnilistTypes = map[string]string{
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

// AnilistStatuses is status of anime & manga.
var AnilistStatuses = map[string]string{
	"FINISHED":         "Finished",
	"RELEASING":        "Releasing",
	"NOT_YET_RELEASED": "Not Yet Released",
	"CANCELLED":        "Cancelled",
	"HIATUS":           "Hiatus",
}

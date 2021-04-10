package internal

type cacheModel struct {
	Commands []string `json:"commands"`
	Page     int      `json:"page"`
	LastPage int      `json:"lastPage"`
	Info     bool     `json:"info"`
}

type response struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Meta    metaCount `json:"meta"`
}

type metaCount struct {
	Count int `json:"count"`
}

// Search.

type searchResponse struct {
	response
	Data []searchData `json:"data"`
}

type searchData struct {
	ID    int    `json:"id"`
	Title string `json:"title"` // Anime & manga.
	Name  string `json:"name"`  // Character & people.
}

// Anime & manga.

type animeMangaResponse struct {
	response
	Data animeMangaData `json:"data"`
}

type animeMangaData struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	AltTitles  altTitle `json:"alternativeTitles"`
	Image      string   `json:"image"`
	Synopsis   string   `json:"synopsis"`
	Score      float64  `json:"score"`
	Voter      int      `json:"voter"`
	Rank       int      `json:"rank"`
	Popularity int      `json:"popularity"`
	Member     int      `json:"member"`
	Favorite   int      `json:"favorite"`
	Type       int      `json:"type"`
	Episode    int      `json:"episode"`
	Volume     int      `json:"volume"`
	Chapter    int      `json:"chapter"`
	Status     int      `json:"status"`
	Airing     airing   `json:"airing"`
	Publishing airing   `json:"publishing"`
	Duration   string   `json:"duration"`
	Premiered  string   `json:"premiered"`
	Source     int      `json:"source"`
	Rating     int      `json:"rating"`
}

type altTitle struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Synonym  string `json:"synonym"`
}

type airing struct {
	Start date   `json:"start"`
	End   date   `json:"end"`
	Day   string `json:"day"`
	Time  string `json:"time"`
}

type date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// Character & people.

type charPeopleResponse struct {
	response
	Data charPeopleData `json:"data"`
}

type charPeopleData struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	Nicknames        []string `json:"nicknames"`
	JapaneseName     string   `json:"japaneseName"`
	Favorite         int      `json:"favorite"`
	About            string   `json:"about"`
	GivenName        string   `json:"givenName"`
	FamilyName       string   `json:"familyName"`
	AlternativeNames []string `json:"alternativeNames"`
	Birthday         date     `json:"birthday"`
	Website          string   `json:"website"`
	More             string   `json:"more"`
}

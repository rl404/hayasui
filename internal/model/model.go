package model

// Command is cached command.
type Command struct {
	Commands []string `json:"commands"`
	Page     int      `json:"page"`
	LastPage int      `json:"lastPage"`
	Info     bool     `json:"info"`
}

// Response is base api response model.
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Meta    Meta   `json:"meta"`
}

// Meta is api response meta field model.
type Meta struct {
	Count int `json:"count"`
}

// ResponseError is error api response model.
type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// Date is common date format model.
type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// ResponseSearch is api search response model.
type ResponseSearch struct {
	Response
	Data []DataSearch `json:"data"`
}

// DataSearch is api search response data model.
type DataSearch struct {
	ID    int    `json:"id"`
	Title string `json:"title"` // Anime & manga.
	Name  string `json:"name"`  // Character & people.
}

// ResponseAnimeManga is api anime & manga response model.
type ResponseAnimeManga struct {
	Response
	Data DataAnimeManga `json:"data"`
}

// DataAnimeManga is api anime & manga response data model.
type DataAnimeManga struct {
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
	Status     int      `json:"status"`

	// Anime only.
	Episode   int    `json:"episode"`
	Airing    airing `json:"airing"`
	Duration  string `json:"duration"`
	Premiered string `json:"premiered"`
	Source    int    `json:"source"`
	Rating    int    `json:"rating"`

	// Manga only.
	Volume     int    `json:"volume"`
	Chapter    int    `json:"chapter"`
	Publishing airing `json:"publishing"`
}

type altTitle struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Synonym  string `json:"synonym"`
}

type airing struct {
	Start Date   `json:"start"`
	End   Date   `json:"end"`
	Day   string `json:"day"`
	Time  string `json:"time"`
}

// ResponseCharPeople is api character & people response model.
type ResponseCharPeople struct {
	Response
	Data DataCharPeople `json:"data"`
}

// DataCharPeople is api character & people response data model.
type DataCharPeople struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Favorite int    `json:"favorite"`

	// Character only.
	Nicknames    []string `json:"nicknames"`
	JapaneseName string   `json:"japaneseName"`
	About        string   `json:"about"`

	// People only.
	GivenName        string   `json:"givenName"`
	FamilyName       string   `json:"familyName"`
	AlternativeNames []string `json:"alternativeNames"`
	Birthday         Date     `json:"birthday"`
	Website          string   `json:"website"`
	More             string   `json:"more"`
}

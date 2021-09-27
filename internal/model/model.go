package model

// Command is cached command.
type Command struct {
	Commands []string `json:"commands"`
	Page     int      `json:"page"`
	LastPage int      `json:"lastPage"`
	Info     bool     `json:"info"`
	Type     int      `json:"type"`
}

// Date is common date format model.
type Date struct {
	Year  int
	Month int
	Day   int
}

// DataSearchAnimeManga is api anime & manga search response data model.
type DataSearchAnimeManga struct {
	ID    int
	Title string
	Type  string
	Score int
}

// DataSearchCharPeople is api character & people search response data model.
type DataSearchCharPeople struct {
	ID   int
	Name string
}

// DataAnimeManga is api anime & manga response data model.
type DataAnimeManga struct {
	ID       int
	Title    Title
	Image    string
	Synopsis string
	Score    int
	Rankings []string
	Member   int
	Favorite int
	Type     string
	Status   string

	// Anime only.
	Episode int
	Airing  Airing

	// Manga only.
	Volume     int
	Chapter    int
	Publishing Airing
}

// Title is anime & manga titles.
type Title struct {
	English string
	Romaji  string
	Native  string
}

// Airing is anime & manga airing/publishing date.
type Airing struct {
	Start Date
	End   Date
}

// DataCharPeople is api character & people response data model.
type DataCharPeople struct {
	ID       int
	Name     string
	Image    string
	Favorite int

	// Character only.
	Nicknames    []string
	JapaneseName string
	About        string

	// People only.
	AlternativeNames []string
	Birthday         Date
	More             string
}

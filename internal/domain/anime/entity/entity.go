package entity

// Anime is entity for anime.
type Anime struct {
	ID            int
	URL           string
	Title         string
	TitleEnglish  string
	TitleJapanese string
	TitleSynonyms []string
	Synopsis      string
	Image         string
	Score         float64
	Member        int
	Favorite      int
	Type          string
	Status        string
	Episode       int
	Ranking       string
	StartDate     Date
	EndDate       Date
}

// Manga is entity for manga.
type Manga struct {
	ID            int
	URL           string
	Title         string
	TitleEnglish  string
	TitleJapanese string
	TitleSynonyms []string
	Synopsis      string
	Image         string
	Score         float64
	Member        int
	Favorite      int
	Type          string
	Status        string
	Chapter       int
	Ranking       string
	StartDate     Date
	EndDate       Date
}

// Character is entity for character.
type Character struct {
	ID           int
	URL          string
	Name         string
	NameJapanese string
	Nicknames    []string
	Favorite     int
	About        string
	Image        string
}

// People is entity for people.
type People struct {
	ID               int
	URL              string
	Name             string
	AlternativeNames []string
	Birthday         Date
	Favorite         int
	About            string
	Image            string
}

// Date is entity for date.
type Date struct {
	Year  int
	Month int
	Day   int
}

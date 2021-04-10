package internal

import (
	"fmt"
	"net/http"
	"net/url"
)

// API contains all api functions.
type API interface {
	SearchAnime(query string, page int) ([]searchData, int, error)
	SearchManga(query string, page int) ([]searchData, int, error)
	SearchCharacter(query string, page int) ([]searchData, int, error)
	SearchPeople(query string, page int) ([]searchData, int, error)
	GetAnime(id int) (*animeMangaData, int, error)
	GetManga(id int) (*animeMangaData, int, error)
	GetCharacter(id int) (*charPeopleData, int, error)
	GetPeople(id int) (*charPeopleData, int, error)
}

type api struct {
	http Requester
}

// NewAPI to create new api.
func NewAPI(host string) (API, error) {
	return &api{
		http: newRequester(host),
	}, nil
}

// SearchAnime to search anime.
func (a *api) SearchAnime(query string, page int) (m []searchData, cnt int, err error) {
	var res searchResponse
	url := fmt.Sprintf("/search/anime?title=%s&page=%v&limit=%v", url.QueryEscape(query), page, dataPerPage)
	if _, err := a.http.Get(url, &res); err != nil {
		return nil, 0, err
	}
	if res.Status != http.StatusOK {
		return []searchData{}, 0, errNot200
	}
	return res.Data, res.Meta.Count, nil
}

// SearchManga to search manga.
func (a *api) SearchManga(query string, page int) ([]searchData, int, error) {
	var res searchResponse
	url := fmt.Sprintf("/search/manga?title=%s&page=%v&limit=%v", url.QueryEscape(query), page, dataPerPage)
	if _, err := a.http.Get(url, &res); err != nil {
		return nil, 0, err
	}
	if res.Status != http.StatusOK {
		return []searchData{}, 0, errNot200
	}
	return res.Data, res.Meta.Count, nil
}

// SearchCharacter to search character.
func (a *api) SearchCharacter(query string, page int) ([]searchData, int, error) {
	var res searchResponse
	url := fmt.Sprintf("/search/character?name=%s&page=%v&limit=%v", url.QueryEscape(query), page, dataPerPage)
	if _, err := a.http.Get(url, &res); err != nil {
		return nil, 0, err
	}
	if res.Status != http.StatusOK {
		return []searchData{}, 0, errNot200
	}
	return res.Data, res.Meta.Count, nil
}

// SearchPeople to search people.
func (a *api) SearchPeople(query string, page int) ([]searchData, int, error) {
	var res searchResponse
	url := fmt.Sprintf("/search/people?name=%s&page=%v&limit=%v", url.QueryEscape(query), page, dataPerPage)
	if _, err := a.http.Get(url, &res); err != nil {
		return nil, 0, err
	}
	if res.Status != http.StatusOK {
		return []searchData{}, 0, errNot200
	}
	return res.Data, res.Meta.Count, nil
}

// GetAnime to get anime.
func (a *api) GetAnime(id int) (*animeMangaData, int, error) {
	var res animeMangaResponse
	url := fmt.Sprintf("/anime/%v", id)
	if code, err := a.http.Get(url, &res); err != nil || code != http.StatusOK {
		return nil, code, err
	}
	if res.Status != http.StatusOK {
		return nil, res.Status, errNot200
	}
	return &res.Data, http.StatusOK, nil
}

// GetManga to get manga.
func (a *api) GetManga(id int) (*animeMangaData, int, error) {
	var res animeMangaResponse
	url := fmt.Sprintf("/manga/%v", id)
	if code, err := a.http.Get(url, &res); err != nil || code != http.StatusOK {
		return nil, code, err
	}
	if res.Status != http.StatusOK {
		return nil, res.Status, errNot200
	}
	return &res.Data, http.StatusOK, nil
}

// GetCharacter to get character.
func (a *api) GetCharacter(id int) (*charPeopleData, int, error) {
	var res charPeopleResponse
	url := fmt.Sprintf("/character/%v", id)
	if code, err := a.http.Get(url, &res); err != nil || code != http.StatusOK {
		return nil, code, err
	}
	if res.Status != http.StatusOK {
		return nil, res.Status, errNot200
	}
	return &res.Data, http.StatusOK, nil
}

// GetPeople to get people.
func (a *api) GetPeople(id int) (*charPeopleData, int, error) {
	var res charPeopleResponse
	url := fmt.Sprintf("/people/%v", id)
	if code, err := a.http.Get(url, &res); err != nil || code != http.StatusOK {
		return nil, code, err
	}
	if res.Status != http.StatusOK {
		return nil, res.Status, errNot200
	}
	return &res.Data, http.StatusOK, nil
}

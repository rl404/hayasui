package http

import (
	"fmt"
	"net/url"

	"github.com/rl404/hayasui/internal/api"
	"github.com/rl404/hayasui/internal/constant"
	"github.com/rl404/hayasui/internal/model"
)

type http struct {
	http requester
}

// New to create new http api.
func New(host string) api.API {
	return &http{
		http: newRequester(host),
	}
}

// SearchAnime to search anime.
func (h *http) SearchAnime(query string, page int) (m []model.DataSearch, cnt int, err error) {
	var res model.ResponseSearch
	url := fmt.Sprintf("/search/anime?title=%s&page=%v&limit=%v", url.QueryEscape(query), page, constant.DataPerPage)
	if _, err := h.http.Get(url, &res); err != nil {
		return nil, 0, err
	}
	return res.Data, res.Meta.Count, nil
}

// SearchManga to search manga.
func (h *http) SearchManga(query string, page int) ([]model.DataSearch, int, error) {
	var res model.ResponseSearch
	url := fmt.Sprintf("/search/manga?title=%s&page=%v&limit=%v", url.QueryEscape(query), page, constant.DataPerPage)
	if _, err := h.http.Get(url, &res); err != nil {
		return nil, 0, err
	}
	return res.Data, res.Meta.Count, nil
}

// SearchCharacter to search character.
func (h *http) SearchCharacter(query string, page int) ([]model.DataSearch, int, error) {
	var res model.ResponseSearch
	url := fmt.Sprintf("/search/character?name=%s&page=%v&limit=%v", url.QueryEscape(query), page, constant.DataPerPage)
	if _, err := h.http.Get(url, &res); err != nil {
		return nil, 0, err
	}
	return res.Data, res.Meta.Count, nil
}

// SearchPeople to search people.
func (h *http) SearchPeople(query string, page int) ([]model.DataSearch, int, error) {
	var res model.ResponseSearch
	url := fmt.Sprintf("/search/people?name=%s&page=%v&limit=%v", url.QueryEscape(query), page, constant.DataPerPage)
	if _, err := h.http.Get(url, &res); err != nil {
		return nil, 0, err
	}
	return res.Data, res.Meta.Count, nil
}

// GetAnime to get anime.
func (h *http) GetAnime(id int) (*model.DataAnimeManga, error) {
	var res model.ResponseAnimeManga
	url := fmt.Sprintf("/anime/%v", id)
	if _, err := h.http.Get(url, &res); err != nil {
		return nil, err
	}
	return &res.Data, nil
}

// GetManga to get manga.
func (h *http) GetManga(id int) (*model.DataAnimeManga, error) {
	var res model.ResponseAnimeManga
	url := fmt.Sprintf("/manga/%v", id)
	if _, err := h.http.Get(url, &res); err != nil {
		return nil, err
	}
	return &res.Data, nil
}

// GetCharacter to get character.
func (h *http) GetCharacter(id int) (*model.DataCharPeople, error) {
	var res model.ResponseCharPeople
	url := fmt.Sprintf("/character/%v", id)
	if _, err := h.http.Get(url, &res); err != nil {
		return nil, err
	}
	return &res.Data, nil
}

// GetPeople to get people.
func (h *http) GetPeople(id int) (*model.DataCharPeople, error) {
	var res model.ResponseCharPeople
	url := fmt.Sprintf("/people/%v", id)
	if _, err := h.http.Get(url, &res); err != nil {
		return nil, err
	}
	return &res.Data, nil
}

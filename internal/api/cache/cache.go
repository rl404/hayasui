package cache

import (
	"github.com/rl404/hayasui/internal/api"
	_cache "github.com/rl404/hayasui/internal/cache"
	"github.com/rl404/hayasui/internal/constant"
	"github.com/rl404/hayasui/internal/model"
)

type cache struct {
	cacher _cache.Cacher
	api    api.API
}

// New to create new cacher.
func New(c _cache.Cacher, a api.API) api.API {
	return &cache{
		cacher: c,
		api:    a,
	}
}

// SearchAnime to search anime.
func (c *cache) SearchAnime(query string, page int) ([]model.DataSearchAnimeManga, int, error) {
	// Cache model.
	var data struct {
		Data  []model.DataSearchAnimeManga
		Count int
	}

	// Get from cache.
	if c.cacher.Get(&data, "search", constant.TypeAnime, query, page) == nil {
		return data.Data, data.Count, nil
	}

	// Call api.
	d, cnt, err := c.api.SearchAnime(query, page)
	if err != nil {
		return d, cnt, err
	}

	// Save to cache.
	data.Data, data.Count = d, cnt
	go c.cacher.Set(data, "search", constant.TypeAnime, query, page)

	return d, cnt, nil
}

// SearchManga to search manga.
func (c *cache) SearchManga(query string, page int) ([]model.DataSearchAnimeManga, int, error) {
	// Cache model.
	var data struct {
		Data  []model.DataSearchAnimeManga
		Count int
	}

	// Get from cache.
	if c.cacher.Get(&data, "search", constant.TypeManga, query, page) == nil {
		return data.Data, data.Count, nil
	}

	// Call api.
	d, cnt, err := c.api.SearchManga(query, page)
	if err != nil {
		return d, cnt, err
	}

	// Save to cache.
	data.Data, data.Count = d, cnt
	go c.cacher.Set(data, "search", constant.TypeManga, query, page)

	return d, cnt, nil
}

// SearchCharacter to search character.
func (c *cache) SearchCharacter(query string, page int) ([]model.DataSearchCharPeople, int, error) {
	// Cache model.
	var data struct {
		Data  []model.DataSearchCharPeople
		Count int
	}

	// Get from cache.
	if c.cacher.Get(&data, "search", constant.TypeCharacter, query, page) == nil {
		return data.Data, data.Count, nil
	}

	// Call api.
	d, cnt, err := c.api.SearchCharacter(query, page)
	if err != nil {
		return d, cnt, err
	}

	// Save to cache.
	data.Data, data.Count = d, cnt
	go c.cacher.Set(data, "search", constant.TypeCharacter, query, page)

	return d, cnt, nil
}

// SearchPeople to search people.
func (c *cache) SearchPeople(query string, page int) ([]model.DataSearchCharPeople, int, error) {
	// Cache model.
	var data struct {
		Data  []model.DataSearchCharPeople
		Count int
	}

	// Get from cache.
	if c.cacher.Get(&data, "search", constant.TypePeople, query, page) == nil {
		return data.Data, data.Count, nil
	}

	// Call api.
	d, cnt, err := c.api.SearchPeople(query, page)
	if err != nil {
		return d, cnt, err
	}

	// Save to cache.
	data.Data, data.Count = d, cnt
	go c.cacher.Set(data, "search", constant.TypePeople, query, page)

	return d, cnt, nil
}

// GetAnime to get anime.
func (c *cache) GetAnime(id int) (data *model.DataAnimeManga, err error) {
	// Get from cache.
	if c.cacher.Get(&data, constant.TypeAnime, id) == nil {
		return data, nil
	}

	// Call api.
	data, err = c.api.GetAnime(id)
	if err != nil {
		return nil, err
	}

	// Save to cache.
	go c.cacher.Set(data, constant.TypeAnime, id)

	return data, nil
}

// GetManga to get manga.
func (c *cache) GetManga(id int) (data *model.DataAnimeManga, err error) {
	// Get from cache.
	if c.cacher.Get(&data, constant.TypeManga, id) == nil {
		return data, nil
	}

	// Call api.
	data, err = c.api.GetManga(id)
	if err != nil {
		return nil, err
	}

	// Save to cache.
	go c.cacher.Set(data, constant.TypeManga, id)

	return data, nil
}

// GetCharacter to get anime.
func (c *cache) GetCharacter(id int) (data *model.DataCharPeople, err error) {
	// Get from cache.
	if c.cacher.Get(&data, constant.TypeCharacter, id) == nil {
		return data, nil
	}

	// Call api.
	data, err = c.api.GetCharacter(id)
	if err != nil {
		return nil, err
	}

	// Save to cache.
	go c.cacher.Set(data, constant.TypeCharacter, id)

	return data, nil
}

// GetPeople to get anime.
func (c *cache) GetPeople(id int) (data *model.DataCharPeople, err error) {
	// Get from cache.
	if c.cacher.Get(&data, constant.TypePeople, id) == nil {
		return data, nil
	}

	// Call api.
	data, err = c.api.GetPeople(id)
	if err != nil {
		return nil, err
	}

	// Save to cache.
	go c.cacher.Set(data, constant.TypePeople, id)

	return data, nil
}

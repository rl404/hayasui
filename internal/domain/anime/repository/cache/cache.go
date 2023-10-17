package cache

import (
	"context"

	"github.com/rl404/fairy/cache"
	"github.com/rl404/fairy/errors/stack"
	"github.com/rl404/hayasui/internal/domain/anime/entity"
	"github.com/rl404/hayasui/internal/domain/anime/repository"
	"github.com/rl404/hayasui/internal/utils"
)

// Cache contains functions for anime cache.
type Cache struct {
	cacher cache.Cacher
	repo   repository.Repository
}

// New to create new anime cache.
func New(cacher cache.Cacher, repo repository.Repository) *Cache {
	return &Cache{
		cacher: cacher,
		repo:   repo,
	}
}

// GetAnime to get anime from cache.
func (c *Cache) GetAnime(ctx context.Context, id int) (data *entity.Anime, _ error) {
	key := utils.GetKey("anime", id)
	if err := c.cacher.Get(ctx, key, &data); err == nil {
		return data, nil
	}

	data, err := c.repo.GetAnime(ctx, id)
	if err != nil {
		return nil, stack.Wrap(ctx, err)
	}

	if err := c.cacher.Set(ctx, key, data); err != nil {
		return nil, stack.Wrap(ctx, err)
	}

	return data, nil
}

// GetManga to get manga from cache.
func (c *Cache) GetManga(ctx context.Context, id int) (data *entity.Manga, _ error) {
	key := utils.GetKey("manga", id)
	if err := c.cacher.Get(ctx, key, &data); err == nil {
		return data, nil
	}

	data, err := c.repo.GetManga(ctx, id)
	if err != nil {
		return nil, stack.Wrap(ctx, err)
	}

	if err := c.cacher.Set(ctx, key, data); err != nil {
		return nil, stack.Wrap(ctx, err)
	}

	return data, nil
}

// GetCharacter to get character from cache.
func (c *Cache) GetCharacter(ctx context.Context, id int) (data *entity.Character, _ error) {
	key := utils.GetKey("character", id)
	if err := c.cacher.Get(ctx, key, &data); err == nil {
		return data, nil
	}

	data, err := c.repo.GetCharacter(ctx, id)
	if err != nil {
		return nil, stack.Wrap(ctx, err)
	}

	if err := c.cacher.Set(ctx, key, data); err != nil {
		return nil, stack.Wrap(ctx, err)
	}

	return data, nil
}

// GetPeople to get people from cache.
func (c *Cache) GetPeople(ctx context.Context, id int) (data *entity.People, _ error) {
	key := utils.GetKey("people", id)
	if err := c.cacher.Get(ctx, key, &data); err == nil {
		return data, nil
	}

	data, err := c.repo.GetPeople(ctx, id)
	if err != nil {
		return nil, stack.Wrap(ctx, err)
	}

	if err := c.cacher.Set(ctx, key, data); err != nil {
		return nil, stack.Wrap(ctx, err)
	}

	return data, nil
}

// SearchAnime to search anime.
func (c *Cache) SearchAnime(ctx context.Context, query string, page int) ([]entity.Anime, int, error) {
	var data struct {
		Data  []entity.Anime
		Count int
	}

	key := utils.GetKey("search", "anime", page, query)
	if c.cacher.Get(ctx, key, &data) == nil {
		return data.Data, data.Count, nil
	}

	d, cnt, err := c.repo.SearchAnime(ctx, query, page)
	if err != nil {
		return d, cnt, err
	}

	data.Data = d
	data.Count = cnt

	if err := c.cacher.Set(ctx, key, data); err != nil {
		return nil, 0, stack.Wrap(ctx, err)
	}

	return d, cnt, nil
}

// SearchManga to search manga.
func (c *Cache) SearchManga(ctx context.Context, query string, page int) ([]entity.Manga, int, error) {
	var data struct {
		Data  []entity.Manga
		Count int
	}

	key := utils.GetKey("search", "manga", page, query)
	if c.cacher.Get(ctx, key, &data) == nil {
		return data.Data, data.Count, nil
	}

	d, cnt, err := c.repo.SearchManga(ctx, query, page)
	if err != nil {
		return d, cnt, err
	}

	data.Data = d
	data.Count = cnt

	if err := c.cacher.Set(ctx, key, data); err != nil {
		return nil, 0, stack.Wrap(ctx, err)
	}

	return d, cnt, nil
}

// SearchCharacter to search character.
func (c *Cache) SearchCharacter(ctx context.Context, query string, page int) ([]entity.Character, int, error) {
	var data struct {
		Data  []entity.Character
		Count int
	}

	key := utils.GetKey("search", "character", page, query)
	if c.cacher.Get(ctx, key, &data) == nil {
		return data.Data, data.Count, nil
	}

	d, cnt, err := c.repo.SearchCharacter(ctx, query, page)
	if err != nil {
		return d, cnt, err
	}

	data.Data = d
	data.Count = cnt

	if err := c.cacher.Set(ctx, key, data); err != nil {
		return nil, 0, stack.Wrap(ctx, err)
	}

	return d, cnt, nil
}

// SearchPeople to search people.
func (c *Cache) SearchPeople(ctx context.Context, query string, page int) ([]entity.People, int, error) {
	var data struct {
		Data  []entity.People
		Count int
	}

	key := utils.GetKey("search", "people", page, query)
	if c.cacher.Get(ctx, key, &data) == nil {
		return data.Data, data.Count, nil
	}

	d, cnt, err := c.repo.SearchPeople(ctx, query, page)
	if err != nil {
		return d, cnt, err
	}

	data.Data = d
	data.Count = cnt

	if err := c.cacher.Set(ctx, key, data); err != nil {
		return nil, 0, stack.Wrap(ctx, err)
	}

	return d, cnt, nil
}

package repository

import (
	"context"

	"github.com/rl404/hayasui/internal/domain/anime/entity"
)

// Repository contains functions for anime domain.
type Repository interface {
	GetAnime(ctx context.Context, id int) (*entity.Anime, error)
	GetManga(ctx context.Context, id int) (*entity.Manga, error)
	GetCharacter(ctx context.Context, id int) (*entity.Character, error)
	GetPeople(ctx context.Context, id int) (*entity.People, error)
	SearchAnime(ctx context.Context, query string, page int) ([]entity.Anime, int, error)
	SearchManga(ctx context.Context, query string, page int) ([]entity.Manga, int, error)
	SearchCharacter(ctx context.Context, query string, page int) ([]entity.Character, int, error)
	SearchPeople(ctx context.Context, query string, page int) ([]entity.People, int, error)
}

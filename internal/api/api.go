package api

import "github.com/rl404/hayasui/internal/model"

// API contains all api functions.
type API interface {
	SearchAnime(query string, page int) ([]model.DataSearchAnimeManga, int, error)
	SearchManga(query string, page int) ([]model.DataSearchAnimeManga, int, error)
	SearchCharacter(query string, page int) ([]model.DataSearchCharPeople, int, error)
	SearchPeople(query string, page int) ([]model.DataSearchCharPeople, int, error)
	GetAnime(id int) (*model.DataAnimeManga, error)
	GetManga(id int) (*model.DataAnimeManga, error)
	GetCharacter(id int) (*model.DataCharPeople, error)
	GetPeople(id int) (*model.DataCharPeople, error)
}

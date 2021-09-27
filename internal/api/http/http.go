package http

import (
	"github.com/rl404/hayasui/internal/api"
	"github.com/rl404/hayasui/internal/constant"
	"github.com/rl404/hayasui/internal/model"
	"github.com/rl404/hayasui/internal/utils"
	"github.com/rl404/verniy"
)

type http struct {
	client *verniy.Client
}

// New to create new http api.
func New() api.API {
	return &http{
		client: verniy.New(),
	}
}

// SearchAnime to search anime.
func (h *http) SearchAnime(query string, page int) (m []model.DataSearchAnimeManga, cnt int, err error) {
	data, err := h.client.SearchAnime(verniy.PageParamMedia{
		Search: query,
	}, page, constant.DataPerPage,
		verniy.MediaFieldID,
		verniy.MediaFieldTitle(
			verniy.MediaTitleFieldRomaji,
		),
		verniy.MediaFieldFormat,
		verniy.MediaFieldAverageScore,
	)
	if err != nil {
		return nil, 0, err
	}

	res := make([]model.DataSearchAnimeManga, len(data.Media))
	for i, d := range data.Media {
		res[i] = model.DataSearchAnimeManga{
			ID:    d.ID,
			Title: utils.PtrToStr(d.Title.Romaji),
			Type:  h.formatToStr(d.Format),
			Score: utils.PtrToInt(d.AverageScore),
		}
	}

	return res, *data.PageInfo.Total, nil
}

// SearchManga to search manga.
func (h *http) SearchManga(query string, page int) ([]model.DataSearchAnimeManga, int, error) {
	data, err := h.client.SearchManga(verniy.PageParamMedia{
		Search: query,
	}, page, constant.DataPerPage,
		verniy.MediaFieldID,
		verniy.MediaFieldTitle(
			verniy.MediaTitleFieldRomaji,
		),
		verniy.MediaFieldFormat,
		verniy.MediaFieldAverageScore,
	)
	if err != nil {
		return nil, 0, err
	}

	res := make([]model.DataSearchAnimeManga, len(data.Media))
	for i, d := range data.Media {
		res[i] = model.DataSearchAnimeManga{
			ID:    d.ID,
			Title: utils.PtrToStr(d.Title.Romaji),
			Type:  h.formatToStr(d.Format),
			Score: utils.PtrToInt(d.AverageScore),
		}
	}

	return res, *data.PageInfo.Total, nil
}

// SearchCharacter to search character.
func (h *http) SearchCharacter(query string, page int) ([]model.DataSearchCharPeople, int, error) {
	data, err := h.client.SearchCharacter(verniy.PageParamCharacters{
		Search: query,
	}, page, constant.DataPerPage,
		verniy.CharacterFieldID,
		verniy.CharacterFieldName(
			verniy.CharacterNameFieldFull,
		),
	)
	if err != nil {
		return nil, 0, err
	}

	res := make([]model.DataSearchCharPeople, len(data.Characters))
	for i, d := range data.Characters {
		res[i] = model.DataSearchCharPeople{
			ID:   d.ID,
			Name: utils.PtrToStr(d.Name.Full),
		}
	}

	return res, *data.PageInfo.Total, nil
}

// SearchPeople to search people.
func (h *http) SearchPeople(query string, page int) ([]model.DataSearchCharPeople, int, error) {
	data, err := h.client.SearchStaff(verniy.PageParamStaff{
		Search: query,
	}, page, constant.DataPerPage,
		verniy.StaffFieldID,
		verniy.StaffFieldName(
			verniy.StaffNameFieldFull,
		),
	)
	if err != nil {
		return nil, 0, err
	}

	res := make([]model.DataSearchCharPeople, len(data.Staff))
	for i, d := range data.Staff {
		res[i] = model.DataSearchCharPeople{
			ID:   d.ID,
			Name: utils.PtrToStr(d.Name.Full),
		}
	}

	return res, *data.PageInfo.Total, nil
}

// GetAnime to get anime.
func (h *http) GetAnime(id int) (*model.DataAnimeManga, error) {
	data, err := h.client.GetAnime(id,
		verniy.MediaFieldID,
		verniy.MediaFieldTitle(
			verniy.MediaTitleFieldEnglish,
			verniy.MediaTitleFieldNative,
			verniy.MediaTitleFieldRomaji,
		),
		verniy.MediaFieldCoverImage(
			verniy.MediaCoverImageFieldLarge,
		),
		verniy.MediaFieldDescription,
		verniy.MediaFieldAverageScore,
		verniy.MediaFieldPopularity,
		verniy.MediaFieldFavourites,
		verniy.MediaFieldFormat,
		verniy.MediaFieldStatusV2,
		verniy.MediaFieldEpisodes,
		verniy.MediaFieldStartDate,
		verniy.MediaFieldEndDate,
		verniy.MediaFieldRankings(
			verniy.MediaRankFieldRank,
			verniy.MediaRankFieldAllTime,
			verniy.MediaRankFieldContext,
		),
	)
	if err != nil {
		return nil, err
	}

	return &model.DataAnimeManga{
		ID: data.ID,
		Title: model.Title{
			English: utils.PtrToStr(data.Title.English),
			Romaji:  utils.PtrToStr(data.Title.Romaji),
			Native:  utils.PtrToStr(data.Title.Native),
		},
		Image:    utils.PtrToStr(data.CoverImage.Large),
		Synopsis: utils.PtrToStr(data.Description),
		Score:    utils.PtrToInt(data.AverageScore),
		Rankings: h.getRanking(data.Rankings),
		Member:   utils.PtrToInt(data.Popularity),
		Favorite: utils.PtrToInt(data.Favourites),
		Type:     h.formatToStr(data.Format),
		Status:   h.statusToStr(data.Status),
		Episode:  utils.PtrToInt(data.Episodes),
		Airing: model.Airing{
			Start: model.Date{
				Year:  utils.PtrToInt(data.StartDate.Year),
				Month: utils.PtrToInt(data.StartDate.Month),
				Day:   utils.PtrToInt(data.StartDate.Day),
			},
			End: model.Date{
				Year:  utils.PtrToInt(data.EndDate.Year),
				Month: utils.PtrToInt(data.EndDate.Month),
				Day:   utils.PtrToInt(data.EndDate.Day),
			},
		},
	}, nil
}

// GetManga to get manga.
func (h *http) GetManga(id int) (*model.DataAnimeManga, error) {
	data, err := h.client.GetManga(id,
		verniy.MediaFieldID,
		verniy.MediaFieldTitle(
			verniy.MediaTitleFieldEnglish,
			verniy.MediaTitleFieldNative,
			verniy.MediaTitleFieldRomaji,
		),
		verniy.MediaFieldCoverImage(
			verniy.MediaCoverImageFieldLarge,
		),
		verniy.MediaFieldDescription,
		verniy.MediaFieldAverageScore,
		verniy.MediaFieldPopularity,
		verniy.MediaFieldFavourites,
		verniy.MediaFieldFormat,
		verniy.MediaFieldStatusV2,
		verniy.MediaFieldVolumes,
		verniy.MediaFieldChapters,
		verniy.MediaFieldStartDate,
		verniy.MediaFieldEndDate,
		verniy.MediaFieldRankings(
			verniy.MediaRankFieldRank,
			verniy.MediaRankFieldAllTime,
			verniy.MediaRankFieldContext,
		),
	)
	if err != nil {
		return nil, err
	}

	return &model.DataAnimeManga{
		ID: data.ID,
		Title: model.Title{
			English: utils.PtrToStr(data.Title.English),
			Romaji:  utils.PtrToStr(data.Title.Romaji),
			Native:  utils.PtrToStr(data.Title.Native),
		},
		Image:    utils.PtrToStr(data.CoverImage.Large),
		Synopsis: utils.PtrToStr(data.Description),
		Score:    utils.PtrToInt(data.AverageScore),
		Rankings: h.getRanking(data.Rankings),
		Member:   utils.PtrToInt(data.Popularity),
		Favorite: utils.PtrToInt(data.Favourites),
		Type:     h.formatToStr(data.Format),
		Status:   h.statusToStr(data.Status),
		Volume:   utils.PtrToInt(data.Volumes),
		Chapter:  utils.PtrToInt(data.Chapters),
		Publishing: model.Airing{
			Start: model.Date{
				Year:  utils.PtrToInt(data.StartDate.Year),
				Month: utils.PtrToInt(data.StartDate.Month),
				Day:   utils.PtrToInt(data.StartDate.Day),
			},
			End: model.Date{
				Year:  utils.PtrToInt(data.EndDate.Year),
				Month: utils.PtrToInt(data.EndDate.Month),
				Day:   utils.PtrToInt(data.EndDate.Day),
			},
		},
	}, nil
}

// GetCharacter to get character.
func (h *http) GetCharacter(id int) (*model.DataCharPeople, error) {
	data, err := h.client.GetCharacter(id,
		verniy.CharacterFieldID,
		verniy.CharacterFieldName(
			verniy.CharacterNameFieldFull,
			verniy.CharacterNameFieldAlternative,
			verniy.CharacterNameFieldNative,
		),
		verniy.CharacterFieldImage(
			verniy.CharacterImageFieldLarge,
		),
		verniy.CharacterFieldFavourite,
		verniy.CharacterFieldDescription,
	)
	if err != nil {
		return nil, err
	}

	return &model.DataCharPeople{
		ID:           data.ID,
		Name:         utils.PtrToStr(data.Name.Full),
		Image:        utils.PtrToStr(data.Image.Large),
		Favorite:     utils.PtrToInt(data.Favourites),
		Nicknames:    data.Name.Alternative,
		JapaneseName: utils.PtrToStr(data.Name.Native),
		About:        utils.PtrToStr(data.Description),
	}, nil
}

// GetPeople to get people.
func (h *http) GetPeople(id int) (*model.DataCharPeople, error) {
	data, err := h.client.GetStaff(id,
		verniy.StaffFieldID,
		verniy.StaffFieldName(
			verniy.StaffNameFieldFull,
			verniy.StaffNameFieldAlternative,
		),
		verniy.StaffFieldImage(
			verniy.StaffImageFieldLarge,
		),
		verniy.StaffFieldFavourites,
		verniy.StaffFieldDateOfBirth,
		verniy.StaffFieldDescription,
	)
	if err != nil {
		return nil, err
	}

	return &model.DataCharPeople{
		ID:               data.ID,
		Name:             utils.PtrToStr(data.Name.Full),
		Image:            utils.PtrToStr(data.Image.Large),
		Favorite:         utils.PtrToInt(data.Favourites),
		AlternativeNames: data.Name.Alternative,
		Birthday: model.Date{
			Year:  utils.PtrToInt(data.DateOfBirth.Year),
			Month: utils.PtrToInt(data.DateOfBirth.Month),
			Day:   utils.PtrToInt(data.DateOfBirth.Day),
		},
		More: utils.PtrToStr(data.Description),
	}, nil
}

package anilist

import (
	"context"

	"github.com/rl404/fairy/errors/stack"
	"github.com/rl404/hayasui/internal/domain/anime/entity"
	"github.com/rl404/hayasui/internal/utils"
	"github.com/rl404/verniy"
)

// GetAnime to get anime.
func (c *Client) GetAnime(ctx context.Context, id int) (*entity.Anime, error) {
	data, err := c.client.GetAnimeWithContext(ctx, id,
		verniy.MediaFieldID,
		verniy.MediaFieldTitle(
			verniy.MediaTitleFieldEnglish,
			verniy.MediaTitleFieldNative,
			verniy.MediaTitleFieldRomaji,
		),
		verniy.MediaFieldCoverImage(
			verniy.MediaCoverImageFieldLarge,
		),
		verniy.MediaFieldSynonyms,
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
		return nil, stack.Wrap(ctx, err)
	}

	return &entity.Anime{
		ID:            data.ID,
		URL:           c.getURL("anime", data.ID),
		Title:         utils.PtrToStr(data.Title.Romaji),
		TitleEnglish:  utils.PtrToStr(data.Title.English),
		TitleJapanese: utils.PtrToStr(data.Title.Native),
		TitleSynonyms: data.Synonyms,
		Synopsis:      c.htmlToStr(utils.PtrToStr(data.Description)),
		Image:         utils.PtrToStr(data.CoverImage.Large),
		Score:         float64(utils.PtrToInt(data.AverageScore)) / 10,
		Member:        utils.PtrToInt(data.Popularity),
		Favorite:      utils.PtrToInt(data.Favourites),
		Type:          c.formatToStr(data.Format),
		Status:        c.statusToStr(data.Status),
		Episode:       utils.PtrToInt(data.Episodes),
		Ranking:       c.rankingToStr(data.Rankings),
		StartDate: entity.Date{
			Year:  utils.PtrToInt(data.StartDate.Year),
			Month: utils.PtrToInt(data.StartDate.Month),
			Day:   utils.PtrToInt(data.StartDate.Day),
		},
		EndDate: entity.Date{
			Year:  utils.PtrToInt(data.EndDate.Year),
			Month: utils.PtrToInt(data.EndDate.Month),
			Day:   utils.PtrToInt(data.EndDate.Day),
		},
	}, nil
}

// SearchAnime to search anime.
func (c *Client) SearchAnime(ctx context.Context, query string, page int) ([]entity.Anime, int, error) {
	data, err := c.client.SearchAnimeWithContext(ctx, verniy.PageParamMedia{
		Search: query,
	}, page, 10, verniy.MediaFieldID,
		verniy.MediaFieldTitle(
			verniy.MediaTitleFieldRomaji,
		),
		verniy.MediaFieldFormat,
		verniy.MediaFieldAverageScore,
	)
	if err != nil {
		return nil, 0, stack.Wrap(ctx, err)
	}

	res := make([]entity.Anime, len(data.Media))
	for i, d := range data.Media {
		res[i] = entity.Anime{
			ID:    d.ID,
			Title: utils.PtrToStr(d.Title.Romaji),
			Type:  c.formatToStr(d.Format),
			Score: float64(utils.PtrToInt(d.AverageScore)) / 10,
		}
	}

	return res, *data.PageInfo.Total, nil
}

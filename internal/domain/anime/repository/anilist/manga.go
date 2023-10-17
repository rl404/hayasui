package anilist

import (
	"context"

	"github.com/rl404/fairy/errors/stack"
	"github.com/rl404/hayasui/internal/domain/anime/entity"
	"github.com/rl404/hayasui/internal/utils"
	"github.com/rl404/verniy"
)

// GetManga to get manga.
func (c *Client) GetManga(ctx context.Context, id int) (*entity.Manga, error) {
	data, err := c.client.GetMangaWithContext(ctx, id,
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
		return nil, stack.Wrap(ctx, err)
	}

	return &entity.Manga{
		ID:            data.ID,
		URL:           c.getURL("manga", data.ID),
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
		Chapter:       utils.PtrToInt(data.Chapters),
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

// SearchManga to search manga.
func (c *Client) SearchManga(ctx context.Context, query string, page int) ([]entity.Manga, int, error) {
	data, err := c.client.SearchMangaWithContext(ctx, verniy.PageParamMedia{
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

	res := make([]entity.Manga, len(data.Media))
	for i, d := range data.Media {
		res[i] = entity.Manga{
			ID:    d.ID,
			Title: utils.PtrToStr(d.Title.Romaji),
			Type:  c.formatToStr(d.Format),
			Score: float64(utils.PtrToInt(d.AverageScore)) / 10,
		}
	}

	return res, *data.PageInfo.Total, nil
}

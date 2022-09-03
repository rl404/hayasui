package anilist

import (
	"context"

	"github.com/rl404/hayasui/internal/domain/anime/entity"
	"github.com/rl404/hayasui/internal/errors"
	"github.com/rl404/hayasui/internal/utils"
	"github.com/rl404/verniy"
)

// GetCharacter to get character.
func (c *Client) GetCharacter(ctx context.Context, id int) (*entity.Character, error) {
	data, err := c.client.GetCharacterWithContext(ctx, id,
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
		return nil, errors.Wrap(ctx, err)
	}

	return &entity.Character{
		ID:           data.ID,
		URL:          c.getURL("character", data.ID),
		Name:         utils.PtrToStr(data.Name.Full),
		NameJapanese: utils.PtrToStr(data.Name.Native),
		Nicknames:    data.Name.Alternative,
		Favorite:     utils.PtrToInt(data.Favourites),
		About:        utils.PtrToStr(data.Description),
		Image:        utils.PtrToStr(data.Image.Large),
	}, nil
}

// SearchCharacter to search character.
func (c *Client) SearchCharacter(ctx context.Context, query string, page int) ([]entity.Character, int, error) {
	data, err := c.client.SearchCharacterWithContext(ctx, verniy.PageParamCharacters{
		Search: query,
	}, page, 10,
		verniy.CharacterFieldID,
		verniy.CharacterFieldName(
			verniy.CharacterNameFieldFull,
		),
	)
	if err != nil {
		return nil, 0, errors.Wrap(ctx, err)
	}

	res := make([]entity.Character, len(data.Characters))
	for i, d := range data.Characters {
		res[i] = entity.Character{
			ID:   d.ID,
			Name: utils.PtrToStr(d.Name.Full),
		}
	}

	return res, *data.PageInfo.Total, nil
}

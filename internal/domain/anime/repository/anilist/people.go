package anilist

import (
	"context"

	"github.com/rl404/hayasui/internal/domain/anime/entity"
	"github.com/rl404/hayasui/internal/errors"
	"github.com/rl404/hayasui/internal/utils"
	"github.com/rl404/verniy"
)

// GetPeople to get people.
func (c *Client) GetPeople(ctx context.Context, id int) (*entity.People, error) {
	data, err := c.client.GetStaffWithContext(ctx, id,
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
		return nil, errors.Wrap(ctx, err)
	}

	return &entity.People{
		ID:               data.ID,
		URL:              c.getURL("staff", data.ID),
		Name:             utils.PtrToStr(data.Name.Full),
		AlternativeNames: data.Name.Alternative,
		Favorite:         utils.PtrToInt(data.Favourites),
		About:            utils.PtrToStr(data.Description),
		Image:            utils.PtrToStr(data.Image.Large),
		Birthday: entity.Date{
			Year:  utils.PtrToInt(data.DateOfBirth.Year),
			Month: utils.PtrToInt(data.DateOfBirth.Month),
			Day:   utils.PtrToInt(data.DateOfBirth.Day),
		},
	}, nil
}

// SearchPeople to search people.
func (c *Client) SearchPeople(ctx context.Context, query string, page int) ([]entity.People, int, error) {
	data, err := c.client.SearchStaffWithContext(ctx, verniy.PageParamStaff{
		Search: query,
	}, page, 10,
		verniy.StaffFieldID,
		verniy.StaffFieldName(
			verniy.StaffNameFieldFull,
		),
	)
	if err != nil {
		return nil, 0, errors.Wrap(ctx, err)
	}

	res := make([]entity.People, len(data.Staff))
	for i, d := range data.Staff {
		res[i] = entity.People{
			ID:   d.ID,
			Name: utils.PtrToStr(d.Name.Full),
		}
	}

	return res, *data.PageInfo.Total, nil
}

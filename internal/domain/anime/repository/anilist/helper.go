package anilist

import (
	"fmt"

	"github.com/rl404/hayasui/internal/domain/anime/entity"
	"github.com/rl404/hayasui/internal/utils"
	"github.com/rl404/verniy"
)

func (c *Client) getURL(path ...interface{}) string {
	url := "https://anilist.co"
	for _, p := range path {
		url += fmt.Sprintf("/%v", p)
	}
	return url
}

func (c *Client) formatToStr(f *verniy.MediaFormat) string {
	if f == nil {
		return ""
	}
	return entity.AnilistTypes[string(*f)]
}

func (c *Client) statusToStr(s *verniy.MediaStatus) string {
	if s == nil {
		return ""
	}
	return entity.AnilistStatuses[string(*s)]
}

func (c *Client) rankingToStr(rank []verniy.MediaRank) (a string) {
	for _, r := range rank {
		if utils.PtrToBool(r.AllTime) {
			a += fmt.Sprintf("#%s %s\n", utils.Thousands(r.Rank), r.Context)
		}
	}
	return a
}

func (c *Client) htmlToStr(str string) string {
	s, _ := c.converter.ConvertString(str)
	return s
}

package http

import (
	"fmt"

	"github.com/rl404/hayasui/internal/utils"
	"github.com/rl404/verniy"
)

func (h *http) formatToStr(f *verniy.MediaFormat) string {
	if f == nil {
		return ""
	}
	return string(*f)
}

func (h *http) statusToStr(s *verniy.MediaStatus) string {
	if s == nil {
		return ""
	}
	return string(*s)
}

func (h *http) getRanking(rank []verniy.MediaRank) (a []string) {
	for _, r := range rank {
		if utils.PtrToBool(r.AllTime) {
			a = append(a, fmt.Sprintf("#%s %s", utils.Thousands(r.Rank), r.Context))
		}
	}
	return a
}

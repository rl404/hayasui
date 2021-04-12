package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// ReadyHandler for handling after ready.
type ReadyHandler struct {
	prefix string
}

// NewReadyHandler to create new discord ready handler.
func NewReadyHandler(prefix string) *ReadyHandler {
	return &ReadyHandler{
		prefix: prefix,
	}
}

// Handler to get handler function.
func (h *ReadyHandler) Handler() func(*discordgo.Session, *discordgo.Ready) {
	return func(s *discordgo.Session, _ *discordgo.Ready) {
		s.UpdateListeningStatus(fmt.Sprintf("%shelp for command list", h.prefix))
	}
}

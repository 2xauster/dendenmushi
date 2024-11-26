package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/jauster101/dendenmushi/core/logger"
)

func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {
	username := s.State.User.Username
	logger.Info(fmt.Sprintf("Logged in as %s", username))
}
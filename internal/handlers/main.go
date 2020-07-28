package handlers

import (
	"github.com/arturoguerra/rolewatcher/internal/config"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type (
	handlers struct {
		Logger *logrus.Logger
		Config *config.Config
	}

	// Handlers exports functions to discordgo
	Handlers interface {
		OnReady(*discordgo.Session, *discordgo.Ready)
		ReactionRolesAdd(*discordgo.Session, *discordgo.MessageReactionAdd)
		ReactionRolesRemove(*discordgo.Session, *discordgo.MessageReactionRemove)
		RoleWatcher(*discordgo.Session, *discordgo.GuildMemberUpdate)
		MemberChunks(*discordgo.Session, *discordgo.GuildMembersChunk)
	}
)

// New creates a new discord event handler
func New(log *logrus.Logger, cfg *config.Config) (Handlers, error) {
	return &handlers{
		Logger: log,
		Config: cfg,
	}, nil
}

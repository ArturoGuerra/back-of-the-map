package handlers

import "github.com/bwmarrin/discordgo"

func (h *handlers) OnReady(s *discordgo.Session, event *discordgo.Ready) {
	h.Logger.Info("Starting RoleWatcher (Back of the Map Bot :)")

	if guild, err := s.Guild(h.Config.GuildID); err == nil && guild != nil {
		h.Logger.Infof("Processing members...")

		for _, member := range guild.Members {
			h.roleHandler(s, member)
		}

		h.Logger.Info("Done processing %d members", len(guild.Members))
	}
}

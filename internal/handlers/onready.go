package handlers

import "github.com/bwmarrin/discordgo"

func (h *handlers) OnReady(s *discordgo.Session, event *discordgo.Ready) {
	h.Logger.Info("Starting RoleWatcher Back of the Map Bot :)")

	if err := s.RequestGuildMembers(h.Config.GuildID, "", 0, false); err != nil {
		h.Logger.Error(err)
	}
}

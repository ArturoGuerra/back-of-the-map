package handlers

import "github.com/bwmarrin/discordgo"

func (h *handlers) MemberChunks(s *discordgo.Session, chunk *discordgo.GuildMembersChunk) {
	if chunk.GuildID == h.Config.GuildID {
		h.Logger.Info("Processing members...")

		for _, member := range chunk.Members {
			h.roleHandler(s, member)
		}

		h.Logger.Infof("Done processing %d members", len(chunk.Members))
	}
}

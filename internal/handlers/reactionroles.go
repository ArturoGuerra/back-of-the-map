package handlers

import "github.com/bwmarrin/discordgo"

func (h *handlers) reactionRoles(s *discordgo.Session, mr *discordgo.MessageReaction, addrole bool) {

	for _, rr := range h.Config.ReactionRoles {
		if rr.Emoji == mr.Emoji.ID && mr.MessageID == rr.Message {
			if addrole {
				s.GuildMemberRoleAdd(mr.GuildID, mr.UserID, rr.Role)
			} else {
				s.GuildMemberRoleRemove(mr.GuildID, mr.UserID, rr.Role)
			}
		}
	}
}

func (h *handlers) ReactionRolesAdd(s *discordgo.Session, mr *discordgo.MessageReactionAdd) {
	h.reactionRoles(s, mr.MessageReaction, true)

}

func (h *handlers) ReactionRolesRemove(s *discordgo.Session, mr *discordgo.MessageReactionRemove) {
	h.reactionRoles(s, mr.MessageReaction, false)

}

package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (h *handlers) roleHandler(s *discordgo.Session, m *discordgo.Member) {
	userString := fmt.Sprintf("%s#%s-%s", m.User.Username, m.User.Discriminator, m.User.ID)
	for _, rolemapper := range h.Config.Roles {
		hasRole := hasRole(rolemapper.Give, m.Roles)
		needsRole := hasRoles(m.Roles, rolemapper.Watch)

		if needsRole && !hasRole {
			if err := s.GuildMemberRoleAdd(m.GuildID, m.User.ID, rolemapper.Give); err != nil {
				h.Logger.Errorf(err.Error())
			} else {
				h.Logger.Infof("Giving %s to %s", rolemapper.Give, userString)
			}
		} else if !needsRole && hasRole {
			if err := s.GuildMemberRoleRemove(m.GuildID, m.User.ID, rolemapper.Give); err != nil {
				h.Logger.Errorf(err.Error())
			} else {
				h.Logger.Infof("Taking %s from %s", rolemapper.Give, userString)
			}
		}
	}
}

func (h *handlers) RoleWatcher(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {
	h.roleHandler(s, m.Member)
}

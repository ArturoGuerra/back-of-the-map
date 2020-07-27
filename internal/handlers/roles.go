package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func hasRole(role string, roles []string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}

	return false
}

func hasRoles(r1 []string, r2 []string) bool {
	for _, rr1 := range r1 {
		for _, rr2 := range r2 {
			if rr2 == rr1 {
				return true
			}
		}
	}

	return false
}

func (h *handlers) roleHandler(s *discordgo.Session, m *discordgo.Member) {
	userString := fmt.Sprintf("%s#%s-%s", m.User.Username, m.User.Discriminator, m.User.ID)
	for _, rolemapper := range h.Config.Roles {
		hasGive := hasRole(rolemapper.Give, m.Roles)
		rolecheck := hasRoles(m.Roles, rolemapper.Watch)

		give := (rolecheck && !hasGive)
		take := (!rolecheck && hasGive)

		if give {
			if err := s.GuildMemberRoleAdd(m.GuildID, m.User.ID, rolemapper.Give); err != nil {
				h.Logger.Errorf(err.Error())
			} else {
				h.Logger.Infof("Giving %s to %s", rolemapper.Give, userString)
			}
		} else if take {
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

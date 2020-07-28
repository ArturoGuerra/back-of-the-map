package handlers

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

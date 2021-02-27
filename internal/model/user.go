package model

type User struct {
	ID    string   `json:"user_id" binding:"required"`
	Email string   `json:"email" binding:"required"`
	Scope []string `json:"scope" binding:"required"`
	Role  string   `json:"role"`
}

func (user User) Can(permissions ...string) bool {
	scopeIndex := map[string]bool{}
	for idx := range user.Scope {
		scopeIndex[user.Scope[idx]] = true
	}

	for idx := range permissions {
		if _, ok := scopeIndex[permissions[idx]]; ok {
			return true
		}
	}
	return false
}

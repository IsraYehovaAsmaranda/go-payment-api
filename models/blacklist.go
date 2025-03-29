package models

type BlacklistedToken struct {
	Token string `json:"token"`
}

type BlacklistedTokenCollection struct {
	BlacklistedTokens []BlacklistedToken `json:"blacklisted_token"`
}

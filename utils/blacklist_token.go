package utils

import (
	"encoding/json"
	"os"

	"github.com/IsraYehovaAsmaranda/go-payment-api/models"
)

var blacklistFile = "storage/blacklist.json"

func GetBlackListedTokens() (models.BlacklistedTokenCollection, error) {
	file, err := os.ReadFile(blacklistFile)
	if err != nil {
		return models.BlacklistedTokenCollection{}, err
	}

	var tokens models.BlacklistedTokenCollection
	err = json.Unmarshal(file, &tokens)
	if err != nil {
		return models.BlacklistedTokenCollection{}, err
	}

	return tokens, nil
}

func BlacklistToken(token string) error {
	tokens, err := GetBlackListedTokens()
	if err != nil {
		return err
	}

	tokens.BlacklistedTokens = append(tokens.BlacklistedTokens, models.BlacklistedToken{
		Token: token,
	})

	data, err := json.MarshalIndent(tokens, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(blacklistFile, data, 0644)
}

func IsTokenBlacklisted(token string) bool {
	tokens, err := GetBlackListedTokens()
	if err != nil {
		return false
	}

	for _, t := range tokens.BlacklistedTokens {
		if t.Token == token {
			return true
		}
	}

	return false
}

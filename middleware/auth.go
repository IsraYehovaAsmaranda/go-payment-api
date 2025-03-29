package middleware

import (
	"net/http"
	"strings"

	"github.com/IsraYehovaAsmaranda/go-payment-api/helpers"
	"github.com/IsraYehovaAsmaranda/go-payment-api/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized: Missing Authorization Header", "Unauthorized")
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if utils.IsTokenBlacklisted(token) {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized: Token is blacklisted", "Unauthorized")
			return
		}

		_, err := utils.VerifyJWT(token)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, err.Error(), "Invalid Token")
			return
		}

		next.ServeHTTP(w, r)
	})
}

package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/IsraYehovaAsmaranda/go-payment-api/helpers"
	"github.com/IsraYehovaAsmaranda/go-payment-api/models"
	"github.com/IsraYehovaAsmaranda/go-payment-api/utils"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error(), "Invalid Request")
		return
	}

	users, err := readUsersFromJson()
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to register user")
		return
	}

	for _, user := range users.Data {
		if user.Username == newUser.Username {
			helpers.RespondWithError(w, http.StatusBadRequest, "Username already exists", "Username already exists")
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to register user")
		return
	}

	newUser.Password = string(hashedPassword)

	users.Data = append(users.Data, newUser)

	err = saveUsersToJson(users)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to register user")
		return
	}

	response := models.RegisterResponse{
		Username: newUser.Username,
		Name:     newUser.Name,
	}
	helpers.RespondWithJSON(w, http.StatusCreated, response, "User Registered Successfully")
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error(), "Invalid request")
		return
	}

	users, err := readUsersFromJson()
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to login user")
		return
	}

	for _, user := range users.Data {
		if user.Username == loginRequest.Username {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
			if err != nil {
				helpers.RespondWithError(w, http.StatusUnauthorized, err.Error(), "Invalid credentials")
				return
			}

			token, err := utils.GenerateJWT(user.Username)
			if err != nil {
				helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to generate token")
				return
			}

			response := models.LoginResponse{
				Username: user.Username,
				Name:     user.Name,
				Token:    token,
			}

			helpers.RespondWithJSON(w, http.StatusOK, response, "Login successful")
			return
		}
	}

	helpers.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials", "Invalid credentials")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	token := strings.TrimPrefix(authHeader, "Bearer ")
	err := utils.BlacklistToken(token)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to logout user")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, nil, "Logout successful")
}

func readUsersFromJson() (models.UserCollection, error) {
	file, err := os.ReadFile("storage/users.json")
	if err != nil {
		return models.UserCollection{}, err
	}

	var users models.UserCollection
	err = json.Unmarshal(file, &users)
	if err != nil {
		return models.UserCollection{}, err
	}

	return users, nil
}

func saveUsersToJson(users models.UserCollection) error {
	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("storage/users.json", data, 0644)
}

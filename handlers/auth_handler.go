package handlers

import (
	"encoding/json"
	"net/http"
	"os"

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
	}

	users, err := readUsersFromJson()
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to login user")
		return
	}

	for _, user := range users.Data {
		if user.Username == loginRequest.Username && user.Password == loginRequest.Password {
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

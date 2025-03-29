package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/IsraYehovaAsmaranda/go-payment-api/helpers"
	"github.com/IsraYehovaAsmaranda/go-payment-api/models"
	"github.com/IsraYehovaAsmaranda/go-payment-api/utils"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var registerRequest models.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		utils.SaveActivityLog("Register Failed - Invalid Request", models.User{})
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error(), "Invalid Request")
		return
	}

	users, err := models.ReadUsersFromJson()
	if err != nil {
		utils.SaveActivityLog("Register Failed - Failed to register user", models.User{})
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to register user")
		return
	}

	for _, user := range users.Data {
		if user.Username == registerRequest.Username {
			utils.SaveActivityLog("Register Failed - Username already exists", models.User{})
			helpers.RespondWithError(w, http.StatusBadRequest, "Username already exists", "Username already exists")
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SaveActivityLog("Register Failed - Failed to encrypt password", models.User{})
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed encrypt password")
		return
	}

	var newUser models.User = models.User{
		Username: registerRequest.Username,
		Name:     registerRequest.Name,
		Password: string(hashedPassword),
		Balance:  0,
	}
	users.Data = append(users.Data, newUser)

	err = models.SaveUsersToJson(users)
	if err != nil {
		utils.SaveActivityLog("Register Failed - Failed to register user", models.User{})
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to register user")
		return
	}

	response := models.RegisterResponse{
		Username: registerRequest.Username,
		Name:     registerRequest.Name,
	}

	utils.SaveActivityLog("User Registered Successfully", newUser)
	helpers.RespondWithJSON(w, http.StatusCreated, response, "User Registered Successfully")
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		utils.SaveActivityLog("Login Failed - Invalid Request", models.User{})
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error(), "Invalid request")
		return
	}

	users, err := models.ReadUsersFromJson()
	if err != nil {
		utils.SaveActivityLog("Login Failed - Failed to login user", models.User{})
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to login user")
		return
	}

	for _, user := range users.Data {
		if user.Username == loginRequest.Username {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
			if err != nil {
				utils.SaveActivityLog("Login Failed - Invallid Credentials", user)
				helpers.RespondWithError(w, http.StatusUnauthorized, err.Error(), "Invalid credentials")
				return
			}

			token, err := utils.GenerateJWT(user.Username)
			if err != nil {
				utils.SaveActivityLog("Login Failed - Failed to generate token", user)
				helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to generate token")
				return
			}

			response := models.LoginResponse{
				Username: user.Username,
				Name:     user.Name,
				Token:    token,
			}

			utils.SaveActivityLog("Login Successful", user)
			helpers.RespondWithJSON(w, http.StatusOK, response, "Login successful")
			return
		}
	}

	utils.SaveActivityLog("Login Failed - Invalid Credentials", models.User{})
	helpers.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials", "Invalid credentials")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	token := strings.TrimPrefix(authHeader, "Bearer ")
	err := utils.BlacklistToken(token)
	if err != nil {
		utils.SaveActivityLog("Logout Failed - Failed to logout user", models.User{})
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to logout user")
		return
	}

	username, _ := utils.GetUsernameFromToken(token)
	user, err := models.GetUserByUsername(username)
	if err != nil {
		utils.SaveActivityLog("Logout Failed - User not found", models.User{})
		helpers.RespondWithError(w, http.StatusNotFound, err.Error(), "User not found")
		return
	}

	utils.SaveActivityLog("Logout Successful", user)
	helpers.RespondWithJSON(w, http.StatusOK, nil, "Logout successful")
}

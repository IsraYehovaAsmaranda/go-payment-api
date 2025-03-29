package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/IsraYehovaAsmaranda/go-payment-api/helpers"
	"github.com/IsraYehovaAsmaranda/go-payment-api/models"
	"github.com/IsraYehovaAsmaranda/go-payment-api/utils"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	var paymentRequest models.PaymentRequest

	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		utils.SaveActivityLog("Payment Failed - Invalid Request", models.User{})
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error(), "Invalid Request")
		return
	}

	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	username, _ := utils.GetUsernameFromToken(token)

	user, err := models.GetUserByUsername(username)
	if err != nil {
		utils.SaveActivityLog("Payment Failed - User not found", models.User{})
		helpers.RespondWithError(w, http.StatusNotFound, err.Error(), "User Not Found")
		return
	}

	if username == paymentRequest.Username {
		utils.SaveActivityLog("Payment Failed - User tried to transfer to their own account", user)
		helpers.RespondWithError(w, http.StatusBadRequest, "You cannot transfer to your own account", "You cannot transfer to your own account")
		return
	}

	targetUser, err := models.GetUserByUsername(paymentRequest.Username)
	if err != nil {
		utils.SaveActivityLog("Payment Failed - Target User not found", user)
		helpers.RespondWithError(w, http.StatusNotFound, err.Error(), "Target User Not Found")
		return
	}

	if user.Balance < paymentRequest.Amount {
		utils.SaveActivityLog("Payment Failed - Insufficient balance", user)
		helpers.RespondWithError(w, http.StatusBadRequest, "Insufficient balance", "Insufficient balance")
		return
	}

	err = models.UpdateUserBalance(username, user.Balance-paymentRequest.Amount)
	if err != nil {
		utils.SaveActivityLog("Payment Failed - Failed to update user balance", user)
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to update user balance")
		return
	}

	err = models.UpdateUserBalance(paymentRequest.Username, targetUser.Balance+paymentRequest.Amount)
	if err != nil {
		utils.SaveActivityLog("Payment Failed - Failed to update target user balance", user)
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to update target user balance")
		return
	}

	userResponse := models.UserResponse{
		Username: targetUser.Username,
		Name:     targetUser.Name,
	}

	newPayment := models.Payment{
		Id:         time.Now().Nanosecond(),
		Amount:     paymentRequest.Amount,
		TargetUser: userResponse,
	}

	err = models.SavePaymentToJSON(newPayment)
	if err != nil {
		utils.SaveActivityLog("Payment Failed - Failed to save payment", user)
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), "Failed to save payment")
		return
	}

	utils.SaveActivityLog("Payment Successful", user)
	helpers.RespondWithJSON(w, http.StatusCreated, newPayment, "Payment successful")
}

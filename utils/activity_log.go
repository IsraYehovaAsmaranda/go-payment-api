package utils

import (
	"fmt"
	"time"

	"github.com/IsraYehovaAsmaranda/go-payment-api/models"
)

func SaveActivityLog(activity string, user models.User) {
	activityLog := models.ActivityLog{
		Activity:  activity,
		CreatedAt: time.Now(),
		CreatedBy: models.UserResponse{
			Username: user.Username,
			Name:     user.Name,
		},
	}

	err := models.SaveActivityLogToJSON(activityLog)
	if err != nil {
		fmt.Println("Failed to save activity log:", err)
	}
}

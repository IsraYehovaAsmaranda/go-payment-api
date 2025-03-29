package models

import (
	"encoding/json"
	"os"
	"time"
)

type ActivityLog struct {
	ID        int64        `json:"id"`
	Activity  string       `json:"activity"`
	CreatedAt time.Time    `json:"created_at"`
	CreatedBy UserResponse `json:"created_by"`
}

type ActivityLogCollection struct {
	Data []ActivityLog `json:"data"`
}

func GetAllActivityLogs() (ActivityLogCollection, error) {
	file, err := os.ReadFile("storage/activity_logs.json")
	if err != nil {
		return ActivityLogCollection{}, err
	}

	var activityLogs ActivityLogCollection
	err = json.Unmarshal(file, &activityLogs)
	if err != nil {
		return ActivityLogCollection{}, err
	}

	return activityLogs, nil
}

func SaveActivityLogToJSON(activityLog ActivityLog) error {
	activityLogsCollection, err := GetAllActivityLogs()
	if err != nil {
		return err
	}

	activityLogsCollection.Data = append(activityLogsCollection.Data, activityLog)

	data, err := json.MarshalIndent(activityLogsCollection, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("storage/activity_logs.json", data, 0644)
}

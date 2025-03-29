package models

import (
	"encoding/json"
	"errors"
	"os"
)

type User struct {
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

type UserCollection struct {
	Data []User `json:"data"`
}

type UserResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

func ReadUsersFromJson() (UserCollection, error) {
	file, err := os.ReadFile("storage/users.json")
	if err != nil {
		return UserCollection{}, err
	}

	var users UserCollection
	err = json.Unmarshal(file, &users)
	if err != nil {
		return UserCollection{}, err
	}

	return users, nil
}

func SaveUsersToJson(users UserCollection) error {
	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("storage/users.json", data, 0644)
}

func GetUserByUsername(username string) (User, error) {
	users, err := ReadUsersFromJson()
	if err != nil {
		return User{}, err
	}

	for _, user := range users.Data {
		if user.Username == username {
			return user, nil
		}
	}

	return User{}, errors.New("user not found")
}

func UpdateUserBalance(username string, balance float64) error {
	users, err := ReadUsersFromJson()
	if err != nil {
		return err
	}

	for i, user := range users.Data {
		if user.Username == username {
			users.Data[i].Balance = balance
			return SaveUsersToJson(users)
		}
	}

	return nil
}

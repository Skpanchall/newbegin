package storage

import (
	"encoding/json"
	"os"

	"github.com/Skpanchall/newbegin/simplerProject/model"
)

func SaveUserToFile(users map[int]model.User) error {
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("user.json", jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
func GetUsersFromFile() (map[int]model.User, error) {
	jsonFile, err := os.ReadFile("user.json")
	if err != nil {
		return nil, err
	}
	var users map[int]model.User
	err = json.Unmarshal(jsonFile, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

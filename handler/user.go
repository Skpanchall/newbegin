package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Skpanchall/newbegin/model"
	"github.com/Skpanchall/newbegin/storage"
	"github.com/Skpanchall/newbegin/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		users = make(map[int]model.User)
	}
	json.NewEncoder(w).Encode(users)
}
func WelcomeAPI(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(w, "Welcome to Go API")
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(w, "Get User API")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		users = make(map[int]model.User)
	}
	payload := model.User{}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.SendResponse(w, err.Error())
		return
	}
	users[payload.ID] = payload
	err = storage.SaveUserToFile(users)
	if err != nil {
		utils.SendResponse(w, err.Error())
		return
	}
	json.NewEncoder(w).Encode(users)
}

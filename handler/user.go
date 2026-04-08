package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		utils.SendResponse(w, err.Error())
		return
	}
	// get a query parameter id
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.SendResponse(w, "id is required")
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.SendResponse(w, "invalid id")
		return
	}
	payload := model.User{}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.SendResponse(w, err.Error())
		return
	}

	exist, ok := users[idInt]
	if !ok {
		utils.SendResponse(w, "User not found")
		return
	}
	if payload.Name != "" {
		exist.Name = payload.Name
	}
	if payload.Email != "" {
		exist.Email = payload.Email
	}
	users[idInt] = exist
	err = storage.SaveUserToFile(users)
	if err != nil {
		utils.SendResponse(w, err.Error())
		return
	}
	json.NewEncoder(w).Encode(users)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		utils.SendResponse(w, err.Error())
		return
	}
	// get a query parameter id
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.SendResponse(w, "id is required")
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.SendResponse(w, "invalid id")
		return
	}
	delete(users, idInt)
	err = storage.SaveUserToFile(users)
	if err != nil {
		utils.SendResponse(w, err.Error())
		return
	}
}
func HandleUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPut {
		UpdateUser(w, r)
		return
	} else if r.Method == http.MethodDelete {
		DeleteUser(w, r)
		return
	}
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

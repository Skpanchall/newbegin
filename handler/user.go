package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Skpanchall/newbegin/model"
	"github.com/Skpanchall/newbegin/storage"
	"github.com/Skpanchall/newbegin/utils"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUsers(w, r)
	default:
		utils.SendResponse(w, "Method not allowed")
	}
}
func HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUserByID(w, r)
	case http.MethodPut:
		UpdateUser(w, r)
	case http.MethodDelete:
		DeleteUser(w, r)
	case http.MethodPost:
		CreateUser(w, r)
	default:
		utils.SendResponse(w, "Method not allowed")
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		users = make(map[int]model.User)
	}
	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
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
	defer r.Body.Close()

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
	w.WriteHeader(http.StatusOK)
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
	if _, ok := users[idInt]; !ok {
		utils.SendResponse(w, "User not found")
		return
	}
	delete(users, idInt)
	err = storage.SaveUserToFile(users)
	if err != nil {
		utils.SendResponse(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	utils.SendResponse(w, "User deleted successfully")
}
func GetUserByID(w http.ResponseWriter, r *http.Request) {

	users, err := storage.GetUsersFromFile()
	if err != nil {
		users = make(map[int]model.User)
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		utils.SendResponse(w, "invalid id")
		return
	}
	if user, ok := users[id]; ok {
		json.NewEncoder(w).Encode(user)
		return
	}
	utils.SendResponse(w, "User not found")
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

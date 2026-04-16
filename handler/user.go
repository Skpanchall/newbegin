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
		utils.SendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
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
		utils.SendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		users = make(map[int]model.User)
	}
	utils.SendSuccessResponse(w, users, http.StatusOK)
}
func WelcomeAPI(w http.ResponseWriter, r *http.Request) {
	utils.SendSuccessResponse(w, "Welcome to Go API", http.StatusOK)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// get a query parameter id
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.SendErrorResponse(w, "id is required", http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.SendErrorResponse(w, "invalid id", http.StatusBadRequest)
		return
	}
	payload := model.User{}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	exist, ok := users[idInt]
	if !ok {
		utils.SendErrorResponse(w, "User not found", http.StatusNotFound)
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
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendSuccessResponse(w, users, http.StatusOK)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// get a query parameter id
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.SendErrorResponse(w, "id is required", http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.SendErrorResponse(w, "invalid id", http.StatusBadRequest)
		return
	}
	if _, ok := users[idInt]; !ok {
		utils.SendErrorResponse(w, "User not found", http.StatusNotFound)
		return
	}
	delete(users, idInt)
	err = storage.SaveUserToFile(users)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendSuccessResponse(w, "User deleted successfully", http.StatusOK)
}
func GetUserByID(w http.ResponseWriter, r *http.Request) {

	users, err := storage.GetUsersFromFile()
	if err != nil {
		users = make(map[int]model.User)
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		utils.SendErrorResponse(w, "invalid id", http.StatusBadRequest)
		return
	}
	if user, ok := users[id]; ok {
		utils.SendSuccessResponse(w, user, http.StatusOK)
		return
	}
	utils.SendErrorResponse(w, "User not found", http.StatusNotFound)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		users = make(map[int]model.User)
	}
	payload := model.User{}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	users[payload.ID] = payload
	err = storage.SaveUserToFile(users)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendSuccessResponse(w, users, http.StatusOK)
}

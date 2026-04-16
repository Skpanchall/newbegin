package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Skpanchall/newbegin/model"
	"github.com/Skpanchall/newbegin/storage"
	"github.com/Skpanchall/newbegin/utils"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return GetUsers(w, r)
	default:
		return &utils.ErrError{Message: "Method not allowed", Code: http.StatusMethodNotAllowed}
	}

}
func HandleUser(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return GetUserByID(w, r)
	case http.MethodPut:
		return UpdateUser(w, r)
	case http.MethodDelete:
		return DeleteUser(w, r)
	case http.MethodPost:
		return CreateUser(w, r)
	default:
		return &utils.ErrError{Message: "Method not allowed", Code: http.StatusMethodNotAllowed}
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		return &utils.ErrError{Message: err.Error(), Code: http.StatusInternalServerError}
	}
	utils.SendSuccessResponse(w, users, http.StatusOK)
	return nil
}
func WelcomeAPI(w http.ResponseWriter, r *http.Request) {
	utils.SendSuccessResponse(w, "Welcome to Go API", http.StatusOK)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) error {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		return &utils.ErrError{Message: err.Error(), Code: http.StatusInternalServerError}
	}
	// get a query parameter id
	id := r.URL.Query().Get("id")
	if id == "" {
		return &utils.ErrError{Message: "id is required", Code: http.StatusBadRequest}
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return &utils.ErrError{Message: "invalid id", Code: http.StatusBadRequest}
	}
	payload := model.User{}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return &utils.ErrError{Message: err.Error(), Code: http.StatusBadRequest}

	}
	defer r.Body.Close()

	exist, ok := users[idInt]
	if !ok {
		return &utils.ErrError{Message: "User not found", Code: http.StatusNotFound}
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
		return &utils.ErrError{Message: err.Error(), Code: http.StatusInternalServerError}
	}
	utils.SendSuccessResponse(w, users, http.StatusOK)
	return nil
}
func DeleteUser(w http.ResponseWriter, r *http.Request) error {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		return &utils.ErrError{Message: err.Error(), Code: http.StatusInternalServerError}
	}
	// get a query parameter id
	id := r.URL.Query().Get("id")
	if id == "" {
		return &utils.ErrError{Message: "id is required", Code: http.StatusBadRequest}
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return &utils.ErrError{Message: "invalid id", Code: http.StatusBadRequest}
	}
	if _, ok := users[idInt]; !ok {
		return &utils.ErrError{Message: "User not found", Code: http.StatusNotFound}
	}
	delete(users, idInt)
	err = storage.SaveUserToFile(users)
	if err != nil {
		return &utils.ErrError{Message: err.Error(), Code: http.StatusInternalServerError}
	}
	utils.SendSuccessResponse(w, "User deleted successfully", http.StatusOK)
	return nil
}
func GetUserByID(w http.ResponseWriter, r *http.Request) error {

	users, err := storage.GetUsersFromFile()
	if err != nil {
		users = make(map[int]model.User)
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return &utils.ErrError{Message: "invalid id", Code: http.StatusBadRequest}
	}
	if user, ok := users[id]; ok {
		utils.SendSuccessResponse(w, user, http.StatusOK)
		return nil
	}
	return &utils.ErrError{Message: "User not found", Code: http.StatusNotFound}
}

func CreateUser(w http.ResponseWriter, r *http.Request) error {
	users, err := storage.GetUsersFromFile()
	if err != nil {
		users = make(map[int]model.User)
	}
	payload := model.User{}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return &utils.ErrError{Message: err.Error(), Code: http.StatusBadRequest}
	}
	users[payload.ID] = payload
	err = storage.SaveUserToFile(users)
	if err != nil {
		return &utils.ErrError{Message: err.Error(), Code: http.StatusInternalServerError}
	}
	utils.SendSuccessResponse(w, users, http.StatusOK)
	return nil
}

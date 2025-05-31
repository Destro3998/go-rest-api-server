package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Destro3998/go-rest-api/models"
	"github.com/gorilla/mux"
)

var users = []models.User{}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, u := range users {
		if u.ID == params["id"] {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.NotFound(w, r)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, u := range users {
		if u.ID == params["id"] {
			json.NewDecoder(r.Body).Decode(&users[i])
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, u := range users {
		if u.ID == params["id"] {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

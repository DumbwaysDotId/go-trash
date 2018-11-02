package handler

import (
	"encoding/json"
	"first-go/db"
	"first-go/model"
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	defer db.Close()

	var users []model.User
	db.Find(&users)

	respondJSON(w, http.StatusOK, users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	defer db.Close()

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)

	respondJSON(w, http.StatusOK, "success create")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Called")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Called")
}

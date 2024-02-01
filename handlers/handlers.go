package handlers

import (
	"encoding/json"
	"fmt"
	"golang/helpers"
	"golang/models"
	"golang/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandlerHomePage(w http.ResponseWriter, r *http.Request) {
	message := models.Text{"Welcome to API Exercises"}
	data, err := json.Marshal(message)
	helpers.CheckError(err)
	fmt.Fprint(w, string(data))
}

func HandlerGetUsers(w http.ResponseWriter, r *http.Request) {
	result := services.GetUsers()
	datas, err := json.Marshal(result)
	helpers.CheckError(err)
	fmt.Fprint(w, string(datas))
}

func HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	userID, err := strconv.Atoi(variables["id"])
	helpers.CheckError(err)

	result := services.GetUser(userID)
	data, err := json.Marshal(result)
	fmt.Fprintf(w, string(data))
}

func HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.Users
	err := json.NewDecoder(r.Body).Decode(&newUser)
	helpers.CheckError(err)

	responseData, _ := json.Marshal(newUser)

	services.CreateUser(newUser)
	fmt.Fprintf(w, string(responseData))
}

func HandlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	userId, err := strconv.Atoi(variables["id"])
	helpers.CheckError(err)

	services.DeleteUser(userId)

}

func HandlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	userID, err := strconv.Atoi(variables["id"])
	helpers.CheckError(err)

	var userData models.Users
	err = json.NewDecoder(r.Body).Decode(&userData)

	userData.ID = userID

	responseData, _ := json.Marshal(userData)

	services.UpdateUser(userData, userID)

	fmt.Fprintf(w, string(responseData))

}

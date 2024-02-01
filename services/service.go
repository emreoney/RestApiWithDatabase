package services

import (
	"golang/database"
	"golang/models"
)

func GetUsers() []models.Users {
	return database.GetUsersFromDb()
}

func GetUser(userID int) models.Users {
	return database.GetUserFromDb(userID)
}

func CreateUser(user models.Users) {
	database.InsertUserToDb(user)
}

func UpdateUser(user models.Users, userID int) {
	database.UpdateUserInDb(user, userID)
}

func DeleteUser(userID int) {
	database.DeleteUserFromDb(userID)
}

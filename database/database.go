package database

import (
	"database/sql"
	"fmt"
	"golang/helpers"
	"golang/models"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func getDbValues() {
	err := godotenv.Load()
	helpers.CheckError(err)
}

func ConnectWithDb() {
	var err error

	getDbValues()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", connectionString)
	helpers.CheckError(err)
}

func GetUsersFromDb() []models.Users {
	row, err := db.Query("SELECT * FROM USERS")
	helpers.CheckError(err)

	var userDatas []models.Users

	for row.Next() {
		var userData models.Users
		err := row.Scan(&userData.ID, &userData.Firstname, &userData.Lastname, &userData.Age)
		helpers.CheckError(err)
		userDatas = append(userDatas, userData)
	}

	return userDatas
}

func GetUserFromDb(userId int) models.Users {
	row, err := db.Query("SELECT * FROM USERS WHERE ID=$1", userId)
	helpers.CheckError(err)

	var usr models.Users

	for row.Next() {
		err := row.Scan(&usr.ID, &usr.Firstname, &usr.Lastname, &usr.Age)
		helpers.CheckError(err)
	}

	return usr
}

func InsertUserToDb(userData models.Users) {
	row, err := db.Exec("insert into users (id,firstname,lastname,age) values ($1,$2,$3,$4)", userData.ID, userData.Firstname, userData.Lastname, userData.Age)
	helpers.CheckError(err)

	countAffectedRow, err := row.RowsAffected()
	helpers.CheckError(err)

	fmt.Printf("Etkilenen satir sayisi %d", countAffectedRow)
}

func DeleteUserFromDb(userId int) {
	row, err := db.Exec("DELETE FROM USERS WHERE ID=$1", userId)
	helpers.CheckError(err)

	countAffectedRow, _ := row.RowsAffected()

	fmt.Printf("Etkilenen satir sayisi: %d", countAffectedRow)
}

func UpdateUserInDb(userData models.Users, userId int) {
	row, err := db.Exec("update users set firstname=$1, lastname=$2 , age=$3 where id=$4", userData.Firstname, userData.Lastname, userData.Age, userId)
	helpers.CheckError(err)

	countAffectedRow, _ := row.RowsAffected()
	fmt.Printf("Etkilenen satir sayisi: %d", countAffectedRow)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// const DNS = "root@127.0.0.1:3306/godb"
const DNS = "root:root@tcp(127.0.0.1:3306)/godb"

type User struct {
	gorm.Model
	name  string `json:"name"`
	email string `json:"email"`
}

func InitialMigration() {

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to the database")
	}
	DB.AutoMigrate(&User{})

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

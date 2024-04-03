package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// const DNS = "root@127.0.0.1:3306/godb"
//const DNS = "root:root@tcp(127.0.0.1:3306)/godb"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func InitialMigration() {

	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	if dbUser == "" || dbPass == "" {
		panic("Failed to read the DB_USER or DB_PASS environment variables")
	}

	dns := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/godb", dbUser, dbPass)

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
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

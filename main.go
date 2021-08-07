// TODO - create email template
// TODO - create Docker package

// TODO - change json responses
// * Updated json response for creating warehouse and sending email *
package main

import (
	"api/config"
	"api/models"
	"api/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var Warehouses []models.Warehouse

var db *sql.DB

func main() {
	godotenv.Load()
	fmt.Println("Welcome")
	db = config.ConnectDB()
	authDB := utils.AuthorizeFunc(db)
	fmt.Println(db)

	r := config.CreateRouter(db, authDB)

	log.Fatal(http.ListenAndServe(":8080", r))
}

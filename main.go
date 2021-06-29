//TODO function for updating and deleting magazine
package main

import (
	"api/config"
	"api/models"
	"api/routes"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var Magazines []models.Magazine

var db *sql.DB

func main() {
	godotenv.Load()
	fmt.Println("Welcome")
	db = config.ConnectDB()
	fmt.Println(db)
	r := mux.NewRouter()

	r.HandleFunc("/getMagazines", routes.GetAllMagazines(db)).Methods("GET")
	r.HandleFunc("/createMagazine", routes.CreateMagazine(db)).Methods("POST")
	r.HandleFunc("/reserveMagazine", routes.ReserveMagazine(db)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}

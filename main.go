//TODO function for updating and deleting magazine
package main

import (
	"api/config"
	"api/models"
	"api/routes"
	"api/utils"
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
	authDB := utils.AuthorizeFunc(db)
	fmt.Println(db)
	r := mux.NewRouter()

	r.HandleFunc("/getMagazines", routes.GetAllMagazines(db, authDB)).Methods("GET")
	r.HandleFunc("/createMagazine", routes.CreateMagazine(db, authDB)).Methods("POST")
	r.HandleFunc("/reserveMagazine", routes.ReserveMagazine(db, authDB)).Methods("PUT")
	r.HandleFunc("/getReservedMagazine", routes.GetReservedMagazines(db, authDB)).Methods("GET")
	r.HandleFunc("/removeMagazine/{compID}", routes.DeleteMagazine(db, authDB)).Methods("DELETE")
	r.HandleFunc("/findMagazine", routes.FindMagazines(db, authDB)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

// TODO - create email template
// TODO - create function that is deleting clients, finding clients by Client.Login
// TODO - after creating function that is finding clients, refactore code of email-tool/email.go
// TODO - update endpoints's names
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

	// MAGAZINES ROUTES
	r.HandleFunc("/getMagazines", routes.GetAllMagazines(db, authDB)).Methods("GET")
	r.HandleFunc("/createMagazine", routes.CreateMagazine(db, authDB)).Methods("POST")
	r.HandleFunc("/reserveMagazine", routes.ReserveMagazine(db, authDB)).Methods("PUT")
	r.HandleFunc("/getReservedMagazine", routes.GetReservedMagazines(db, authDB)).Methods("GET")
	r.HandleFunc("/removeMagazine/{compID}", routes.DeleteMagazine(db, authDB)).Methods("DELETE")
	r.HandleFunc("/getMagazine/{compID}", routes.GetMagazine(db, authDB)).Methods("GET")
	r.HandleFunc("/findMagazine", routes.FindMagazines(db, authDB)).Methods("POST")
	r.HandleFunc("/getMagazineReservedBy/{name}", routes.GetMagazineReservedBy(db, authDB)).Methods("GET")

	// CLIENTS ROUTES
	r.HandleFunc("/createClient", routes.CreateClient(db, authDB)).Methods("POST")
	r.HandleFunc("/getClients", routes.GetAllClients(db, authDB)).Methods("GET")

	// CONTACT FUNCTIONS
	r.HandleFunc("/sendEmail/{holder}", routes.GetSendEmailToClient(db, authDB)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

package config

import (
	"api/models"
	"api/routes"
	"database/sql"

	"github.com/gorilla/mux"
)

func CreateRouter(db *sql.DB, authDB []models.Authorize) *mux.Router {
	r := mux.NewRouter()

	// * WAREHOUSES ROUTES *
	r.HandleFunc("/getAllWarehouses", routes.GetAllWarehouses(db, authDB)).Methods("GET")
	r.HandleFunc("/createWarehouse", routes.CreateWarehouse(db, authDB)).Methods("POST")
	r.HandleFunc("/reserveWarehouse", routes.ReserveWarehouse(db, authDB)).Methods("PUT")
	r.HandleFunc("/getReservedWarehouses", routes.GetReservedWarehouses(db, authDB)).Methods("GET")
	r.HandleFunc("/removeWarehouse/{compID}", routes.DeleteWarehouse(db, authDB)).Methods("DELETE")
	r.HandleFunc("/getWarehouse/{compID}", routes.GetWarehouse(db, authDB)).Methods("GET")
	r.HandleFunc("/findWarehouse", routes.FindWarehouses(db, authDB)).Methods("POST")
	r.HandleFunc("/getWarehousesReservedBy/{name}", routes.GetWarehouseReservedBy(db, authDB)).Methods("GET")
	r.HandleFunc("/updateWarehousePrice", routes.UpdateWarehousePrice(db, authDB)).Methods("PUT")

	// * CLIENTS ROUTES *
	r.HandleFunc("/createClient", routes.CreateClient(db, authDB)).Methods("POST")
	r.HandleFunc("/getAllClients", routes.GetAllClients(db, authDB)).Methods("GET")
	r.HandleFunc("/findClient/{holder}", routes.GetClient(db, authDB)).Methods("GET")
	r.HandleFunc("/removeClient/{holder}", routes.DeleteClient(db, authDB)).Methods("DELETE")
	r.HandleFunc("/updateClient", routes.UpdateClient(db, authDB)).Methods("PUT")

	// * CONTACT FUNCTIONS *
	r.HandleFunc("/sendEmail/{holder}", routes.GetSendEmailToClient(db, authDB)).Methods("GET")
	r.HandleFunc("/getPhoneNumber/{holder}", routes.GetPhoneNumber(db, authDB)).Methods("GET")

	return r
}

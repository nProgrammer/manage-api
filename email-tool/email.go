package emailtool

import (
	"api/models"
	"api/repositories"
	"database/sql"
	"fmt"
)

func SendEmail(db *sql.DB, holder string) { // FIXME - clean up
	var mag models.Warehouse
	var mags []models.Warehouse
	mags = repositories.GetWarehouseReservedByDB(db, mag, mags, holder)
	// mags - list of warehouses reserved by $holder
	var clients []models.Client
	clients = repositories.FindClientDB(holder, db, clients)
	email := clients[0].Email
	message := prepareMessage(mags)
	fmt.Printf(email + message)
	sendEmail(message, email)
}

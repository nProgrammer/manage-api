package emailtool

import (
	"api/models"
	"api/repositories"
	"database/sql"
	"fmt"
	"net/smtp"
	"os"
	"strconv"
)

func SendEmail(db *sql.DB, holder string) { // FIXME - clean up
	var mag models.Warehouse
	var mags []models.Warehouse
	mags = repositories.GetWarehouseReservedByDB(db, mag, mags, holder)
	// mags - list of warehouses reserved by $holder
	var client models.Client
	var clients []models.Client
	rows, _ := db.Query("select * from clients where login=$1;", holder)
	for rows.Next() {
		_ = rows.Scan(&client.ID, &client.Name, &client.Login, &client.Phone, &client.Email)
		clients = append(clients, client)
	}
	email := clients[0].Email
	i := 0
	var message string
	message += "You reserved " + strconv.Itoa(len(mags)) + " warehouse(s)"
	for i < len(mags) {
		message += "\n" + mags[i].DateTillReserved + " id of warehouse " + mags[i].CompanyID
		i++
	}
	fmt.Printf(email + message)
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	smtpHost := "smtp.gmail.com"
	var to []string
	msg := []byte(message)
	to = append(to, email)
	smtpPort := "587"
	auth := smtp.PlainAuth("", from, password, smtpHost)
	_ = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
}

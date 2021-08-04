package emailtool

import (
	"api/models"
	"net/smtp"
	"os"
	"strconv"
)

func prepareMessage(mags []models.Warehouse) string {
	i := 0
	var message string
	message += "You reserved " + strconv.Itoa(len(mags)) + " warehouse(s)"
	for i < len(mags) {
		message += "\n" + mags[i].DateTillReserved + " id of warehouse " + mags[i].CompanyID
		i++
	}
	return message
}

func sendEmail(message string, email string) {
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

package routes

import (
	"api/controllers"
	"api/models"
	"api/repositories"
	"api/utils"
	"database/sql"
	"fmt"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllMagazines(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		if mainauS == authDB[0].MainAuth {
			controllers.GetMagazine(db, rw)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

func GetMagazine(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		if mainauS == authDB[0].MainAuth {
			params := mux.Vars(r)
			id := params["compID"]
			controllers.GetMagazineCT(db, rw, id)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

func GetReservedMagazines(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		if mainauS == authDB[0].MainAuth {
			controllers.GetReservedMagazines(db, rw)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

func GetMagazineReservedBy(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		if mainauS == authDB[0].MainAuth {
			params := mux.Vars(r)
			name := params["name"]
			controllers.GetMagazineReservedByCT(db, rw, name)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

func GetAllClients(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		if mainauS == authDB[0].MainAuth {
			controllers.GetClientsCT(db, rw)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

func GetSendEmailToClient(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		params := mux.Vars(r)
		holder := params["holder"]
		if mainauS == authDB[0].MainAuth {
			var mag models.Magazine
			var mags []models.Magazine
			mags = repositories.GetMagazineReservedByDB(db, mag, mags, holder)
			// mags - list of magazines reserved by $holder
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
			message += "You reserved " + strconv.Itoa(len(mags)) + " magazine(s)"
			for i < len(mags) {
				message += "\n" + mags[i].DateTillReserved + " id of magazine " + mags[0].CompanyID
				i++
			}
			fmt.Printf(email + message)
			from := "wagnernorbert836@gmail.com"
			password := "Te$$@2012@"
			smtpHost := "smtp.gmail.com"
			var to []string
			msg := []byte(message)
			to = append(to, email)
			smtpPort := "587"
			auth := smtp.PlainAuth("", from, password, smtpHost)
			_ = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

package routes

import (
	"api/controllers"
	emailtool "api/email-tool"
	"api/models"
	"api/utils"
	"database/sql"
	"net/http"

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
			emailtool.SendEmail(db, holder)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

func GetClient(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		params := mux.Vars(r)
		holder := params["holder"]
		if mainauS == authDB[0].MainAuth {
			controllers.FindClientCT(db, rw, holder)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

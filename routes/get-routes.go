package routes

import (
	"api/controllers"
	emailtool "api/email-tool"
	"api/models"
	"api/utils"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllWarehouses(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		if mainauS == authDB[0].MainAuth {
			controllers.GetWarehouse(db, rw)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

func GetWarehouse(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		if mainauS == authDB[0].MainAuth {
			params := mux.Vars(r)
			id := params["compID"]
			controllers.GetWarehouseCT(db, rw, id)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

func GetReservedWarehouses(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		if mainauS == authDB[0].MainAuth {
			controllers.GetReservedWarehouses(db, rw)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

func GetWarehouseReservedBy(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		if mainauS == authDB[0].MainAuth {
			params := mux.Vars(r)
			name := params["name"]
			controllers.GetWarehouseReservedByCT(db, rw, name)
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
			json.NewEncoder(rw).Encode("Message sent")
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
			client := controllers.FindClientCT(db, holder)
			json.NewEncoder(rw).Encode(client)

		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

func GetPhoneNumber(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		params := mux.Vars(r)
		holder := params["holder"]
		if mainauS == authDB[0].MainAuth {
			client := controllers.FindClientCT(db, holder)
			json.NewEncoder(rw).Encode(client.Phone)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

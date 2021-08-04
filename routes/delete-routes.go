package routes

import (
	"api/controllers"
	"api/models"
	"api/utils"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteMagazine(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["compID"]
		mainauS := utils.AuthorizeMethod(r, authDB)
		if mainauS == authDB[0].MainAuth {
			controllers.DeleteMagazineCT(db, rw, id)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

func DeleteClient(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		params := mux.Vars(r)
		holder := params["holder"]
		if mainauS == authDB[0].MainAuth {
			controllers.DeleteClientCT(db, rw, holder)
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

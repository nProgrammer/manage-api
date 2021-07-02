package routes

import (
	"api/controllers"
	"api/models"
	"api/utils"
	"database/sql"
	"net/http"
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

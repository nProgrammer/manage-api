package routes

import (
	"api/controllers"
	"api/utils"
	"database/sql"
	"net/http"
	"strconv"
)

func CreateMagazine(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		id := controllers.CreateMagazine(r, db)
		idS := strconv.Itoa(id)
		utils.JsonResponse(idS, true, rw)
	}
}

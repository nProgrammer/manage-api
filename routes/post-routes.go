package routes

import (
	"api/controllers"
	"api/models"
	"api/utils"
	"database/sql"
	"encoding/json"
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

func FindMagazines(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var magazineToFind models.Magazine
		json.NewDecoder(r.Body).Decode(&magazineToFind)
		controllers.FindMagazineCT(db, magazineToFind, rw)
	}
}

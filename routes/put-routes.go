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

func ReserveMagazine(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var magazine models.Magazine

		json.NewDecoder(r.Body).Decode(&magazine)

		if magazine.Holder != "" && magazine.Holder != " " {
			magazine.IsReserved = true
		} else {
			utils.JsonResponse("Magazine holder name is empty", false, rw)
		}
		if magazine.CompanyID != "" && magazine.IsReserved == true && magazine.DateTillReserved != "" {
			a := controllers.ReserveMagazineCT(db, magazine)
			res := strconv.FormatInt(a, 10)
			utils.JsonResponse(res, true, rw)
		} else {
			utils.JsonResponse("Missing data!", false, rw)
		}
	}
}
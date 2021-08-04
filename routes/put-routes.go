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

func ReserveWarehouse(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)

		if mainauS == authDB[0].MainAuth {

			var warehouse models.Warehouse

			json.NewDecoder(r.Body).Decode(&warehouse)

			if warehouse.Holder != "" && warehouse.Holder != " " {
				warehouse.IsReserved = true
			} else {
				utils.JsonResponse("Warehouse holder name is empty", false, rw)
			}
			if warehouse.CompanyID != "" && warehouse.IsReserved == true && warehouse.DateTillReserved != "" {
				a := controllers.ReserveWarehouseCT(db, warehouse)
				res := strconv.FormatInt(a, 10)
				utils.JsonResponse(res, true, rw)
			} else {
				utils.JsonResponse("Missing data!", false, rw)
			}
		} else {
			utils.JsonResponse("Bad token!", false, rw)
		}
	}
}

func UpdateClient(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)

		if mainauS == authDB[0].MainAuth {

			var client models.Client

			json.NewDecoder(r.Body).Decode(&client)

			if client.Login != "" && client.Login != " " {

			} else {
				utils.JsonResponse("Client holder name is empty", false, rw)
			}
			a := controllers.UpdateClientCT(db, client)
			res := strconv.FormatInt(a, 10)
			utils.JsonResponse(res, true, rw)
		} else {
			utils.JsonResponse("Missing data!", false, rw)
		}
	}
}

func UpdateWarehousePrice(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)

		if mainauS == authDB[0].MainAuth {

			var warehouse models.Warehouse

			json.NewDecoder(r.Body).Decode(&warehouse)
			a := controllers.UpdateWarehousePriceCT(db, warehouse)
			res := strconv.FormatInt(a, 10)
			utils.JsonResponse(res, true, rw)
		} else {
			utils.JsonResponse("Missing data!", false, rw)
		}
	}
}

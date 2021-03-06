package routes

import (
	"api/controllers"
	"api/models"
	"api/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func CreateWarehouse(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		if mainauS == authDB[0].MainAuth {
			controllers.CreateWarehouse(r, db, rw)
		} else {
			utils.JsonResponse("Bad Token!", false, rw)
		}
	}
}

func FindWarehouses(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		var warehouseToFind models.Warehouse
		json.NewDecoder(r.Body).Decode(&warehouseToFind)
		if mainauS == authDB[0].MainAuth {
			controllers.FindWarehouseCT(db, warehouseToFind, rw)
		} else {
			utils.JsonResponse("Bad Token!", false, rw)
		}
	}
}

func CreateClient(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		var newClient models.Client
		json.NewDecoder(r.Body).Decode(&newClient)
		fmt.Println(newClient)
		if mainauS == authDB[0].MainAuth {
			log.Println("WORKS")
			id := controllers.CreateClientCT(db, newClient, rw)
			utils.JsonResponse(strconv.Itoa(id), true, rw)
		} else {
			utils.JsonResponse("Bad Token!", false, rw)
		}
	}
}

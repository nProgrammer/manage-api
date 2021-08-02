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

func CreateMagazine(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		if mainauS == authDB[0].MainAuth {
			id := controllers.CreateMagazine(r, db)
			idS := strconv.Itoa(id)
			utils.JsonResponse(idS, true, rw)
		} else {
			utils.JsonResponse("Bad Token!", false, rw)
		}
	}
}

func FindMagazines(db *sql.DB, authDB []models.Authorize) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		mainauS := utils.AuthorizeMethod(r, authDB)
		var magazineToFind models.Magazine
		json.NewDecoder(r.Body).Decode(&magazineToFind)
		if mainauS == authDB[0].MainAuth {
			controllers.FindMagazineCT(db, magazineToFind, rw)
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

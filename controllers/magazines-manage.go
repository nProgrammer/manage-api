package controllers

import (
	"api/models"
	"api/repositories"
	"api/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

func CreateMagazine(r *http.Request, db *sql.DB) int {
	var newMag models.Magazine
	json.NewDecoder(r.Body).Decode(&newMag)
	id := repositories.CreateMagazineDB(db, newMag)
	return id
}

func GetMagazine(db *sql.DB, rw http.ResponseWriter) {
	var magazines []models.Magazine
	var magazine models.Magazine
	magazines = repositories.GetMagazinesDB(db, magazine, magazines)
	json.NewEncoder(rw).Encode(magazines)
}

func GetMagazineCT(db *sql.DB, rw http.ResponseWriter, id string) {
	var magazines []models.Magazine
	var magazine models.Magazine
	magazines = repositories.GetMagazineDB(db, magazine, magazines, id)
	json.NewEncoder(rw).Encode(magazines)
}

func ReserveMagazineCT(db *sql.DB, magazine models.Magazine) int64 {
	rowsUpd := repositories.ReserveMagazineDB(db, magazine)
	return rowsUpd
}

func GetReservedMagazines(db *sql.DB, rw http.ResponseWriter) {
	var magazines []models.Magazine
	var magazine models.Magazine
	magazines = repositories.GetReservedMagazinesDB(db, magazine, magazines)
	json.NewEncoder(rw).Encode(magazines)
}

func DeleteMagazineCT(db *sql.DB, rw http.ResponseWriter, compID string) {
	text := "Deleted " + compID
	repositories.DeleteMagazineDB(db, compID)
	utils.JsonResponse(text, true, rw)
}

func FindMagazineCT(db *sql.DB, magazineToFind models.Magazine, rw http.ResponseWriter) {
	var magazines []models.Magazine
	magazines = repositories.FindMagazinesDB(db, magazineToFind, magazines)
	json.NewEncoder(rw).Encode(magazines)
}

package controllers

import (
	"api/models"
	"api/repositories"
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

func ReserveMagazineCT(db *sql.DB, magazine models.Magazine) int64 {
	rowsUpd := repositories.ReserveMagazineDB(db, magazine)
	return rowsUpd
}

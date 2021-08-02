package controllers

import (
	"api/models"
	"api/repositories"
	"database/sql"
	"net/http"
)

func CreateClientCT(db *sql.DB, client models.Client, rw http.ResponseWriter) int {
	id := repositories.CreateClientDB(db, client)
	return id
}

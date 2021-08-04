package controllers

import (
	"api/models"
	"api/repositories"
	"api/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

func CreateClientCT(db *sql.DB, client models.Client, rw http.ResponseWriter) int {
	id := repositories.CreateClientDB(db, client)
	return id
}

func GetClientsCT(db *sql.DB, rw http.ResponseWriter) {
	var clients []models.Client
	var client models.Client
	clients = repositories.GetClientsDB(db, client, clients)
	json.NewEncoder(rw).Encode(clients)
}

func FindClientCT(db *sql.DB, rw http.ResponseWriter, login string) {
	var clients []models.Client
	clients = repositories.FindClientDB(login, db, clients)
	json.NewEncoder(rw).Encode(clients)
}

func DeleteClientCT(db *sql.DB, rw http.ResponseWriter, holder string) {
	text := "Deleted " + holder
	repositories.DeleteClientDB(db, holder)
	utils.JsonResponse(text, true, rw)
}

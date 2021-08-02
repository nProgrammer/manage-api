package controllers

import (
	"api/models"
	"api/repositories"
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

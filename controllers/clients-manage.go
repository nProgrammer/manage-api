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

func UpdateClientCT(db *sql.DB, client models.Client) int64 {
	rowsUpd := repositories.UpdateClientDB(db, client)
	return rowsUpd
}

func GetClientsCT(db *sql.DB, rw http.ResponseWriter) {
	var clients []models.Client
	var client models.Client
	clients = repositories.GetClientsDB(db, client, clients)
	json.NewEncoder(rw).Encode(clients)
}

func FindClientCT(db *sql.DB, login string) models.Client {
	var clients []models.Client
	clients = repositories.FindClientDB(login, db, clients)
	return clients[0]
}

func DeleteClientCT(db *sql.DB, rw http.ResponseWriter, holder string) {
	text := "Deleted " + holder
	repositories.DeleteClientDB(db, holder)
	utils.JsonResponse(text, true, rw)
}
